package requests

type CreateVariantTemplate struct {
	Title string `validate:"required,min=3,max=30"`
	Tasks string `validate:"required,max=300"`
	Tags  string `validate:"required,max=300"`
}

type GetVariantTemplate struct {
	ID uint `validate:"required"` // ID обязательно
}

type EditVariantTemplate struct {
	ID    uint   `validate:"required"`
	Title string `validate:"required,min=3,max=30"`
	Tasks string `validate:"required,max=300"`
	Tags  string `validate:"required,max=300"`
}

type DeleteVariantTemplate struct {
	ID uint `validate:"required"`
}
