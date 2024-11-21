package project

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type ProjectType int32

const (
	Maven ProjectType = iota
)

type Project struct {
	projectUrl  string
	projectType ProjectType
}

func Create(projectUrl string, projectType ProjectType) *Project {
	return &Project{
		projectUrl:  projectUrl,
		projectType: projectType,
	}
}

func (ts *Project) GetCloneUrl() string {
	return ts.projectUrl
}

func (ts *Project) GetTestCommand(testsuites []string) (*exec.Cmd, error) {
	switch ts.projectType {
	case Maven:
		return exec.Command("mvn", "test", "-DskipIT", fmt.Sprintf("-Dtest=%s", strings.Join(testsuites, ","))), nil
	default:
		return nil, errors.New("unsupported project type")
	}
}

func (ts *Project) GetSetupCommand() (*exec.Cmd, error) {
	switch ts.projectType {
	case Maven:
		return exec.Command("mvn", "clean", "install", "-DskipTests", "-DskipIT"), nil
	default:
		return nil, errors.New("unsupported project type")
	}
}

func (ts *Project) GetDirName() string {
	return "Blahaj"
}
