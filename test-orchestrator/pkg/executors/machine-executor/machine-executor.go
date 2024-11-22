package machineexecutor

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type MachineExecutor struct{}

func Create() domain.TestExecutor {
	return &MachineExecutor{}
}
