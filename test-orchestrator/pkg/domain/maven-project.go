package domain

import (
	"fmt"
	"os/exec"
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

func (ts *MavenProject) GetTestCommand(testsuites []string) *exec.Cmd {
	return exec.Command("mvn", "test", "-DskipIT", fmt.Sprintf("-Dtest=%s", strings.Join(testsuites, ",")))
}

func (ts *MavenProject) GetSetupCommand() *exec.Cmd {
	return exec.Command("mvn", "clean", "install", "-DskipTests", "-DskipIT")
}
