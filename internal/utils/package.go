package utils

import (
	"encoding/json"

	"github.com/bwmarrin/snowflake"
)

func GenId() (string, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return "", err
	}
	id := node.Generate().String()

	return id, nil
}

func MarshalJSONToRaw(v any) *json.RawMessage {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	rawMsg := json.RawMessage(raw)
	return &rawMsg
}
