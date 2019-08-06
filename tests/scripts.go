package tests

import (
	"io/ioutil"
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
	resp := gn.get(kpath)
	if resp.StatusCode != 404 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}

	// insert the key
	resp = gn.post(rpath, &models.Message{
		Key:   key,
		Value: value,
	})
	if resp.StatusCode != 200 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}

	// get it's value
	resp = gn.get(kpath)
	if resp.StatusCode != 200 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		stored := models.Object{}
		err = stored.UnmarshalJSON(data)
		gn.panicIf("unmarshal", err)
		gn.compare(value, stored)
	} else {
		gn.panicIf("body read", err)
	}

	// change stored data
	value.Name = "John"
	resp = gn.post(kpath, &models.Message{
		Value: value,
	})
	if resp.StatusCode != 200 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}

	// check the value
	resp = gn.get(kpath)
	if resp.StatusCode != 200 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	if data, err := ioutil.ReadAll(resp.Body); err != nil {
		stored := models.Object{}
		err = stored.UnmarshalJSON(data)
		gn.panicIf("unmarshal", err)
		gn.compare(value, stored)
	} else {
		gn.panicIf("body read", err)
	}

	// delete the value
	resp = gn.delete("kpath")
	if resp.StatusCode != 200 {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}

	return true
}
