package kubernetes

import (
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/project"
)

// TODO: This should happen in the Kubernetes Job to a Persistant Volume
func CloneTestProject(project *project.Project, volume *PersistantVolume) ([]string, error) {
	log.Println("Setting up project...")
	cmd := exec.Command("git", "clone", project.GetCloneUrl(), volume.GetPath())

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	setup, _ := project.GetSetupCommand()

	setup.Dir = volume.GetPath()

	if err := setup.Run(); err != nil {
		return nil, err
	}

	// Find Test Files
	log.Println("Discovering test files...")

	getTests := exec.Command("find", ".", "-type", "f", "-name", "*Test*.java")
	getTests.Dir = volume.GetPath()

	tests := []string{}

	if output, err := getTests.Output(); err != nil {
		return nil, err
	} else {
		for _, testPath := range strings.Split(string(output), "\n") {
			filename := path.Base(testPath)
			extension := path.Ext(filename)
			tests = append(tests, filename[0:len(filename)-len(extension)])
		}
	}

	return tests, nil
}
