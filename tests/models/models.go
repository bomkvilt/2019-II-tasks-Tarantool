package models

import (
	"bytes"
	"io"
)

// --------------------------| message |-------------------------- \

// Message is a storage request payload
// easyjson:json
type Message struct {
	Key   string      `json:"key,omnitempty"`
	Value interface{} `json:"value"`
}

// ToReader marshales self to JSON and places to to a reader
func (msg *Message) ToReader() io.Reader {
	if msg == nil {
		return nil
	}

	data, err := msg.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(data)
}

// --------------------------| value types |-------------------------- \

// Object is a generic message value
// easyjson:json
type Object struct {
	Name string      `json:"name"`
	Data interface{} `json:"data,omnitempty"`
}
