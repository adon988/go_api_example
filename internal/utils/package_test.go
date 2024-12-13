package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenSha256IdempotentId(t *testing.T) {
	hashInfos := []string{"memberId123", "info_json"}
	idempotentId := GenSha256IdempotentId(hashInfos)
	assert.Equal(t, "732e6587267a5a1863d225a120c9803c59fab27c8768f1e82466a916d087a876", idempotentId)
}
