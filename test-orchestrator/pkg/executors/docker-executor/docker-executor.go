package dockerexecutor

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type DockerExecutor struct {
	container     *ContainerConfig
	numContainers int
}

func Create(container *ContainerConfig, numContainers int) domain.TestExecutor {
	return &DockerExecutor{
		container,
		numContainers,
	}
}
