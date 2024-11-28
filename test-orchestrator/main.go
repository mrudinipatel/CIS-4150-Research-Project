package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/workspace/docker"
)

type TestResult struct {
	config        int
	numContainers int
	times         []int
}

func main() {
	dockerImage, err := docker.BuildImage(".")
	if err != nil {
		log.Panic(err)
	}

	configs := []docker.ContainerConfig{
		{Image: dockerImage, Memory: "1g", Cpus: "2"}, // t3.micro
		{Image: dockerImage, Memory: "2g", Cpus: "2"}, // t3.small
		{Image: dockerImage, Memory: "4g", Cpus: "2"}, // t3.medium
		{Image: dockerImage, Memory: "2g", Cpus: "1"},
		{Image: dockerImage, Memory: "2g", Cpus: "2"},
		{Image: dockerImage, Memory: "2g", Cpus: "4"},
	}

	containers := []int{1, 2, 3}

	workspace, err := docker.CreateWorkspace(docker.ContainerConfig{Image: dockerImage, Memory: "4g", Cpus: "4"})
	if err != nil {
		log.Panic(err)
	}

	defer workspace.Cleanup()

	// project, err := domain.CreateMavenProject("jsoup", "https://github.com/jhy/jsoup.git", workspace)
	project, err := domain.CreateMavenProjectWithTestModule("guava", "https://github.com/google/guava.git", "guava-tests", workspace)
	if err != nil {
		log.Panic(err)
	}

	results := RunConfigurations(configs, containers, project)
	OutputResults(project, results)
}

func RunConfigurations(configs []docker.ContainerConfig, containers []int, project domain.Project) []TestResult {
	results := []TestResult{}

	for i, config := range configs {
		for _, n := range containers {
			log.Printf("Running Test with %s Memory, %s CPU(s), %d Container(s)\n", config.Memory, config.Cpus, n)

			if times, err := ExecTestSuite(project, config, n); err != nil {
				log.Panic(err)
			} else {
				results = append(results, TestResult{config: i, numContainers: n, times: times})
			}

			// sleep for 2 minutes to allow the CPU to cool
			log.Println("Sleeping for 2 minutes...")
			time.Sleep(time.Duration(time.Minute * 2))
		}
	}

	return results
}

func ExecTestSuite(project domain.Project, config domain.WorkspaceConfig, numContainers int) ([]int, error) {
	times := []int{}
	numIterations := 3

	for i := 0; i < numIterations; i++ {
		start := time.Now()

		if err := project.RunTestsParallelWithConfig(numContainers, config); err != nil {
			return nil, err
		}

		times = append(times, int(time.Since(start).Milliseconds()))
	}

	return times, nil
}

func OutputResults(project domain.Project, results []TestResult) {
	filename := fmt.Sprintf("%s.csv", project.GetName())
	f, err := os.Create(filename)

	if err != nil {
		log.Panic(err)
	}

	defer f.Close()

	f.WriteString("Config,NContainers,Times\n")

	for _, result := range results {
		f.WriteString(fmt.Sprintf("%d,%d,", result.config, result.numContainers))

		for _, time := range result.times {
			f.WriteString(fmt.Sprintf("%d ", time))
		}

		f.WriteString("\n")
	}
}
