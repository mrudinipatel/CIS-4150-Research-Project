package controllers

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type TestController struct {
	Executor domain.TestExecutor
}

func (t *TestController) ExecTestSuite(proj domain.Project) error {
	volume, err := t.Executor.CreateWorkspace()

	if err != nil {
		return err
	}

	tests, err := t.Executor.SetupProject(proj, volume)

	if err != nil {
		return err
	}

	return t.Executor.ExecuteTests(proj, volume, tests)
}
