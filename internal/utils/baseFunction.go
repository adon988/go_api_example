package utils

import "fmt"

var allowEditorRoles = map[int32]string{
	1: "admin",
	2: "editor",
}

func CheckRoleWithEditorPermission(RoleId int32) error {
	if _, ok := allowEditorRoles[RoleId]; !ok {
		return fmt.Errorf("permission denied")
	}
	return nil
}

var allowQuestionTypes = map[string]bool{
	"multiple_choice": true,
	"true_false":      true,
	"full_in_blank":   true,
}

func CheckQuestionTypes(questionType []string) error {
	for _, qt := range questionType {
		if _, ok := allowQuestionTypes[qt]; !ok {
			return fmt.Errorf("invalid question type")
		}
	}
	return nil
}
