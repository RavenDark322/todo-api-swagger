package models

import "time"

// User представляет пользователя системы.
type User struct {
	// ID — уникальный идентификатор пользователя.
	ID int `json:"id"`

	// Email — электронная почта пользователя.
	Email string `json:"email"`

	// Name — имя пользователя.
	Name string `json:"name"`

	// CreatedAt — дата и время создания пользователя.
	CreatedAt time.Time `json:"created_at"`
}

// CreateUserRequest содержит данные для создания нового пользователя.
type CreateUserRequest struct {
	// Email — электронная почта нового пользователя.
	Email string `json:"email" validate:"required,email"`

	// Name — имя нового пользователя, максимум 100 символов.
	Name string `json:"name" validate:"required,max=100"`
}

// UpdateUserRequest содержит поля, которые можно изменить у существующего пользователя.
type UpdateUserRequest struct {
	// Email — новое значение электронной почты пользователя.
	Email *string `json:"email" validate:"omitempty,email"`

	// Name — новое имя пользователя, максимум 100 символов.
	Name *string `json:"name" validate:"omitempty,max=100"`
}

// ErrorResponse описывает ответ API при возникновении ошибки.
type ErrorResponse struct {
	// Error — краткий тип или название ошибки.
	Error string `json:"error"`

	// Message — подробное описание ошибки.
	Message string `json:"message"`
}

// SuccessResponse описывает стандартный успешный ответ API.
type SuccessResponse struct {
	// Success показывает, успешно ли выполнен запрос.
	Success bool `json:"success"`

	// Data содержит данные ответа.
	Data interface{} `json:"data,omitempty"`
}
