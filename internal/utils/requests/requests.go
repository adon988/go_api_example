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
