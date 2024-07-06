package core

import (
	"consul-companion/internal/cfg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getListProjects(errCh chan error) (projects []Project) {

	dir, err := os.Open(cfg.ROOT_PROJECT_PATH)
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
			log.Println("skip", project.Name())
			continue

		}
		if project.Name() == "containerd" {
			continue
		}
		projects = append(projects, Project{
			project.Name(),
			filepath.Join(cfg.ROOT_PROJECT_PATH, project.Name()),
			filepath.Join(cfg.ROOT_PROJECT_PATH, project.Name(), ".env"), nil})
	}
	return projects
}

func getListEnv(projects []Project, errCh chan error) (tmpProjects []Project) {

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

			tmpProjects = append(tmpProjects,
				Project{project.Name,
					project.Path,
					project.DotEnv, append(project.Env, Env{key, value})})
		}
	}

	return tmpProjects
}
