package controllers

type MemberCreateVerify struct {
	Name     string `form:"name" json:"name" binding:"" example:"test"`
	Age      int64  `form:"age" json:"age" binding:"" example:"18"`
	Email    string `form:"email" json:"email" binding:"email" example:"example@example.com"`
	Birthday string `form:"birthday" json:"birthday" example:"2021-01-01"`
	Gender   int    `json:"gender" binding:"" example:"1"`
}
type MemberUpdateVerify struct {
	Name     string `json:"name" binding:"" example:"test"`
	Age      int64  `json:"age" binding:"" example:"18"`
	Email    string `json:"email" binding:"email" example:"example@example.com"`
	Birthday string `json:"birthday" example:"2021-01-01"`
	Gender   int    `json:"gender" binding:"" example:"1"`
}

type LoginVerify struct {
	Username string `json:"username" binding:"required" message:"username is required" example:"test"`
	Password string `json:"password" binding:"required" message:"password is required" example:"123456"`
}
