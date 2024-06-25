package core

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

type ServiceData struct {
	Name     string
	Project  string
	Tags     []string
	Port     string
	Interval string
	Timeout  string
}

var CONFDIR string

func RunCreatesServices(errCh chan error) {

	prjs := getListProjects(errCh)
	p := getListEnv(prjs, errCh)
	var services []ServiceData

	for _, prj := range p {
		for _, env := range prj.Env {
			if strings.Contains(env.Key, "#") {
				continue
			}
			convertServiceFile(env, prj, &services)
		}
	}

	for _, service := range services {
		targetDir := "/tmp/consul-companion"
		mkDir(targetDir)

		tmpFile := path.Join(targetDir, service.Name+"-"+service.Project+".tmp")
		currentFile := path.Join(CONFDIR, service.Name+"-"+service.Project+".hcl")

		createServiceFile(service, tmpFile)
		ok := DiffChecksum(currentFile, tmpFile)

		if !ok {
			log.Println("File changed", currentFile)
			createServiceFile(service, currentFile)
		}
	}
}

func createServiceFile(service ServiceData, path string) {

	// Создаем шаблон для файла
	tmpl := `## -----------------------------
## Generation by consul-companion
## -----------------------------

service {
	name = "{{.Name}}"
	id = "{{.Name}}-{{.Project}}"
	tags = [{{range .Tags}}"{{.}}",{{end}}]
	port = {{.Port}}

	check
	{
	  name = "{{.Name}} status check",
	  id =  "{{.Name}}-{{.Project}}",
	  service_id = "{{.Name}}-{{.Project}}",
		tcp  = "localhost:{{.Port}}",
		interval = "{{.Interval}}",
		timeout = "{{.Timeout}}"
	}
   }`

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	t := template.Must(template.New("serviceTemplate").Parse(tmpl))
	err = t.Execute(file, service)
	if err != nil {
		panic(err)
	}
}

func convertServiceFile(env Env, prj Project, services *[]ServiceData) {

	if strings.Contains(env.Key, "EXT") {
		svcName := strings.Replace(env.Key, "EXT_", "", 1)
		svcName = strings.Replace(svcName, "_PORT", "", 1)
		svcName = strings.ToLower(svcName)

		parts := strings.Split(prj.Name, "-")
		index := len(parts) - 1

		partsPrjName := strings.Split(prj.Name, "-")
		projectName := strings.Join(partsPrjName[:len(parts)-1], "-")

		*services = append(*services, ServiceData{
			Name:     svcName,
			Project:  prj.Name,
			Tags:     []string{svcName, parts[index], projectName, fmt.Sprintf(projectName + "-" + parts[index])},
			Port:     env.Value,
			Interval: "5s",
			Timeout:  "5s",
		})
	}
}
