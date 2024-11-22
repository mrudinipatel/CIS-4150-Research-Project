package dockerexecutor

import (
	"os/exec"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type Volume struct {
	name string
}

func (v *Volume) GetPath() string {
	return "/tmp/" + v.GetName()
}

func (v *Volume) GetName() string {
	return v.name
}

// CreateWorkspace implements domain.TestExecutor.
func (d *DockerExecutor) CreateWorkspace() (domain.Workspace, error) {
	cmd := exec.Command("docker", "volume", "create")

	if name, err := cmd.Output(); err != nil {
		return nil, err
	} else {
		return &Volume{
			name: strings.TrimSpace(string(name)),
		}, nil
	}
}

func (d *DockerExecutor) CleanupWorkspace(workspace domain.Workspace) error {
	cmd := exec.Command("docker", "volume", "rm", workspace.GetName())
	return cmd.Run()
}
