package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkOrganizationPermissionAllow(t *testing.T) {
	sampleAllowEditorRoles := map[int32]string{
		1: "admin",
		2: "editor",
	}
	for k, _ := range sampleAllowEditorRoles {
		err := checkRoleWithEditorPermission(k)
		assert.Nil(t, err)
	}
}

func Test_checkOrganizationPermissionDisallow(t *testing.T) {
	sampleDisallowedRoles := map[int32]string{
		3: "viewer",
		4: "guest",
		5: "member",
	}
	for k, _ := range sampleDisallowedRoles {
		err := checkRoleWithEditorPermission(k)
		assert.Equal(t, "permission denied", err.Error())
	}
}
