package models

var ModelsMap = map[string]interface{}{
	"Authentication":         &Authentication{},
	"Organization":           &Organization{},
	"OrganizationPermission": &OrganizationPermission{},
	"CoursePermission":       &CoursePermission{},
	"UnitPermission":         &UnitPermission{},
	"Unit":                   &Unit{},
	"Word":                   &Word{},
	"Course":                 &Course{},
	"Member":                 &Member{},
	"Quiz":                   &Quiz{},
	"QuizAnswerRecord":       &QuizAnswerRecord{},
}
