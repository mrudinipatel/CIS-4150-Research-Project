package docker

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type DockerWorkspace struct {
	name      string
	path      string
	container ContainerConfig
}

type ContainerConfig struct {
	Image  *DockerImage
	Memory string
	Cpus   string
}

func CreateWorkspace(config domain.WorkspaceConfig) (domain.Workspace, error) {
	cmd := exec.Command("docker", "volume", "create")

	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	name := strings.TrimSpace(string(out))
	container := config.(ContainerConfig)

	return &DockerWorkspace{
		name:      name,
		path:      "/tmp/" + name,
		container: container,
	}, nil
}

func (w *DockerWorkspace) Cleanup() error {
	cmd := exec.Command("docker", "volume", "rm", w.name)
	return cmd.Run()
}

func (w *DockerWorkspace) Run(command string) (string, error) {
	return w.RunWithConfig(command, w.container)
}

func (w *DockerWorkspace) RunWithConfig(command string, config domain.WorkspaceConfig) (string, error) {
	container := config.(ContainerConfig)

	cmd := exec.Command(
		"docker",
		"run",
		"--rm",
		"-m", container.Memory,
		fmt.Sprintf("--cpus=%s", container.Cpus),
		"-v", fmt.Sprintf("%s:%s", w.name, w.path),
		// FIXME: we're only working with maven for now, make generic in the future
		"-v", "maven:/root/.m2",
		"-e", fmt.Sprintf("MAVEN_OPTS=-Xmx%s -Xms%s", container.Memory, container.Memory),
		// -----------
		"-w", w.path,
		"--entrypoint", "/bin/bash",
		container.Image.GetTag(),
		"-c",
		command,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	} else {
		return stdout.String(), nil
	}
}

type result struct {
	output string
	err    error
}

func (w *DockerWorkspace) RunParallelWithConfig(commands []string, config domain.WorkspaceConfig) error {
	outputResult := false

	var wg sync.WaitGroup
	wg.Add(len(commands))

	out := make(chan result)
	defer close(out)

	for _, command := range commands {
		go func(co chan<- result, wg *sync.WaitGroup) {
			output, err := w.RunWithConfig(command, config)
			co <- result{output, err}
			wg.Done()
		}(out, &wg)
	}

	go func() {
		for result := range out {
			if outputResult && result.err != nil {
				log.Println(result.output)
			}
		}
	}()

	wg.Wait()
	return nil
}
