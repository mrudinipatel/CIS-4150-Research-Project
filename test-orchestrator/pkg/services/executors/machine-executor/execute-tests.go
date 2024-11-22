package machineexecutor

import (
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// ExecuteTests implements domain.TestExecutor.
func (m *MachineExecutor) ExecuteTests(project domain.Project, workspace domain.Workspace, testset *testset.TestSet) error {
	var results = make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(5)

	for _, tests := range testset.Split(5) {
		go m.runTests(tests, project, workspace, results, &wg)
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

func (m *MachineExecutor) runTests(tests []string, project domain.Project, workspace domain.Workspace, co chan<- interface{}, wg *sync.WaitGroup) {
	arguments := strings.Split(project.GetTestCommand(tests), " ")

	cmd := exec.Command(arguments[0], arguments[1:]...)
	cmd.Dir = workspace.GetPath()

	if output, err := cmd.Output(); err != nil {
		co <- err
	} else {
		co <- string(output)
	}

	wg.Done()
}
