package models

var ModelsMap = map[string]interface{}{
	"Authentication":          &Authentication{},
	"Organization":            &Organization{},
	"OriganizationPermission": &OrganizationPermission{},
	"CoursePermission":        &CoursePermission{},
	"UnitPermission":          &UnitPermission{},
	"Unit":                    &Unit{},
	"Word":                    &Word{},
	"Course":                  &Course{},
	"Role":                    &Role{},
	"Member":                  &Member{},
}
