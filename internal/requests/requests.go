package requests

type MemberUpdateRequest struct {
	Name     string `json:"name" binding:"" example:"test"`
	Email    string `json:"email" binding:"email" example:"example@example.com"`
	Birthday string `json:"birthday" example:"2021-01-01"`
	Gender   int32  `json:"gender" binding:"" example:"1"`
}

type LoginRequeset struct {
	Username string `json:"username" binding:"required" message:"username is required" example:"test"`
	Password string `json:"password" binding:"required" message:"password is required" example:"123456"`
}

type RoleCreateRequest struct {
	Title    string `json:"title" binding:"required" example:"role title"`
	RoleType string `json:"role_type" binding:"required" example:"role type"`
	Image    string `json:"image" binding:"required" example:"role.png"`
}

type RoleUpdateRequest struct {
	Id       int32  `json:"id" binding:"required" example:"1"`
	Title    string `json:"title" example:"role title update"`
	RoleType string `json:"role_type" example:"role type update"`
	Image    string `json:"image" example:"role_update.png"`
}
type RoleDeleteRequest struct {
	Id int32 `json:"id" binding:"required" example:"1"`
}

type OrganizationCreateRequest struct {
	Title          string `json:"title" binding:"required" example:"organization title"`
	Order          int32  `json:"order" binding:"required" example:"1"`
	SourceLanguage string `json:"source_language" binding:"required" example:"en"`
	TargetLanguage string `json:"target_language" binding:"required" example:"zh"`
	Publish        int32  `json:"publish" binding:"required" example:"1"`
}

type OrganizationUpdateRequest struct {
	Id             string `json:"id" binding:"required" example:"1"`
	Title          string `json:"title" example:"organization title update"`
	Order          int32  `json:"order" example:"1"`
	SourceLanguage string `json:"source_language" example:"en"`
	TargetLanguage string `json:"target_language" example:"zh"`
	Publish        int32  `json:"publish" example:"1"`
}

type OrganizationDeleteRequest struct {
	Id string `json:"id" binding:"required" example:"1"`
}

type AssignRoleToMemberRequest struct {
	MemberId       string `json:"member_id" binding:"required" example:"1"`
	RoleId         int32  `json:"role_id" binding:"required" example:"1"`
	OrganizationId string `json:"organization_id" binding:"required" example:"1"`
}

type CourseCreateRequest struct {
	Title          string `json:"title" binding:"required" example:"course title"`
	OrganizationId string `json:"organization_id" binding:"required" example:"1"`
	Order          int32  `json:"order" binding:"required" example:"1"`
	Publish        int32  `json:"publish" binding:"required" example:"1"`
}

type CourseUpdateReqeust struct {
	Id             string `json:"id" binding:"required" example:"1"`
	Title          string `json:"title" example:"course title update"`
	OrganizationId string `json:"organization_id" example:"1"`
	Order          int32  `json:"order" example:"1"`
	Publish        int32  `json:"publish" example:"1"`
}

type CourseDeleteReqeust struct {
	Id string `json:"id" binding:"required" example:"1"`
}
type AssignCourseRequest struct {
	MemberId       string `json:"member_id" binding:"required" example:"1"`
	CourseId       string `json:"course_id" binding:"required" example:"1"`
	OrganizationId string `json:"organization_id" binding:"required" example:"1"`
	RoleId         int32  `json:"role_id" binding:"required" example:"1"`
}

type UnitCreateRequest struct {
	Title    string `json:"title" binding:"required" example:"unit title"`
	CourseId string `json:"course_id" binding:"required" example:"1"`
	Order    int32  `json:"order" binding:"required" example:"1"`
	Publish  int32  `json:"publish" binding:"required" example:"1"`
}

type UnitUpdateRequest struct {
	Id       string `json:"id" binding:"required" example:"1"`
	Title    string `json:"title" example:"unit title update"`
	CourseId string `json:"course_id" example:"1"`
	Order    int32  `json:"order" example:"1"`
	Publish  int32  `json:"publish" example:"1"`
}

type UnitDeleteRequest struct {
	Id string `json:"id" binding:"required" example:"1"`
}

type AssignUnitPermissionRequest struct {
	MemberId string `json:"member_id" binding:"required" example:"1"`
	CourseId string `json:"course_id" binding:"required" example:"1"`
	UnitId   string `json:"unit_id" binding:"required" example:"1"`
	RoleId   int32  `json:"role_id" binding:"required" example:"1"`
}

type WordCreateRequest struct {
	UnitId        string `json:"unit_id" binding:"required" example:"1"`
	Word          string `json:"word" binding:"required" example:"word"`
	Definition    string `json:"definition" binding:"required" example:"definition"`
	Image         string `json:"image" example:"image"`
	Pronunciation string `json:"pronunciation" example:"pronunciation"`
	Description   string `json:"description" example:"description"`
	Comment       string `json:"comment" example:"comment"`
	Order         int32  `json:"order" binding:"required" example:"1"`
}

type WordUpdateRequest struct {
	Id            string `json:"id" binding:"required" example:"1"`
	UnitId        string `json:"unit_id" binding:"required" example:"1"`
	Word          string `json:"word" binding:"required" example:"word"`
	Definition    string `json:"definition" binding:"required" example:"definition"`
	Image         string `json:"image" example:"image"`
	Pronunciation string `json:"pronunciation" example:"pronunciation"`
	Description   string `json:"description" example:"description"`
	Comment       string `json:"comment" example:"comment"`
	Order         int32  `json:"order" binding:"required" example:"1"`
}

type WordDeleteRequest struct {
	Id string `json:"id" binding:"required" example:"1"`
}

// Quiz - Quiz.Type will be 1
type QuizCreateRequest struct {
	QuestionTypes []string `json:"question_type" binding:"required" example:"multiple_choice, true_false, full_in_blank"`
	Topic         int32    `json:"topic" binding:"required" example:"1"`
	ExamDate      string   `json:"exam_date" example:"30"`
	// Info
	QuizCount      int32  `json:"quiz_count" binding:"required" example:"10"`
	OrganizationId string `json:"organization_id" binding:"required" example:"1"`
	CourseId       string `json:"course_id" binding:"required" example:"1"`
	UnitId         string `json:"unit_id" example:"1"`
	// --/info
	MembersId []string `json:"members_id" example:"1, 2, 3"`
}

// Challenge - Quiz.Type will be 2
type QuizChallengesRequest struct {
	QuestionType []string `json:"question_type" binding:"required" example:"multiple_choice, true_false, full_in_blank"`
	Topic        int32    `json:"topic" binding:"required" example:"1"`
	// Info
	QuizDays       int32  `json:"quiz_days" bindding:"required" example:"30"`
	RetryTimes     int32  `json:"retry_times" binding:"required" example:"2"`
	OrganizationId string `json:"organization_id" binding:"required" example:"1"`
	CourseId       string `json:"course_id" binding:"required" example:"1"`
	UnitId         string `json:"unit_id" example:"1"`
	// --/info
	MembersId []string `json:"members_id" example:"1,2,3"`
}

type QuizListRequest struct {
	Page int32 `json:"page" example:"1"`
}

type QuizUpdateQuizAnswerRecordRequest struct {
	QuizId         string `json:"quiz_id" binding:"required" example:"1"`
	AnswerQuestion string `json:"answer_question" binding:"required" example:"{}"`
}
