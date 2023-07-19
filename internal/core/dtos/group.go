package dtos

type Group struct {
	Name  string `validate:"required"` // Greek letter (alpha, beta and etc.)
	Index int    `validate:"required"`
}

type Transparency struct {
	Transparency int `validate:"required"`
}

type Temperature struct {
	Temperature float32 `validate:"required"`
}
