package dockerexecutor

import (
	"fmt"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

func (d *DockerExecutor) SetupProject(project domain.Project, workspace domain.Workspace) (*domain.TestSet, error) {
	if output, err := d.container.Run(d.setupCommand(project, workspace), workspace); err != nil {
		return nil, err
	} else {
		return domain.NewTestSet(strings.Split(output, "\n")), nil
	}
}

func (d *DockerExecutor) setupCommand(project domain.Project, workspace domain.Workspace) string {
	return fmt.Sprintf(
		"git clone %s %s 1>&2 && %s 1>&2 && %s",
		project.GetCloneUrl(),
		workspace.GetPath(),
		project.GetSetupCommand(),
		project.GetTestFilterCommand(),
	)
}
