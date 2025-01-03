package responses

type InterestsTemplateDTO struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	Interests []string `json:"interests"`
}

type CreateInterestsTemplateDTO struct {
	Status       string               `json:"status"`
	TaskTemplate InterestsTemplateDTO `json:"task_template"`
}

type GetInterestsTemplateDTO struct {
	TaskTemplate InterestsTemplateDTO `json:"task_template"`
}

type EditInterestsTemplateDTO struct {
	Status       string               `json:"status"`
	TaskTemplate InterestsTemplateDTO `json:"task_template"`
}

type DeleteInterestsTemplateDTO struct {
	Status string `json:"status"`
}

type GetAllInterestsTemplatesDTO struct {
	TaskTemplates []InterestsTemplateDTO `json:"task_templates"`
}
