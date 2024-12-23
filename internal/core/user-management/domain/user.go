package domain

type User struct {
	ID       string `form:"id" biding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}
