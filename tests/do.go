package tests

import (
	"io/ioutil"
	"sln/tests/models"
)

func (gn *Generator) doGet(path string, code int, exp *models.Object) bool {
	resp := gn.get(path)
	if resp.StatusCode != code {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	if exp == nil {
		return true
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gn.panicIf("body read", err)
		return false
	}

	tkn := &models.Object{}
	gn.panicIf("unmarshal", tkn.UnmarshalJSON(data))
	return gn.compare(exp, tkn)
}

func (gn *Generator) doPost(path string, code int, msg *models.Message) bool {
	resp := gn.post(path, msg)
	if resp.StatusCode != code {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	return true
}

func (gn *Generator) doPut(path string, code int, msg *models.Message) bool {
	resp := gn.put(path, msg)
	if resp.StatusCode != code {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	return true
}

func (gn *Generator) doDelete(path string, code int) bool {
	resp := gn.delete(path)
	if resp.StatusCode != code {
		gn.unexpectedCode(resp.StatusCode)
		return false
	}
	return true
}
