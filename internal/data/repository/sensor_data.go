package repository

import (
	"TestTaskHaloLab/internal/core/dtos"
	"TestTaskHaloLab/internal/core/models"
	"TestTaskHaloLab/internal/data/repository/scopes"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SensorDataRepository struct {
	db *gorm.DB
}

func NewSensorDataRepository(db *gorm.DB) *SensorDataRepository {
	return &SensorDataRepository{db: db}
}

func (sdr *SensorDataRepository) CreateSensorData(sensorData *models.SensorData) error {
	return sdr.db.Create(&sensorData).Error
}

func (sdr *SensorDataRepository) GetAverageTemperatureByGroupName(groupName string) (dtos.Temperature, error) {
	var avgTemperature float32

	record := sdr.db.
		Model(models.SensorData{}).
		Scopes(scopes.GetWithGroupName(groupName)).
		Select("AVG(temperature)")
	err := record.Error
	if err != nil {
		return dtos.Temperature{}, err
	}

	err = record.
		Row().
		Scan(&avgTemperature)

	return dtos.Temperature{Temperature: avgTemperature}, err
}

func (sdr *SensorDataRepository) GetAverageTransparencyByGroupName(groupName string) (dtos.Transparency, error) {
	var avgTransparency float32
	record := sdr.db.
		Model(models.SensorData{}).
		Scopes(scopes.GetWithGroupName(groupName)).
		Select("AVG(transparency)")

	err := record.Error
	if err != nil {
		return dtos.Transparency{}, err
	}

	err = record.
		Row().
		Scan(&avgTransparency)
	return dtos.Transparency{Transparency: int(avgTransparency)}, err
}

func (sdr *SensorDataRepository) GetMinimumTemperatureOfGroups(groups []dtos.Group) (dtos.Temperature, error) {
	var minTemperature float32
	err := sdr.db.
		Model(&models.SensorData{}).
		Scopes(scopes.GetWithGroups(groups)).
		Select("MIN(temperature)").
		Scan(&minTemperature).Error
	return dtos.Temperature{Temperature: minTemperature}, err
}

func (sdr *SensorDataRepository) GetMaximumTemperatureOfGroups(groups []dtos.Group) (dtos.Temperature, error) {
	var maxTemperature float32
	err := sdr.db.
		Model(&models.SensorData{}).
		Scopes(scopes.GetWithGroups(groups)).
		Select("MAX(temperature)").
		Scan(&maxTemperature).Error
	return dtos.Temperature{Temperature: maxTemperature}, err
}

func (sdr *SensorDataRepository) GetAverageTemperatureByGroupWithTimeInterval(group *dtos.Group, timeInterval *dtos.TimeInterval) (dtos.Temperature, error) {
	var avgTemperature float32
	err := sdr.db.
		Model(&models.SensorData{}).
		Scopes(scopes.GetWithGroup(group)).
		Scopes(scopes.GetWithTimeInterval(timeInterval)).
		Select("AVG(temperature)").
		Scan(&avgTemperature).
		Error
	return dtos.Temperature{Temperature: avgTemperature}, err
}

func (sdr *SensorDataRepository) GetLastSensorDataIDsByGroupName(groupName string) ([]uuid.UUID, error) {
	var dataIDs []uuid.UUID
	err := sdr.db.
		Model(&models.SensorData{}).
		Select("id").
		Where("sensor_name = ?", groupName).
		Order("created_at DESC").
		Limit(1).
		Find(&dataIDs).Error
	return dataIDs, err
}

func (sdr *SensorDataRepository) GetSensorDataIDsInTimeIntervalByGroupName(groupName string, timeInterval *dtos.TimeInterval) ([]uuid.UUID, error) {
	dataIDs := make([]uuid.UUID, 0)
	err := sdr.db.
		Model(&models.SensorData{}).
		Where("sensor_name = ?", groupName).
		Scopes(scopes.GetWithTimeInterval(timeInterval)).
		Find(&dataIDs).Error
	return dataIDs, err
}

func (sdr *SensorDataRepository) GetSensorDataIDsInTimeInterval(timeInterval *dtos.TimeInterval) ([]uuid.UUID, error) {
	dataIDs := make([]uuid.UUID, 0)
	err := sdr.db.
		Model(&models.SensorData{}).
		Scopes(scopes.GetWithTimeInterval(timeInterval)).
		Find(&dataIDs).Error
	return dataIDs, err
}
