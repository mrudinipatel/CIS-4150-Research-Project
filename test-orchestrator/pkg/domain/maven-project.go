package domain

import (
	"fmt"
	"strings"
)

type MavenProject struct {
	projectUrl string
}

func CreateMavenProject(projectUrl string) Project {
	return &MavenProject{
		projectUrl: projectUrl,
	}
}

func (ts *MavenProject) GetCloneUrl() string {
	return ts.projectUrl
}

func (ts *MavenProject) GetTestCommand(testsuites []string) string {
	return fmt.Sprintf("mvn test -DskipIT -Dtest=%s -Dmaven.repo.local=./.m2", strings.Join(testsuites, ","))
}

func (ts *MavenProject) GetSetupCommand() string {
	return "mvn clean install -DskipTests -DskipIT -Dmaven.repo.local=./.m2"
}

func (ts *MavenProject) GetTestFilter() string {
	return "*Test*.java"
}
