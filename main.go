package main

import (
	"fmt"

	"github.com/aksanart/go-graphDB-neo4j-crud/model"
	"github.com/aksanart/go-graphDB-neo4j-crud/routes"
)

func main() {
	fmt.Println("-------- Run --------")
	model.Neo4jConfig.Init()
	defer model.Neo4jConfig.Driver.Close()
	r := routes.SetupRouter()
	r.Run(":1234")
}
