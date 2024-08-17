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
	"Role":                   &Role{},
	"Member":                 &Member{},
}
