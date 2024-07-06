package core

import (
	"consul-companion/internal/cfg"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

func RunCreatesServices(errCh chan error) {
	var services []ServiceData

	litsProjects := getListProjects(errCh)
	projectsWhithServices := getListEnv(litsProjects, errCh)

	for _, project := range projectsWhithServices {
		for _, env := range project.Env {
			if strings.Contains(env.Key, "#") {
				continue
			}
			project.convertServiceFile(env, &services)
		}
	}

	for _, service := range services {
		mkDir(cfg.TMP_DIR)

		tmpFile := path.Join(cfg.TMP_DIR, service.Name+"-"+service.Project+".tmp")
		currentFile := path.Join(cfg.CONFDIR, service.Name+"-"+service.Project+".hcl")

		service.createServiceFile(tmpFile)
		ok := DiffChecksum(currentFile, tmpFile)

		if !ok {
			log.Println("File changed", currentFile)
			service.createServiceFile(currentFile)
		}
	}
}

func (s *ServiceData) createServiceFile(path string) {

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
	err = t.Execute(file, s)
	if err != nil {
		panic(err)
	}
}

func (p *Project) convertServiceFile(env Env, services *[]ServiceData) {

	if strings.Contains(env.Key, "EXT") {
		svcName := strings.Replace(env.Key, "EXT_", "", 1)
		svcName = strings.Replace(svcName, "_PORT", "", 1)
		svcName = strings.ToLower(svcName)

		parts := strings.Split(p.Name, "-")
		index := len(parts) - 1

		partsPrjName := strings.Split(p.Name, "-")
		projectName := strings.Join(partsPrjName[:len(parts)-1], "-")

		*services = append(*services, ServiceData{
			Name:     svcName,
			Project:  p.Name,
			Tags:     []string{svcName, parts[index], projectName, fmt.Sprintf(projectName + "-" + parts[index])},
			Port:     env.Value,
			Interval: "5s",
			Timeout:  "5s",
		})
	}
}
