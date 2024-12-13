package responses

import "time"

type WordResponse struct {
	Id            string `json:"id"`
	UnitId        string `json:"unit_id"`
	Word          string `json:"word"`
	Definition    string `json:"definition"`
	Image         string `json:"image"`
	Pronunciation string `json:"pronunciation"`
	Description   string `json:"description"`
	Comment       string `json:"comment"`
	Order         int32  `json:"order"`
}
type UnitResponse struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	CourseId  string    `json:"course_id"`
	Order     int32     `json:"order"`
	Publish   int32     `json:"publish"`
	CreaterId string    `json:"creator_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrganizationResponse struct {
	Id             string    `json:"id"`
	Title          string    `json:"title"`
	Order          int32     `json:"order"`
	SourceLanguage string    `json:"source_language"`
	TargetLanguage string    `json:"target_language"`
	Publish        int32     `json:"publish"`
	CreaterId      string    `json:"creater_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type MemberinfoResponse struct {
	ID        string    `json:"id" example:"123456"`
	Name      string    `json:"name" example:"test"`
	Birthday  string    `json:"birthday" example:"2021-01-01"`
	Gender    int32     `json:"gender" example:"1"`
	Email     string    `json:"email" example:"example@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2021-01-01 00:00:00"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-01-01 00:00:00"`
}

type GetMemberResonse struct {
	Code int `json:"code" example:"0"`
	Data MemberinfoResponse
	Msg  string `json:"msg" example:"success"`
}

type CourseResponse struct {
	Id             string    `json:"id"`
	Title          string    `json:"title"`
	OrganizationId string    `json:"organization_id"`
	Order          int32     `json:"order"`
	Publish        int32     `json:"publish"`
	CreaterId      string    `json:"creator_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TokenResponse struct {
	Token string `json:"token" example:"jwt token"`
}
type LoginResonse struct {
	Code int `json:"code" example:"0"`
	Data TokenResponse
	Msg  string `json:"msg" example:"success"`
}
