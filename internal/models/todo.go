package models

import "time"

// Todo представляет задачу пользователя.
type Todo struct {
	// ID — уникальный идентификатор задачи.
	ID int `json:"id"`

	// UserID — идентификатор пользователя, которому принадлежит задача.
	UserID int `json:"user_id"`

	// Title — название задачи.
	Title string `json:"title"`

	// Description — подробное описание задачи.
	Description string `json:"description"`

	// Done показывает, выполнена задача или нет.
	Done bool `json:"done"`

	// CreatedAt — дата и время создания задачи.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt — дата и время последнего изменения задачи.
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTodoRequest содержит данные для создания новой задачи.
type CreateTodoRequest struct {
	// UserID — идентификатор пользователя, которому принадлежит задача.
	UserID int `json:"user_id" validate:"required"`

	// Title — название задачи, максимум 200 символов.
	Title string `json:"title" validate:"required,max=200"`

	// Description — описание задачи, максимум 1000 символов.
	Description string `json:"description" validate:"max=1000"`

	// Done показывает, выполнена задача или нет.
	Done bool `json:"done"`
}

// UpdateTodoRequest содержит поля, которые можно изменить у существующей задачи.
type UpdateTodoRequest struct {
	// Title — новое название задачи, максимум 200 символов.
	Title *string `json:"title" validate:"omitempty,max=200"`

	// Description — новое описание задачи, максимум 1000 символов.
	Description *string `json:"description" validate:"omitempty,max=1000"`

	// Done — новое состояние выполнения задачи.
	Done *bool `json:"done"`
}
