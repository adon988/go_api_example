package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermissionRole(t *testing.T) {
	roles := []string{"admin", "editor", "viewer"}
	for _, role := range roles {
		v, ok := PermissionRoleEnum[role]
		if v < 0 || v > 3 {
			assert.Fail(t, fmt.Sprintf("Role %s is not in the range of 0-3", role))
		} else {
			assert.True(t, ok)
		}
	}
}

func TestPermissionRoleNotExist(t *testing.T) {
	Role := "not_exist"
	_, ok := PermissionRoleEnum[Role]
	assert.False(t, ok)
}
