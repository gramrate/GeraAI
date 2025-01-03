package requests

type CreateTask struct {
	Title     string `validate:"required,min=3,max=100"`
	Condition string `validate:"required,max=2000"`
	Answer    string `validate:"required,max=100"`
}

type GetTask struct {
	ID uint `validate:"required"` // ID обязательно
}

type EditTask struct {
	ID        uint   `validate:"required"`
	Title     string `validate:"required,min=3,max=100"`
	Condition string `validate:"required,max=2000"`
	Answer    string `validate:"required,max=100"`
}

type DeleteTask struct {
	ID uint `validate:"required"`
}
