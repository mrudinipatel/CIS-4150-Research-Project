package domain

type Workspace interface {
	GetPath() string
	GetName() string
}

type TestExecutor interface {
	CreateWorkspace() (Workspace, error)
	CleanupWorkspace(Workspace) error
	SetupProject(Project, Workspace) (*TestSet, error)
	ExecuteTests(Project, Workspace, *TestSet) error
}

type Project interface {
	GetCloneUrl() string
	GetTestCommand(testsuites []string) string
	GetSetupCommand() string
	GetTestFilterCommand() string
}
