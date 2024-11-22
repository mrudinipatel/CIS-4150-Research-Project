package dockerexecutor

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// ExecuteTests implements domain.TestExecutor.
func (d *DockerExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *testset.TestSet) error {
	volume, isVolume := workspace.(*Volume)

	if !isVolume {
		return errors.New("expected volume workspace")
	}

	var results = make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(5)

	for _, tests := range testset.Split(5) {
		go func(co chan<- interface{}) {
			cmd := exec.Command(
				"docker",
				"run",
				"--rm",
				"-v", fmt.Sprintf("%s:%s", volume.GetName(), volume.GetPath()),
				"-w", volume.GetPath(),
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
		}(results)
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
