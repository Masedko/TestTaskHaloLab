package repository

import (
	"TestTaskHaloLab/internal/core/dtos"
	"TestTaskHaloLab/internal/core/models"
	"TestTaskHaloLab/internal/data/repository/scopes"
	"gorm.io/gorm"
)

type SensorRepository struct {
	db *gorm.DB
}

func NewSensorRepository(db *gorm.DB) *SensorRepository {
	return &SensorRepository{db: db}
}

func (sr *SensorRepository) CreateSensors(sensors *[]models.Sensor) error {
	return sr.db.Create(&sensors).Error
}

func (sr *SensorRepository) GetSensors() ([]models.Sensor, error) {
	var sensors []models.Sensor
	err := sr.db.Find(&sensors).Error
	return sensors, err
}

func (sr *SensorRepository) GetSensorByGroupNameIndex(group dtos.Group) (models.Sensor, error) {
	var sensor models.Sensor
	err := sr.db.Where(&models.Sensor{Name: group.Name, Index: group.Index}).First(&sensor).Error
	return sensor, err
}

func (sr *SensorRepository) GetGroupsInCoordinates(coordinates *dtos.Coordinates) ([]dtos.Group, error) {
	var groups []dtos.Group
	// https://gorm.io/docs/advanced_query.html#Smart-Select-Fields
	err := sr.db.Model(&models.Sensor{}).Scopes(scopes.InCoordinates(*coordinates)).Find(&groups).Error
	return groups, err
}
