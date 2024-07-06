package core

type Project struct {
	Name   string
	Path   string
	DotEnv string
	Env    []Env
}

type Env struct {
	Key   string
	Value string
}

type ServiceData struct {
	Name     string
	Project  string
	Tags     []string
	Port     string
	Interval string
	Timeout  string
}
