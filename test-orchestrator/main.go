package main

import (
	"log"

	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/controllers"
	"github.com/D3h4n/CIS-4150-Research-Project/test-orchestrator/pkg/domain/project"
)

func main() {
	tc := controllers.TestController{}

	if err := tc.ExecTestSuite(project.Create("https://github.com/jhy/jsoup.git", project.Maven)); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Success")
	}
}
