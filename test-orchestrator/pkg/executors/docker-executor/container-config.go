package dockerexecutor

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type ContainerConfig struct {
	image  string
	memory string
	cpus   string
}

func NewContainerConfig(image string, memory string, cpus string) *ContainerConfig {
	return &ContainerConfig{
		image,
		memory,
		cpus,
	}
}

func (i *ContainerConfig) Run(command string, workspace domain.Workspace) (string, error) {
	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-m", i.memory,
		fmt.Sprintf("--cpus=%s", i.cpus),
		"-v", fmt.Sprintf("%s:%s", workspace.GetName(), workspace.GetPath()),
		"-w", workspace.GetPath(),
		"--entrypoint", "/bin/sh",
		i.image,
		"-c",
		command,
	)

	if output, err := cmd.Output(); err != nil {
		log.Fatal(err)
		return "", err
	} else {
		return string(output), nil
	}
}
