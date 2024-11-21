package machine

import (
	"log"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// ExecuteTests implements domain.TestExecutor.
func (k *MachineExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *testset.TestSet) error {
	log.Println("Running a few tests...")

	var results = make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(5)

	for _, tests := range testset.Split(5) {
		go func(co chan<- interface{}) {
			cmd := project.GetTestCommand(tests)

			cmd.Dir = workspace.GetPath()

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
