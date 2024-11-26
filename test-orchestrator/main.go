package main

import (
	"log"
	"time"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/workspace/docker"
)

func main() {
	dockerImage := docker.NewDockerImage("test")

	configs := []docker.ContainerConfig{
		{Image: dockerImage, Memory: "1g", Cpus: "2"}, // t3.micro
		{Image: dockerImage, Memory: "2g", Cpus: "2"}, // t3.small
		{Image: dockerImage, Memory: "4g", Cpus: "2"}, // t3.medium
	}

	containers := []int{1, 2, 3}

	workspace, err := docker.CreateWorkspace(docker.ContainerConfig{Image: dockerImage, Memory: "4g", Cpus: "4"})

	if err != nil {
		log.Panic(err)
	}

	defer workspace.Cleanup()

	project, err := domain.CreateMavenProjectWithTestModule("https://github.com/google/guava.git", "guava-tests", workspace)
	// project, err := domain.CreateMavenProject("https://github.com/jhy/jsoup.git", workspace)

	if err != nil {
		log.Panic(err)
	}

	for _, config := range configs {
		for _, n := range containers {
			if duration, err := ExecTestSuite(project, config, n); err != nil {
				log.Panic(err)
			} else {
				log.Println(duration)
			}
		}
	}
}

func ExecTestSuite(project domain.Project, config domain.WorkspaceConfig, n int) (time.Duration, error) {
	curr := time.Now()

	if err := project.RunTestsParallelWithConfig(n, config); err != nil {
		return -1, err
	}

	return time.Since(curr), nil
}
