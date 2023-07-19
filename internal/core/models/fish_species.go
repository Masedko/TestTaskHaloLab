package models

import "github.com/google/uuid"

type FishSpecies struct {
	// Names from https://oceana.org/ocean-fishes/
	ID           uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name         string      `gorm:"not null"`
	Count        int         `gorm:"not null"`
	SensorDataID *uuid.UUID  `gorm:"not null"`
	SensorData   *SensorData `gorm:"foreignKey:SensorDataID"`
}
