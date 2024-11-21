package main

import (
	"log"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/controllers"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/services/machine"
)

func main() {
	tc := controllers.TestController{
		Executor: machine.Create(),
	}

	project := domain.CreateMavenProject("https://github.com/jhy/jsoup.git")

	if err := tc.ExecTestSuite(project); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success")
	}
}
