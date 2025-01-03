package requests

type GenerateByInterests struct {
	Condition string   `validate:"required,max=2000"`
	Interests []string `validate:"required,min=0,max=20,dive,max=100"`
}

type GenerateByNoInterests struct {
	Condition string `validate:"required,max=2000"`
}

type GenerateAnswer struct {
	Condition string `validate:"required,max=2000"`
}
