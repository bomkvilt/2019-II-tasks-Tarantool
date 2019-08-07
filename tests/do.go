package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sln/tests/models"
)

func (gn *Generator) checkCode(resp *http.Response, code int) bool {
	if resp.StatusCode == code {
		return true
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		gn.panicIf("body read", err)
	} else {
		_, fn, line, _ := runtime.Caller(2)
		fmt.Printf("unexpected code (%s:%d): \n{\ncode: %d (whanted: %d),\nbody: %s\n}\n",
			fn, line,
			resp.StatusCode,
			code,
			data,
		)
	}
	return false
}

func (gn *Generator) doGet(path string, code int, exp *models.Object) bool {
	resp := gn.get(path)
	if !gn.checkCode(resp, code) {
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
	return gn.checkCode(gn.post(path, msg), code)
}

func (gn *Generator) doPut(path string, code int, msg *models.Message) bool {
	return gn.checkCode(gn.put(path, msg), code)
}

func (gn *Generator) doDelete(path string, code int) bool {
	return gn.checkCode(gn.delete(path), code)
}
