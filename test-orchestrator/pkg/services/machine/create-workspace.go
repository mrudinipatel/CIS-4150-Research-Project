package machine

import (
	"log"
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
	log.Println("Creating Persistant Volume...")
	path := "/tmp/" + randSeq(10)
	os.Mkdir(path, 0750)
	return &Directory{
		path: path,
	}, nil
}
