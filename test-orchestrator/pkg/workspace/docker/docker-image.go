package docker

import (
	"log"
	"os/exec"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type DockerImage struct {
	tag    string
	delete bool
}

func NewDockerImage(tag string) *DockerImage {
	return &DockerImage{
		tag:    tag,
		delete: false,
	}
}

func BuildImage(path string) (*DockerImage, error) {
	imageTag := strings.ToLower(domain.RandSeq(10))
	cmd := exec.Command("docker", "build", path, "-t", imageTag)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(output))
		return nil, err
	}

	return &DockerImage{
		tag:    imageTag,
		delete: true,
	}, nil
}

func (i *DockerImage) GetTag() string {
	return i.tag
}

func (i *DockerImage) Cleanup() error {
	if !i.delete {
		return nil
	}

	cmd := exec.Command("docker", "image", "rm", i.tag)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(output))
		return err
	}

	return nil
}
