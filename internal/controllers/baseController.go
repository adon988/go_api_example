package controllers

import "fmt"

var allowEditorRoles = map[int32]string{
	1: "admin",
	2: "editor",
}

func checkRoleWithEditorPermission(RoleId int32) error {
	if _, ok := allowEditorRoles[RoleId]; !ok {
		return fmt.Errorf("permission denied")
	}
	return nil
}
