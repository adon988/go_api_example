package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkOrganizationPermissionAllow(t *testing.T) {
	sampleAllowEditorRoles := map[int32]string{
		1: "admin",
		2: "editor",
	}
	for k := range sampleAllowEditorRoles {
		err := CheckRoleWithEditorPermission(k)
		assert.Nil(t, err)
	}
}

func Test_checkOrganizationPermissionDisallow(t *testing.T) {
	sampleDisallowedRoles := map[int32]string{
		3: "viewer",
		4: "guest",
		5: "member",
	}
	for k := range sampleDisallowedRoles {
		err := CheckRoleWithEditorPermission(k)
		assert.Equal(t, "permission denied", err.Error())
	}
}
