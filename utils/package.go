package utils

import (
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
