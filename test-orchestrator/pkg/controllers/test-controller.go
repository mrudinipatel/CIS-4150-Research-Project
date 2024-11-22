package controllers

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type TestController struct {
	Executor domain.TestExecutor
}

func (t *TestController) ExecTestSuite(project domain.Project) error {
	workspace, err := t.Executor.CreateWorkspace()

	if err != nil {
		return err
	}

	tests, err := t.Executor.SetupProject(project, workspace)

	if err != nil {
		t.Executor.CleanupWorkspace(workspace)
		return err
	}

	if err := t.Executor.ExecuteTests(project, workspace, tests); err != nil {
		t.Executor.CleanupWorkspace(workspace)
		return err
	}

	return t.Executor.CleanupWorkspace(workspace)
}
