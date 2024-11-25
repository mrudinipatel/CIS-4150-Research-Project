package main

import (
	"fmt"
	"log"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/controllers"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	dockerexecutor "github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/executors/docker-executor"
)

func main() {
	image, err := dockerexecutor.BuildImage(".")

	if err != nil {
		log.Panic(err)
	}

	defer dockerexecutor.DeleteImage(image)

	tc := controllers.TestController{
		Executor: dockerexecutor.Create(
			dockerexecutor.NewContainerConfig(
				image,
				"2g",
				"2",
			),
			4,
		),
	}

	project := domain.CreateMavenProject("https://github.com/jhy/jsoup.git")
	// project := domain.CreateMavenProjectWithTestModule("https://github.com/google/guava.git", "guava-tests")

	if duration, err := tc.ExecTestSuite(project); err != nil {
		log.Panic(err)
	} else {
		fmt.Println(duration)
	}
}
