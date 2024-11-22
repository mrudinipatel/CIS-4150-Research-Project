package dockerexecutor

import (
	"log"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

// ExecuteTests implements domain.TestExecutor.
func (d *DockerExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *domain.TestSet) error {
	var wg sync.WaitGroup
	wg.Add(d.numContainers)

	for _, tests := range testset.Split(d.numContainers) {
		go func(wg *sync.WaitGroup) {
			if output, err := d.container.Run(project.GetTestCommand(tests), workspace); err != nil {
				log.Fatal(output)
			}
			wg.Done()
		}(&wg)
	}

	wg.Wait()
	return nil
}
