package controllers

import (
	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/project"
)

type TestController struct{}

func (t *TestController) ExecTestSuite(proj *project.Project) error {
	// TODO: This should happen in the Kubernetes Job to a Persistant Volume
	// Clone Repo
	log.Println("Setting up project...")

	targetDir := path.Join("tmp", proj.GetDirName())
	cmd := exec.Command("git", "clone", proj.GetCloneUrl(), targetDir)

	if err := cmd.Run(); err != nil {
		return err
	}

	setup, _ := proj.GetSetupCommand()

	if err := setup.Run(); err != nil {
		return err
	}

	// TODO: This should happen in the Kubernetes Job to a Persistant Volume
	// Find Test Files
	log.Println("Discovering test files...")

	getTests := exec.Command("find", ".", "-type", "f", "-name", "*Test*.java")
	getTests.Dir = targetDir

	tests := []string{}

	if output, err := getTests.Output(); err != nil {
		return err
	} else {
		for _, testPath := range strings.Split(string(output), "\n") {
			filename := path.Base(testPath)
			extension := path.Ext(filename)
			tests = append(tests, filename[0:len(filename)-len(extension)])
		}
	}

	// TODO: Tests should be split and ran as separate K8s jobs across nodes
	// Execute Tests
	log.Println("Running a few tests...")

	cmd, err := proj.GetTestCommand(tests[0:2])

	if err != nil {
		return err
	}

	cmd.Dir = targetDir

	if output, err := cmd.Output(); err != nil {
		return err
	} else {
		log.Println(string(output))
	}

	return nil
}
