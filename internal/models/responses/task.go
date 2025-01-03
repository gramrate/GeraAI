package responses

// TaskDTO описывает основную информацию о задаче
type TaskDTO struct {
	ID        uint   `json:"id"`
	AuthorID  uint   `json:"author_id"`
	Title     string `json:"title"`
	Condition string `json:"condition"`
	Answer    string `json:"answer"`
}

// CreateTaskResponseDTO описывает ответ на создание задачи
type CreateTaskResponseDTO struct {
	Status string  `json:"status"`
	Task   TaskDTO `json:"task"`
}

// GetTaskResponseDTO описывает ответ на запрос задачи по ID
type GetTaskResponseDTO struct {
	Status string  `json:"status"`
	Task   TaskDTO `json:"task"`
}

// EditTaskResponseDTO описывает ответ на редактирование задачи
type EditTaskResponseDTO struct {
	Status string  `json:"status"`
	Task   TaskDTO `json:"task"`
}

// DeleteTaskResponseDTO описывает ответ на удаление задачи
type DeleteTaskResponseDTO struct {
	Status string `json:"status"`
}

// GetAllTasksResponseDTO описывает ответ на запрос всех задач
type GetAllTasksResponseDTO struct {
	Tasks []TaskDTO `json:"tasks"`
}
