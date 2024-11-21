package controllers

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/project"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/services/kubernetes"
)

type TestController struct{}

func (t *TestController) ExecTestSuite(proj *project.Project) error {
	volume, err := kubernetes.CreatePersistantVolume()

	if err != nil {
		return err
	}

	tests, err := kubernetes.CloneTestProject(proj, volume)

	if err != nil {
		return err
	}

	return kubernetes.ExecuteTests(proj, volume, tests)
}
