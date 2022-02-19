package service

import "github.com/aksanart/go-graphDB-neo4j-crud/model"

type ControllerInterface interface {
	model.PersonInterface
	VendorInterface
}
type VendorInterface interface {
}
