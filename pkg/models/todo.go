package models

// Todo represents a to-do item as returned by the JSONPlaceholder /todos endpoint.
type Todo struct {
	// UserID is the identifier of the user this to-do belongs to.
	UserID int `json:"userId"`

	// ID is the unique identifier of the to-do item.
	ID int `json:"id"`

	// Title is the description of the task.
	Title string `json:"title"`

	// Completed indicates whether the to-do item has been finished.
	Completed bool `json:"completed"`
}
