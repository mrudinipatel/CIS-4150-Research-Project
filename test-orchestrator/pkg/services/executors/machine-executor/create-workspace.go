package machineexecutor

import (
	"math/rand"
	"os"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type Directory struct {
	path string
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (pv *Directory) GetPath() string {
	return pv.path
}

// CreateWorkspace implements domain.TestExecutor.
func (k *MachineExecutor) CreateWorkspace() (domain.Workspace, error) {
	path := "/tmp/" + randSeq(10)

	err := os.Mkdir(path, 0750)

	if err != nil {
		return nil, err
	}

	return &Directory{
		path: path,
	}, nil
}

func (k *MachineExecutor) CleanupWorkspace(workspace domain.Workspace) error {
	return os.RemoveAll(workspace.GetPath())
}
