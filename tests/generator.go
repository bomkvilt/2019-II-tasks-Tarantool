package tests

import (
	"fmt"
	"net/http"
	"reflect"
	"sln/tests/models"
	"strings"

	"github.com/google/uuid"
)

//go:generate easyjson -output_filename ./models/jsons.gen.go ./models

// --------------------------| generator |-------------------------- \

// Generator is a test script launcher
type Generator struct {
	scripts []func() bool
	conf    Config
}

// NewGenerator creates a new generator
func NewGenerator(conf Config) *Generator {
	gn := &Generator{
		scripts: []func() bool{},
		conf:    conf,
	}
	gn.addScript(gn.scriptSimpleInsertion)
	gn.addScript(gn.scriptNoKey)
	gn.addScript(gn.scriptKeyAlreadyExists)
	gn.addScript(gn.scriptInvalidValue)
	gn.addScript(gn.scriptValidValue)
	return gn
}

// -----------|

// PlayScripts plays scripted test cases
func (gn *Generator) PlayScripts() bool {
	fmt.Println("Playing scripts...")
	for _, script := range gn.scripts {
		if script() {
			continue
		}
		fmt.Println("Error occurred!")
		return false
	}
	fmt.Println("Scripts finished successfully!")
	return true
}

func (gn *Generator) addScript(script func() bool) {
	gn.scripts = append(gn.scripts, script)
}

// -----------| helpers

// genKey generates a rundom key
func (gn *Generator) genKey() string {
	key := uuid.New().String()
	key = key[0:30]
	return key
}

func (gn *Generator) get(path string) *http.Response {
	fmt.Printf("get '%s'\n", path)
	resp, err := http.Get(path)
	gn.panicIf("get", err)
	return resp
}

func (gn *Generator) post(path string, msg *models.Message) *http.Response {
	fmt.Printf("post to '%s' message '%+v'\n", path, msg)
	resp, err := http.Post(path, "application/json", msg.ToReader())
	gn.panicIf("post", err)
	return resp
}

func (gn *Generator) put(path string, msg *models.Message) *http.Response {
	return gn.extraMethod("put", path, msg)
}

func (gn *Generator) delete(path string) *http.Response {
	return gn.extraMethod("delete", path, nil)
}

func (gn *Generator) extraMethod(method, path string, msg *models.Message) *http.Response {
	method = strings.ToUpper(method)
	tag := strings.ToLower(method)

	fmt.Printf("%s path='%s' message='%v'\n", tag, path, msg)
	req, err := http.NewRequest(method, path, msg.ToReader())
	gn.panicIf("new request", err)

	client := http.Client{}
	resp, err := client.Do(req)
	gn.panicIf(tag, err)
	return resp
}

// -----------|

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
