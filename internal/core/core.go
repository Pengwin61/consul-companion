package core

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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

const path = "/opt/project"

func GetProjects() ([]Project, error) {
	var projects []Project

	dir, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer dir.Close()

	folder, err := dir.ReadDir(-1)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, project := range folder {
		if !project.IsDir() {
			continue
		}
		projects = append(projects, Project{project.Name(), filepath.Join(path, project.Name()), filepath.Join(path, project.Name(), ".env"), nil})

	}
	return projects, err
}

func GetEnv(projects []Project) []Project {
	var prjs []Project

	for _, project := range projects {
		content, err := os.ReadFile(project.DotEnv)
		if err != nil {
			log.Fatalf("Ошибка чтения файла: %v", err)
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			pair := strings.Split(line, "=")
			if len(pair) != 2 {
				continue
			}
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])

			// project.Env = append(project.Env, Env{key, value})

			prjs = append(prjs, Project{project.Name, project.Path, project.DotEnv, append(project.Env, Env{key, value})})
		}
	}

	return prjs
}
