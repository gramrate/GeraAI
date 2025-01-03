package requests

type CreateInterestsTemplate struct {
	Title     string   `validate:"required,min=3,max=100"`
	Interests []string `validate:"required,dive,max=100"`
}

type GetInterestsTemplate struct {
	ID uint `validate:"required"`
}

type EditInterestsTemplate struct {
	ID        uint     `validate:"required"`
	Title     string   `validate:"required,min=3,max=100"`
	Interests []string `validate:"required,dive,max=100"`
}

type DeleteInterestsTemplate struct {
	ID uint `validate:"required"`
}
