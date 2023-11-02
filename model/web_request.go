package model

type RegistrasiRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Username string `validate:"required,min=6,max=125,alphanum" json:"username"`
	Password string `validate:"required,min=6,max=125,alphanum" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,min=6,max=125,alphanum" json:"username"`
	Password string `validate:"required,min=6,max=125" json:"password"`
}

type GetAllTodoRequest struct {
	Userid int32 `validate:"required,numeric" json:"userid"`
}

type AddNewTodoRequest struct {
	Todo      string `validate:"required" json:"todo"`
	Complated bool   `validate:"required,boolean" json:"complated"`
	Userid    int32  `validate:"required,numeric" json:"userid"`
}

type GetorDeleteTodoRequest struct {
	Userid int32 `validate:"required,numeric" json:"userid"`
	ID     int32 `validate:"required,numeric" json:"id"`
}

type UpdateStatusTodoRequest struct {
	ID        int32 `validate:"required,numeric" json:"id"`
	Complated bool  `validate:"boolean" json:"complated"`
	Userid    int32 `validate:"required,numeric" json:"userid"`
}

type GetTodoFilterRequest struct {
	Userid int32 `validate:"required,numeric" json:"userid"`
	Limit  int32 `validate:"numeric" json:"limit"`
	Offset int32 `validate:"numeric" json:"offset"`
}
