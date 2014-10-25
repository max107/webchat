package main

import (
	"encoding/json"
	"time"
)

var lastMessages []Message

type Message struct {
	Id        int64     `json:"id"`
	From      string    `json:"from"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func (msg Message) ToJson() []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return data
}
