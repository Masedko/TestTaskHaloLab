package models

import (
	"github.com/google/uuid"
	"time"
)

type SensorData struct {
	// Temperature in Celsius
	// Transparency in %
	// CreatedAt stores data record time
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Temperature  float32   `gorm:"not null"`
	Transparency int       `gorm:"not null"`
	SensorName   string    `gorm:"not null"`
	SensorIndex  int       `gorm:"not null"`
	Sensor       *Sensor   `gorm:"foreignKey:SensorName,SensorIndex"`
	CreatedAt    time.Time
}
