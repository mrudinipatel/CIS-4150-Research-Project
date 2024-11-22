package dockerexecutor

import (
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// ExecuteTests implements domain.TestExecutor.
func (d *DockerExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *testset.TestSet) error {
	var results = make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(5)

	for _, tests := range testset.Split(5) {
		go d.runTests(tests, project, workspace, results, &wg)
	}

	go func() {
		for result := range results {
			switch result.(type) {
			case string:
				log.Println(result)
			case error:
				log.Panic(result)
			}
		}
	}()

	wg.Wait()
	close(results)
	return nil
}

func (d *DockerExecutor) runTests(tests []string, project domain.Project, workspace domain.Workspace, co chan<- interface{}, wg *sync.WaitGroup) {
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-v", fmt.Sprintf("%s:%s", workspace.GetName(), workspace.GetPath()),
		"-w", workspace.GetPath(),
		"--entrypoint", "/bin/sh",
		d.image,
		"-c",
		project.GetTestCommand(tests),
	)

	if output, err := cmd.Output(); err != nil {
		co <- err
	} else {
		co <- string(output)
	}

	wg.Done()
}
