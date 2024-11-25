package dockerexecutor

import (
	"log"
	"os/exec"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

func BuildImage(path string) (string, error) {
	imageTag := strings.ToLower(domain.RandSeq(10))
	cmd := exec.Command("docker", "build", path, "-t", imageTag)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(output))
		return "", err
	}

	return imageTag, nil
}

func DeleteImage(image string) error {
	cmd := exec.Command("docker", "image", "rm", image)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Println(string(output))
		return err
	}

	return nil
}
