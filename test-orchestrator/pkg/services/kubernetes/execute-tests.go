package kubernetes

import (
	"log"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/project"
)

func ExecuteTests(project *project.Project, volume *PersistantVolume, tests []string) error {
	log.Println("Running a few tests...")

	cmd, err := project.GetTestCommand(tests[0:2])

	if err != nil {
		return err
	}

	cmd.Dir = volume.GetPath()

	if output, err := cmd.Output(); err != nil {
		return err
	} else {
		log.Println(string(output))
	}

	return nil
}
