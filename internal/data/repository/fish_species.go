package repository

import (
	"TestTaskHaloLab/internal/core/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FishSpeciesRepository struct {
	db *gorm.DB
}

func NewFishSpeciesRepository(db *gorm.DB) *FishSpeciesRepository {
	return &FishSpeciesRepository{db: db}
}

func (fsr *FishSpeciesRepository) CreateFishSpecies(fishSpecies *models.FishSpecies) error {
	return fsr.db.Create(&fishSpecies).Error
}

func (fsr *FishSpeciesRepository) GetFishSpeciesBySensorDataIDs(sensorDataIDs []uuid.UUID) ([]models.FishSpecies, error) {
	var fishSpecies []models.FishSpecies
	err := fsr.db.
		Model(&models.FishSpecies{}).
		Where("sensor_data_id IN ?", sensorDataIDs).
		Select("name, SUM(count) as count").
		Group("name").
		Find(&fishSpecies).
		Error
	return fishSpecies, err
}

func (fsr *FishSpeciesRepository) GetTopNFishSpeciesBySensorDataIDs(sensorDataIDs []uuid.UUID, n int) ([]models.FishSpecies, error) {
	var fishSpecies []models.FishSpecies
	err := fsr.db.
		Model(&models.FishSpecies{}).
		Where("sensor_data_id IN ?", sensorDataIDs).
		Select("name, SUM(count) as count").
		Group("name").
		Order("count DESC").
		Limit(n).
		Find(&fishSpecies).
		Error
	return fishSpecies, err
}
