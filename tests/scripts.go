package tests

import (
	"sln/tests/models"
)

// -----------| simple insertion

func (gn *Generator) scriptSimpleInsertion() bool {
	var (
		key   = gn.genKey()
		rpath = gn.conf.GetURL()
		kpath = gn.conf.GetURL(key)
		value = &models.Object{
			Name: "Rob",
			Data: "Intel",
		}
	)
	// check the key not exists
	if !gn.doGet(kpath, 404, nil) {
		return false
	}
	// insert the key
	if !gn.doPost(rpath, 200, &models.Message{
		Key:   key,
		Value: value,
	}) {
		return false
	}
	// get it's value
	if !gn.doGet(kpath, 200, value) {
		return false
	}
	// change stored data
	value.Name = "John"
	if !gn.doPut(kpath, 200, &models.Message{Value: value}) {
		return false
	}
	// check the value
	if !gn.doGet(kpath, 200, value) {
		return false
	}
	// delete the value
	if !gn.doDelete(kpath, 200) {
		return false
	}
	return true
}
