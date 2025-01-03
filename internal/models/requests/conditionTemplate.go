package requests

type CreateConditionTemplate struct {
	Title     string `validate:"required,min=3,max=100"`
	Condition string `validate:"required,max=2000"`
}

type GetConditionTemplate struct {
	ID uint `validate:"required"`
}

type EditConditionTemplate struct {
	ID        uint   `validate:"required"`
	Title     string `validate:"required,min=3,max=100"`
	Condition string `validate:"required,max=2000"`
}

type DeleteConditionTemplate struct {
	ID uint `validate:"required"`
}
