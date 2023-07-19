package dtos

type Fish struct {
	Name  string `validate:"required"`
	Count int    `validate:"required"`
}
