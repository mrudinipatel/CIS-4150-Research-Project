package controllers

import (
	"time"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
)

type TestController struct {
	Executor domain.TestExecutor
}

func (t *TestController) ExecTestSuite(project domain.Project) (time.Duration, error) {
	workspace, err := t.Executor.CreateWorkspace()

	if err != nil {
		return -1, err
	}

	defer t.Executor.CleanupWorkspace(workspace)

	tests, err := t.Executor.SetupProject(project, workspace)

	if err != nil {
		return -1, err
	}

	curr := time.Now()

	if err = t.Executor.ExecuteTests(project, workspace, tests); err != nil {
		return -1, err
	}

	return time.Since(curr), nil
}
