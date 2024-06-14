package core

import (
	"os"
	"text/template"
)

type ServiceData struct {
	Name     string
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
	 id = "{{.Name}}"
	 tags = [{{range .Tags}}"{{.}}",{{end}}]
	 port = {{.Port}}
   
	 check
	 {
	   name = "{{.Name}} status check",
	   id =  "{{.Name}}",
	   service_id = "{{.Name}}",
		 tcp  = "localhost:{{.Port}}",
		 interval = "{{.Interval}}",
		 timeout = "{{.Timeout}}"
	 }
   }`

	file, err := os.Create("." + "/" + "temp" + "/" + service.Name + ".hcl")
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
