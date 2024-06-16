package core

import (
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

var Path = "/opt"

func getListProjects(errCh chan error) []Project {
	var projects []Project

	dir, err := os.Open(Path)
	if err != nil {
		errCh <- err
		return nil
	}
	defer dir.Close()

	folder, err := dir.ReadDir(-1)
	if err != nil {
		errCh <- err
		return nil
	}

	for _, project := range folder {
		if !project.IsDir() {
			continue
		}
		if project.Name() == "containerd" {
			continue
		}
		projects = append(projects, Project{project.Name(), filepath.Join(Path, project.Name()), filepath.Join(Path, project.Name(), ".env"), nil})

	}
	return projects
}

func getListEnv(projects []Project, errCh chan error) []Project {
	var prjs []Project

	for _, project := range projects {
		content, err := os.ReadFile(project.DotEnv)
		if err != nil {
			errCh <- err
			continue
		}

		lines := strings.Split(string(content), "\n")
		for _, line := range lines {
			pair := strings.Split(line, "=")
			if len(pair) != 2 {
				continue
			}
			key := strings.TrimSpace(pair[0])
			value := strings.TrimSpace(pair[1])

			prjs = append(prjs, Project{project.Name, project.Path, project.DotEnv, append(project.Env, Env{key, value})})
		}
	}

	return prjs
}
