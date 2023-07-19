package dtos

type Coordinates struct {
	XMin float64 `validate:"required" query:"xMin"`
	XMax float64 `validate:"required" query:"xMax"`
	YMin float64 `validate:"required" query:"yMin"`
	YMax float64 `validate:"required" query:"yMax"`
	ZMin float64 `validate:"required" query:"zMin"`
	ZMax float64 `validate:"required" query:"zMax"`
}

type CoordinatesString struct {
	XMin string `validate:"required" query:"xMin"`
	XMax string `validate:"required" query:"xMax"`
	YMin string `validate:"required" query:"yMin"`
	YMax string `validate:"required" query:"yMax"`
	ZMin string `validate:"required" query:"zMin"`
	ZMax string `validate:"required" query:"zMax"`
}
