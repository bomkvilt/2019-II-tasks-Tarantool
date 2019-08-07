package tests

import (
	"encoding/json"
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

// -----------| no key
func (gn *Generator) scriptNoKey() bool {
	var (
		key   = gn.genKey()
		kpath = gn.conf.GetURL(key)
		value = &models.Object{
			Name: "Rob",
			Data: "Intel",
		}
	)
	if !gn.doGet(kpath, 404, nil) {
		return false
	}
	if !gn.doPut(kpath, 404, &models.Message{
		Value: value,
	}) {
		return false
	}
	if !gn.doDelete(kpath, 404) {
		return false
	}
	return true
}

// -----------| key already exists
func (gn *Generator) scriptKeyAlreadyExists() bool {
	var (
		key   = gn.genKey()
		rpath = gn.conf.GetURL()
		kpath = gn.conf.GetURL(key)
		value = &models.Object{
			Name: "Rob",
			Data: "Intel",
		}
	)
	if !gn.doGet(kpath, 404, nil) {
		return false
	}
	if !gn.doPost(rpath, 200, &models.Message{
		Key:   key,
		Value: value,
	}) {
		return false
	}
	if !gn.doPost(rpath, 409, &models.Message{
		Key:   key,
		Value: value,
	}) {
		return false
	}
	return true
}

// -----------| valid value
func (gn *Generator) scriptValidValue() bool {
	for _, value := range []string{
		`{"a": 10}`, // object
		`[10, 6]`,   // array
		`"aksnc"`,   // string
		`1000000`,   // number
		`true`,
		`false`,
		`{ "a": null }`,
	} {
		var (
			key   = gn.genKey()
			rpath = gn.conf.GetURL()
			kpath = gn.conf.GetURL(key)
		)

		rawValue := json.RawMessage(value)
		if !gn.doPost(rpath, 200, &models.Message{
			Key:   key,
			Value: rawValue,
		}) {
			return false
		}

		if !gn.doPut(kpath, 200, &models.Message{
			Value: rawValue,
		}) {
			return false
		}
	}
	return true
}

// -----------| invalid value
func (gn *Generator) scriptInvalidValue() bool {
	for _, value := range []string{
		`{a: 10}`,    // key without quotes
		`{"a": 10,}`, // extra quote
		`{"a": 10`,   // none closed object
		`[10, 6`,     // none closed array
		`no_quotes`,  // string without quotes
	} {
		var (
			key   = gn.genKey()
			rpath = gn.conf.GetURL()
			kpath = gn.conf.GetURL(key)
		)

		rawValue := json.RawMessage(value)
		if !gn.doPost(rpath, 400, &models.Message{
			Key:   key,
			Value: rawValue,
		}) {
			return false
		}

		if !gn.doPost(rpath, 200, &models.Message{
			Key:   key,
			Value: "",
		}) {
			return false
		}

		if !gn.doPut(kpath, 400, &models.Message{
			Value: rawValue,
		}) {
			return false
		}
	}
	return true
}
