package utils

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"

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

func GenSha256IdempotentId(hashInfos []string) (idempotentId string) {
	hash := sha256.Sum256([]byte(strings.Join(hashInfos, "")))
	idempotentId = fmt.Sprintf("%x", hash)
	return idempotentId
}

func MarshalJSONToRaw(v any) *json.RawMessage {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	rawMsg := json.RawMessage(raw)
	return &rawMsg
}
