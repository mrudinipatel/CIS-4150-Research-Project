package main

import (
	"fmt"
	"log"
	"time"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/workspace/docker"
)

func main() {
	dockerImage, err := docker.BuildImage(".")

	if err != nil {
		log.Panic(err)
	}

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

	// project, err := domain.CreateMavenProjectWithTestModule("https://github.com/google/guava.git", "guava-tests", workspace)
	project, err := domain.CreateMavenProject("https://github.com/jhy/jsoup.git", workspace)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("config,containers,duration")

	for i, config := range configs {
		for _, n := range containers {
			if duration, err := ExecTestSuite(project, config, n); err != nil {
				log.Panic(err)
			} else {
				fmt.Printf("%d,%d,%d\n", i, n, duration)
			}

			time.Sleep(time.Duration(time.Minute * 2)) // sleep for 2 minutes to allow the CPU to cool
		}
	}
}

func ExecTestSuite(project domain.Project, config domain.WorkspaceConfig, numContainers int) (int, error) {
	totalTime := 0
	numIterations := 3

	for i := 0; i < numIterations; i++ {
		start := time.Now()

		if err := project.RunTestsParallelWithConfig(numContainers, config); err != nil {
			return -1, err
		}

		totalTime += int(time.Since(start).Milliseconds())
	}

	avgTime := totalTime / numIterations

	return avgTime, nil
}
