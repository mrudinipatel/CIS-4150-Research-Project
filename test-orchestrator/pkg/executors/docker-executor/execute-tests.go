package dockerexecutor

import (
	"log"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

// ExecuteTests implements domain.TestExecutor.
func (d *DockerExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *domain.TestSet) error {
	var results = make(chan interface{})
	defer close(results)

	var wg sync.WaitGroup
	wg.Add(d.numContainers)

	for _, tests := range testset.Split(d.numContainers) {
		go func(co chan<- interface{}, wg *sync.WaitGroup) {
			if output, err := d.container.Run(project.GetTestCommand(tests), workspace); err != nil {
				co <- err
			} else {
				co <- output
			}

			wg.Done()
		}(results, &wg)
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
	return nil
}
