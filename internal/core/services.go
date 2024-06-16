package core

import (
	"os"
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

func CreateServiceFile(service ServiceData) {

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

	path := "./" + "temp"
	fullName := path + "/" + service.Name + "-" + service.Project + ".hcl"

	file, err := os.Create(fullName)
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
