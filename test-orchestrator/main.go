package main

import (
	"log"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/controllers"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	dockerexecutor "github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/executors/docker-executor"
)

func main() {
	image, err := dockerexecutor.BuildImage(".")

	if err != nil {
		panic(err)
	}

	tc := controllers.TestController{
		Executor: dockerexecutor.Create(
			dockerexecutor.NewContainerConfig(
				image,
				"2g",
				"1",
			),
			5,
		),
	}

	project := domain.CreateMavenProject("https://github.com/jhy/jsoup.git")

	if err := tc.ExecTestSuite(project); err != nil {
		log.Fatal(err)
	}

	dockerexecutor.DeleteImage(image)
}
