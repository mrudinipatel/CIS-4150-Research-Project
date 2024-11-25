package domain

import (
	"fmt"
	"strings"
)

type MavenProject struct {
	projectUrl string
	testModule string
}

func CreateMavenProject(projectUrl string) Project {
	return &MavenProject{
		projectUrl: projectUrl,
		testModule: ".",
	}
}

func CreateMavenProjectWithTestModule(projectUrl string, testModule string) Project {
	return &MavenProject{
		projectUrl,
		testModule,
	}
}

func (ts *MavenProject) GetCloneUrl() string {
	return ts.projectUrl
}

func (ts *MavenProject) GetTestCommand(testsuites []string) string {
	return fmt.Sprintf("cd %s && mvn test -DskipIT -Dtest=%s", ts.testModule, strings.Join(testsuites, ","))
}

func (ts *MavenProject) GetSetupCommand() string {
	return "mvn clean install -DskipTests -DskipIT"
}

func (ts *MavenProject) GetTestFilterCommand() string {
	return fmt.Sprintf("find %s -type f -name *Test*.class -exec basename -s .class {} \\;", ts.testModule)
}
