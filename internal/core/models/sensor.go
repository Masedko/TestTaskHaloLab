package models

type Sensor struct {
	// codeName consists of the Name(greek letters) and Index
	// X, Y, Z are coordinates of the Sensor
	// OutputRate is data output rate (in seconds)
	Name        string  `gorm:"not null;primaryKey"`
	Index       int     `gorm:"not null;primaryKey"`
	CoordinateX float64 `gorm:"not null"`
	CoordinateY float64 `gorm:"not null"`
	CoordinateZ float64 `gorm:"not null"`
	OutputRate  int     `gorm:"not null"`
}
