package model

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Neo4jModel struct {
	Driver neo4j.Driver
}

func (db *Neo4jModel) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Driver, err = neo4j.NewDriver(os.Getenv("NE04J_HOST"), neo4j.BasicAuth(os.Getenv("NE04J_USER"), os.Getenv("NE04J_PASS"), ""))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connected to: ", os.Getenv("NE04J_HOST"), neo4j.BasicAuth(os.Getenv("NE04J_USER"), os.Getenv("NE04J_PASS"), ""))
	DatabaseNEO4J = os.Getenv("NE04J_DB")
}

var Neo4jConfig = &Neo4jModel{}
var DatabaseNEO4J string
