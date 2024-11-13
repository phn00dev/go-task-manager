package dto

type CreateAdminRequest struct {
	Firstname       string `json:"firstname" binding:"required,min=2,max=100"`
	Lastname        string `json:"lastname" binding:"required,min=2,max=100"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type UpdateAdminRequest struct {
	Firstname string `json:"firstname" binding:"omitempty,min=2,max=100"`
	Lastname  string `json:"lastname" binding:"omitempty,min=2,max=100"`
	Email     string `json:"email" binding:"omitempty,email"`
}

type UpdateAdminPasswordRequest struct {
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}
