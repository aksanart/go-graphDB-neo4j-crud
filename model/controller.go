package model

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type PersonInterface interface {
	FindAll() (result interface{}, err error)
	Follow(name, following_name string) (result interface{}, err error)
	UnFollow(name, following_name string) (result interface{}, err error)
	AddPerson(name, age string) (result interface{}, err error)
}
type PersonDB struct {
	Name interface{} `json:"name"`
	Age  interface{} `json:"age"`
}

func (p PersonDB) FindAll() (results interface{}, err error) {
	session := Neo4jConfig.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: DatabaseNEO4J,
	})
	defer session.Close()
	dataResults, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`MATCH (n:Person) RETURN n.name as name,n.age as age `,
			map[string]interface{}{})
		if err != nil {
			return nil, err
		}
		var results []PersonDB
		for records.Next() {
			record := records.Record()
			name, _ := record.Get("name")
			age, _ := record.Get("age")
			results = append(results, PersonDB{
				Name: name.(string),
				Age:  age.(string),
			})
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return dataResults, nil
}

func (p PersonDB) Follow(name, following_name string) (results interface{}, err error) {
	session := Neo4jConfig.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: DatabaseNEO4J,
	})
	defer session.Close()
	dataResults, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`MATCH
			(a:Person),
			(b:Person)
		  WHERE a.name = $name AND b.name = $following_name
		  CREATE (a)-[r:Follow]->(b)`,
			map[string]interface{}{"name": name, "following_name": following_name})
		if err != nil {
			return nil, err
		}
		if records.Next() {
			return records.Record().Values[0], nil
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return dataResults, nil
}

func (p PersonDB) UnFollow(name, following_name string) (results interface{}, err error) {
	session := Neo4jConfig.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: DatabaseNEO4J,
	})
	defer session.Close()
	dataResults, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`Match(n:Person{name:$name})-[r:Follow]->(x:Person{name:$following_name})
			DELETE r`,
			map[string]interface{}{"name": name, "following_name": following_name})
		if err != nil {
			return nil, err
		}
		if records.Next() {
			return records.Record().Values[0], nil
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return dataResults, nil
}

func (p PersonDB) AddPerson(name, age string) (results interface{}, err error) {
	session := Neo4jConfig.Driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: DatabaseNEO4J,
	})
	defer session.Close()
	dataResults, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			`Create(n:Person{name:$name,age:$age})Return n`,
			map[string]interface{}{"name": name, "age": age})
		if err != nil {
			return nil, err
		}
		if records.Next() {
			return records.Record().Values[0], nil
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return dataResults, nil
}
