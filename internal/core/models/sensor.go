package models

import (
	"github.com/google/uuid"
	"time"
)

type Sensor struct {
	// Codename consists of the GroupName and Index
	// X, Y, Z are coordinates of the Sensor
	// OutputRate is data output rate (in seconds)
	GroupName   string       `gorm:"not null;primaryKey"`
	Index       int          `gorm:"not null;primaryKey"`
	CoordinateX float64      `gorm:"not null"`
	CoordinateY float64      `gorm:"not null"`
	CoordinateZ float64      `gorm:"not null"`
	OutputRate  int          `gorm:"not null"`
	SensorData  []SensorData `gorm:"foreignKey:ID"`
}

type SensorData struct {
	// Temperature in Celsius
	// Transparency in %
	// CreatedAt stores data record time
	ID           *uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Temperature  float32       `gorm:"not null"`
	Transparency int           `gorm:"not null"`
	FishSpecies  []FishSpecies `gorm:"foreignKey:Name"`
	CreatedAt    time.Time     `gorm:"not null"`
}

type FishSpecies struct {
	// Names from https://oceana.org/ocean-fishes/
	ID    *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name  string     `gorm:"not null"`
	Count int        `gorm:"not null"`
}
