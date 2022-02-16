package request

type User struct {
	Name     string `binding:"required" json:"name"`
	Age      int    `binding:"required" json:"age"`
	Position string `binding:"required" json:"position"`
}

type UpdateUser struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

type UserId struct {
	Id int `binding:"required" form:"id"`
}
