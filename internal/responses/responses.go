package responses

import (
	"time"
)

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

type QuizListResponse struct {
	QuizList []QuizWithAnswers
	Total    int32 `json:"total" example:"100"`
}

type QuizWithAnswer struct {
	QuizId                    string    `json:"quiz_id" example:"1"`
	QuizAnswerRecordId        string    `json:"quiz_answer_record_id" example:"1"`
	CreaterID                 string    `json:"creater_id" example:"1"`
	QuestionType              string    `json:"question_type" example:"mutiple_choice"`
	Topic                     int32     `json:"topic" example:"1"`
	Type                      int32     `json:"type" example:"1"`
	Info                      string    `json:"info" example:"{}"`
	Content                   string    `json:"content" example:"{}"`
	AnswerQuestion            string    `json:"answer_question" example:"{}"`
	Status                    int32     `json:"status" example:"1"`
	DueDate                   time.Time `json:"due_date" example:"2021-01-01 00:00:00"`
	FailedAnswerCount         int32     `json:"failed_answer_count" example:"10"`
	TotalQuestionCount        int32     `json:"total_question_count" example:"20"`
	FailedLogs                string    `json:"failed_logs" example:"{}"`
	Scope                     int32     `json:"scope" example:"50"`
	QuizAnswerRecordUpdatedAt time.Time `json:"quiz_answer_record_updated_at" example:"2021-01-01 00:00:00"`
}

type QuizWithAnswers struct {
	QuizId             string    `json:"quiz_id" example:"1"`
	QuizAnswerRecordId string    `json:"quiz_answer_record_id" example:"1"`
	CreaterID          string    `json:"creater_id" example:"1"`
	QuestionType       string    `json:"question_type" example:"mutiple_choice"`
	Topic              int32     `json:"topic" example:"1"`
	Type               int32     `json:"type" example:"1"`
	Info               string    `json:"info" example:"{}"`
	Status             int32     `json:"status" example:"1"`
	DueDate            time.Time `json:"due_date" example:"2021-01-01 00:00:00"`
	FailedAnswerCount  int32     `json:"failed_answer_count" example:"10"`
	TotalQuestionCount int32     `json:"total_question_count" example:"20"`
	FailedLogs         string    `json:"failed_logs" example:"{}"`
	Scope              int32     `json:"scope" example:"50"`
}
