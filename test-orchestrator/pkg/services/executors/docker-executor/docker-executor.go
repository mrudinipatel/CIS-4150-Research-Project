package dockerexecutor

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type DockerExecutor struct {
	image string
}

func Create(image string) domain.TestExecutor {
	return &DockerExecutor{
		image: image,
	}
}
