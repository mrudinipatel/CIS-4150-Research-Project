package machine

import (
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// SetupProject implements domain.TestExecutor.
func (k *MachineExecutor) SetupProject(project domain.Project, workspace domain.Workspace) (*testset.TestSet, error) {
	log.Println("Setting up project...")
	cmd := exec.Command("git", "clone", project.GetCloneUrl(), workspace.GetPath())

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	setup := project.GetSetupCommand()

	setup.Dir = workspace.GetPath()

	if err := setup.Run(); err != nil {
		return nil, err
	}

	// Find Test Files
	log.Println("Discovering test files...")

	getTests := exec.Command("find", ".", "-type", "f", "-name", "*Test*.java")
	getTests.Dir = workspace.GetPath()

	tests := []string{}

	if output, err := getTests.Output(); err != nil {
		return nil, err
	} else {
		for _, testPath := range strings.Split(string(output), "\n") {
			filename := path.Base(testPath)
			extension := path.Ext(filename)
			tests = append(tests, filename[0:len(filename)-len(extension)])
		}
	}

	return testset.Create(tests), nil
}
