package dockerexecutor

import (
	"fmt"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

// SetupProject implements domain.TestExecutor.
func (d *DockerExecutor) SetupProject(project domain.Project, workspace domain.Workspace) (*domain.TestSet, error) {
	if output, err := d.container.Run(d.setupCommand(project, workspace), workspace); err != nil {
		return nil, err
	} else {
		return domain.NewTestSet(strings.Split(output, "\n")), nil
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
