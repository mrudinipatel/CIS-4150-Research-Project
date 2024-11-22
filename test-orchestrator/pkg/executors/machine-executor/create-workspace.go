package machineexecutor

import (
	"os"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type Directory struct {
	name string
}

func (d *Directory) GetPath() string {
	return "/tmp/" + d.name
}

func (d *Directory) GetName() string {
	return d.name
}

func (m *MachineExecutor) CreateWorkspace() (domain.Workspace, error) {
	d := &Directory{
		name: domain.RandSeq(10),
	}

	if err := os.Mkdir(d.GetPath(), 0750); err != nil {
		return nil, err
	} else {
		return d, nil
	}
}

func (m *MachineExecutor) CleanupWorkspace(workspace domain.Workspace) error {
	return os.RemoveAll(workspace.GetPath())
}
