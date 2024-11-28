package domain

type Project interface {
	GetName() string
	GetTestCommands(n int) []string
	RunTestsParallelWithConfig(int, WorkspaceConfig) error
}

type Workspace interface {
	Cleanup() error
	Run(string) (string, error)
	RunWithConfig(string, WorkspaceConfig) (string, error)
	RunParallelWithConfig([]string, WorkspaceConfig) error
}

type WorkspaceConfig interface{}
