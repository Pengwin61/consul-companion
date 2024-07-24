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

func getScanFolder(errCh chan error) []string {
	var tmpServices []string

	dir, err := os.Open(cfg.CONFDIR)
	if err != nil {
		errCh <- err
	}

	defer dir.Close()

	folder, err := dir.ReadDir(-1)
	if err != nil {
		errCh <- err
	}

	for _, project := range folder {
		if !project.IsDir() {
			tmpServices = append(tmpServices, filepath.Join(cfg.CONFDIR, project.Name()))
		}

	}
	return tmpServices
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			if !found {
				diff = append(diff, s1)
			}
		}
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}
