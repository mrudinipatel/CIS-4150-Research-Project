package dockerexecutor

import (
	"bytes"
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
		// FIXME: we're only working with maven for now, make generic in the future
		"-v", "maven:/root/.m2",
		"-e", fmt.Sprintf("MAVEN_OPTS=-Xmx%s -Xms%s", i.memory, i.memory),
		// -----------
		"-w", workspace.GetPath(),
		"--entrypoint", "/bin/sh",
		i.image,
		"-c",
		command,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if output, err := cmd.Output(); err != nil {
		log.Println(string(output))
		log.Println(stderr.String())
		return "", err
	} else {
		return string(output), nil
	}
}
