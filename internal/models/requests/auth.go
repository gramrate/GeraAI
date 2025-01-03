package requests

type Login struct {
	Login    string `validate:"required,min=5,max=20,alphanum"`
	Password string `validate:"required,min=8"`
}

type Register struct {
	Login    string `validate:"required,min=5,max=20,alphanum"`
	Username string `validate:"required,max=35"`
	Password string `validate:"required,min=8"`
}
