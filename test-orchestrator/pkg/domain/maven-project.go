package domain

import (
	"fmt"
	"strings"
)

type MavenProject struct {
	projectUrl string
	testModule string
	workspace  Workspace
	tests      *TestSet
}

func CreateMavenProject(projectUrl string, workspace Workspace) (Project, error) {
	return CreateMavenProjectWithTestModule(projectUrl, ".", workspace)
}

func CreateMavenProjectWithTestModule(projectUrl string, testModule string, workspace Workspace) (Project, error) {
	if output, err := workspace.Run(setupCommand(projectUrl, testModule)); err != nil {
		return nil, err
	} else {
		return &MavenProject{
			projectUrl: projectUrl,
			testModule: testModule,
			workspace:  workspace,
			tests:      NewTestSet(strings.Split(output, "\n")),
		}, nil
	}

}

func setupCommand(cloneUrl string, testModule string) string {
	return fmt.Sprintf(
		"git clone %s . 1>&2 && mvn clean install -DskipTests -DskipIT 1>&2 && find %s -type f -name *Test*.class -exec basename -s .class {} \\;",
		cloneUrl,
		testModule,
	)
}

func (ts *MavenProject) GetTestCommands(n int) []string {
	commands := []string{}

	for _, test := range ts.tests.Split(n) {
		commands = append(commands, fmt.Sprintf("cd %s && mvn test -DskipIT -Dtest=%s", ts.testModule, strings.Join(test, ",")))
	}

	return commands
}

func (ts *MavenProject) RunTestsParallelWithConfig(n int, config WorkspaceConfig) error {
	return ts.workspace.RunParallelWithConfig(ts.GetTestCommands(n), config)
}
