package machineexecutor

import (
	"os/exec"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

func (k *MachineExecutor) SetupProject(project domain.Project, workspace domain.Workspace) (*domain.TestSet, error) {
	cmd := exec.Command("git", "clone", project.GetCloneUrl(), workspace.GetPath())

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	arguments := strings.Split(project.GetSetupCommand(), " ")

	cmd = exec.Command(arguments[0], arguments[1:]...)
	cmd.Dir = workspace.GetPath()

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	getTests := exec.Command("find", ".", "-type", "f", "-name", project.GetTestFilter(), "-exec", "basename", "{}", "\\;")
	getTests.Dir = workspace.GetPath()

	if output, err := getTests.Output(); err != nil {
		return nil, err
	} else {
		return domain.NewTestSet(strings.Split(string(output), "\n")), nil
	}
}
