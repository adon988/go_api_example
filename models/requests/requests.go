package requests

type MemberUpdateRequest struct {
	Name     string `json:"name" binding:"" example:"test"`
	Email    string `json:"email" binding:"email" example:"example@example.com"`
	Birthday string `json:"birthday" example:"2021-01-01"`
	Gender   int    `json:"gender" binding:"" example:"1"`
}

type LoginRequeset struct {
	Username string `json:"username" binding:"required" message:"username is required" example:"test"`
	Password string `json:"password" binding:"required" message:"password is required" example:"123456"`
}
