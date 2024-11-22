package dockerexecutor

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

// SetupProject implements domain.TestExecutor.
func (d *DockerExecutor) SetupProject(project domain.Project, workspace domain.Workspace) (*testset.TestSet, error) {
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-v", fmt.Sprintf("%s:%s", workspace.GetName(), workspace.GetPath()),
		"-w", workspace.GetPath(),
		"--entrypoint", "/bin/sh",
		d.image,
		"-c",
		d.setupCommand(project, workspace),
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, err
	} else {
		return testset.Create(strings.Split(string(output), "\n")), nil
	}
}

func (d *DockerExecutor) setupCommand(project domain.Project, workspace domain.Workspace) string {
	return fmt.Sprintf(
		"git clone %s %s > /dev/null 2>&1 && %s > /dev/null 2>&1 && find . -type f -name %s -exec basename {} \\;",
		project.GetCloneUrl(),
		workspace.GetPath(),
		project.GetSetupCommand(),
		project.GetTestFilter(),
	)
}
