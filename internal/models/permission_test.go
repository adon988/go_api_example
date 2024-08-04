package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionRole(t *testing.T) {
	roles := []string{"admin", "editor", "guest"}
	for _, role := range roles {
		_, ok := PermissionRoleEnum[role]
		assert.True(t, ok)
	}
}

func TestPermissionRoleNotExist(t *testing.T) {
	Role := "not_exist"
	_, ok := PermissionRoleEnum[Role]

	assert.False(t, ok)
}
