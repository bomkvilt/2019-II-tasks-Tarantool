package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"sln/tests/models"

	"github.com/google/uuid"
)

//go:generate easyjson -output_filename ./models/jsons.gen.go ./models
//go:generate easyjson -output_filename jsons.gen.go .

// --------------------------| generator |-------------------------- \

// Generator generates test cases for required key-value storage
type Generator struct {
	scripts []func() bool
	conf    Config
}

// NewGenerator creates a new test generator
func NewGenerator(conf Config) *Generator {
	gn := &Generator{
		scripts: []func() bool{},
		conf:    conf,
	}
	gn.addScript(gn.scriptSimpleInsertion)
	return gn
}

// -----------|

// PlayScripts plays scripted test cases
func (gn *Generator) PlayScripts() {
	fmt.Println("Playing scripts...")
	for _, script := range gn.scripts {
		if !script() {
			break
		}
	}
	fmt.Println("Scripts finished sucessfully!")
}

func (gn *Generator) addScript(script func() bool) {
	gn.scripts = append(gn.scripts, script)
}

// -----------| helpers

func (gn *Generator) genKey() string {
	key := uuid.New().String()
	key = key[0:30]
	return key
}

func (gn *Generator) get(path string) *http.Response {
	resp, err := http.Get(path)
	gn.panicIf("get", err)
	return resp
}

func (gn *Generator) post(path string, msg *models.Message) *http.Response {
	resp, err := http.Post(path, "application/json", msg.ToReader())
	gn.panicIf("post", err)
	return resp
}

func (gn *Generator) put(path string, msg *models.Message) *http.Response {
	req, err := http.NewRequest("PUT", path, msg.ToReader())
	gn.panicIf("new request", err)

	client := http.Client{}
	resp, err := client.Do(req)
	gn.panicIf("put", err)
	return resp
}

func (gn *Generator) delete(path string) *http.Response {
	req, err := http.NewRequest("DELETE", path, nil)
	gn.panicIf("new request", err)

	client := http.Client{}
	resp, err := client.Do(req)
	gn.panicIf("delete", err)
	return resp
}

// -----------|

func (gn *Generator) unexpectedCode(code int) {
	fmt.Printf("unexpected code:. %v", code)
}

func (gn *Generator) panicIf(tag string, err error) {
	if err == nil {
		return
	}
	err = fmt.Errorf("unexpected %s error:. %v", tag, err)
	panic(err)
}

func (gn *Generator) compare(exp, tkn interface{}) bool {
	if !reflect.DeepEqual(exp, tkn) {
		fmt.Printf("expected:\n%+v\ntaken:\n%+v\n", exp, tkn)
		return false
	}
	return true
}
