package domain

import (
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/testset"
)

type Workspace interface {
	GetPath() string
	GetName() string
}

type TestExecutor interface {
	CreateWorkspace() (Workspace, error)
	CleanupWorkspace(Workspace) error
	SetupProject(Project, Workspace) (*testset.TestSet, error)
	ExecuteTests(Project, Workspace, *testset.TestSet) error
}

type Project interface {
	GetCloneUrl() string
	GetTestCommand(testsuites []string) string
	GetSetupCommand() string
	GetTestFilter() string
}
