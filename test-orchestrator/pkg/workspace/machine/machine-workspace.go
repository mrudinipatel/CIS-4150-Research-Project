package machine

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type Directory struct {
	name   string
	path   string
	config domain.WorkspaceConfig
}

func CreateWorkspace(config domain.WorkspaceConfig) (domain.Workspace, error) {
	name := domain.RandSeq(10)
	d := &Directory{
		name:   name,
		path:   "/tmp/" + name,
		config: config,
	}

	if err := os.Mkdir(d.path, 0750); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

func (d *Directory) Cleanup() error {
	return os.RemoveAll(d.path)
}

// Run implements domain.Workspace.
func (d *Directory) Run(command string) (string, error) {
	return d.RunWithConfig(command, d.config)
}

// RunWithConfig implements domain.Workspace.
func (d *Directory) RunWithConfig(command string, _ domain.WorkspaceConfig) (string, error) {
	arguments := strings.Split(command, " ")
	cmd := exec.Command(arguments[0], arguments[1:]...)

	cmd.Dir = d.path

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	} else {
		return stdout.String(), nil
	}
}

// RunParallelWithConfig implements domain.Workspace.
func (d *Directory) RunParallelWithConfig(commands []string, config domain.WorkspaceConfig) error {
	var wg sync.WaitGroup
	wg.Add(len(commands))
	out := make(chan string)

	for i, command := range commands {
		go func(out chan<- string, wg *sync.WaitGroup) {
			if _, err := d.RunWithConfig(command, config); err != nil {
				out <- fmt.Sprintf("Test Set %d Completed: Failed\n", i+1)
			} else {
				out <- fmt.Sprintf("Test Set %d Completed: Passed\n", i+1)
			}
			wg.Done()
		}(out, &wg)
	}

	go func() {
		for result := range out {
			log.Println(result)
		}
	}()

	wg.Wait()
	return nil
}
