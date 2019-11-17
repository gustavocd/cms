package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

// Page represents a single page in the CMS
type Page struct {
	ID        int64     `json:"id" db:"id"`
	Label     string    `json:"label" db:"label"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	Slug      string    `json:"slug" db:"slug"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

// Validate validates Page data
func (p Page) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Label, validation.Required.Error("El campo etiqueta es obligatorio")),
		validation.Field(&p.Title, validation.Required.Error("El campo título es obligatorio")),
		validation.Field(&p.Body, validation.Required.Error("El campo contenido es obligatorio")),
		validation.Field(&p.Slug, validation.Required.Error("El campo slug es obligatorio")),
	)
}
