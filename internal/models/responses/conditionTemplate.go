package responses

// DTOs for condition templates

type ConditionTemplateDTO struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Condition string `json:"condition"`
}

type CreateConditionTemplateDTO struct {
	Status       string               `json:"status"`
	TaskTemplate ConditionTemplateDTO `json:"task_template"`
}

type GetConditionTemplateDTO struct {
	TaskTemplate ConditionTemplateDTO `json:"task_template"`
}

type EditConditionTemplateDTO struct {
	Status       string               `json:"status"`
	TaskTemplate ConditionTemplateDTO `json:"task_template"`
}

type DeleteConditionTemplateDTO struct {
	Status string `json:"status"`
}

type GetAllConditionTemplatesDTO struct {
	TaskTemplates []ConditionTemplateDTO `json:"task_templates"`
}
