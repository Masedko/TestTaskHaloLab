package services

import (
	"TestTaskHaloLab/internal/core/dtos"
	"TestTaskHaloLab/internal/core/models"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type GroupServiceCache interface {
	Get(key string) (string, error)
	Set(key string, value string, ttl time.Duration) error
}

type GroupServiceSensorDataRepository interface {
	GetAverageTemperatureByGroupName(groupName string) (dtos.Temperature, error)
	GetAverageTransparencyByGroupName(groupName string) (dtos.Transparency, error)
	GetLastSensorDataIDsByGroupName(groupName string) ([]uuid.UUID, error)
	GetSensorDataIDsInTimeInterval(timeInterval *dtos.TimeInterval) ([]uuid.UUID, error)
}

type GroupServiceFishSpeciesRepository interface {
	GetFishSpeciesBySensorDataIDs(sensorDataIDs []uuid.UUID) ([]models.FishSpecies, error)
	GetTopNFishSpeciesBySensorDataIDs(sensorDataIDs []uuid.UUID, n int) ([]models.FishSpecies, error)
}

type GroupService struct {
	Cache                 GroupServiceCache
	SensorDataRepository  GroupServiceSensorDataRepository
	FishSpeciesRepository GroupServiceFishSpeciesRepository
}

func NewGroupService(cache GroupServiceCache, sensorDataRepository GroupServiceSensorDataRepository, fishSpeciesRepository GroupServiceFishSpeciesRepository) *GroupService {
	return &GroupService{Cache: cache, SensorDataRepository: sensorDataRepository, FishSpeciesRepository: fishSpeciesRepository}
}

func (s *GroupService) TemperatureAverage(groupName string) (*dtos.Temperature, error) {
	// Define a cache key for this method and groupName
	cacheKey := fmt.Sprintf("TemperatureAverage:%s", groupName)

	// Check if the result is already cached
	if val, err := s.Cache.Get(cacheKey); err == nil {
		var temperature dtos.Temperature
		if err := json.Unmarshal([]byte(val), &temperature); err == nil {
			return &temperature, nil
		}
	}

	// If not cached, get the result from the repository
	temperature, err := s.SensorDataRepository.GetAverageTemperatureByGroupName(groupName)
	if err != nil {
		return &dtos.Temperature{}, err
	}

	// Cache the result with a 10s TTL
	if val, err := json.Marshal(temperature); err == nil {
		err := s.Cache.Set(cacheKey, string(val), 10*time.Second)
		if err != nil {
			return nil, err
		}
	}

	return &temperature, nil
}

func (s *GroupService) TransparencyAverage(groupName string) (*dtos.Transparency, error) {
	// Define a cache key for this method and groupName
	cacheKey := fmt.Sprintf("TransparencyAverage:%s", groupName)

	// Check if the result is already cached
	if val, err := s.Cache.Get(cacheKey); err == nil {
		var transparency dtos.Transparency
		if err := json.Unmarshal([]byte(val), &transparency); err == nil {
			return &transparency, nil
		}
	}

	// If not cached, get the result from the repository
	transparency, err := s.SensorDataRepository.GetAverageTransparencyByGroupName(groupName)
	if err != nil {
		return &dtos.Transparency{}, err
	}

	// Cache the result with a 10s TTL
	if val, err := json.Marshal(transparency); err == nil {
		err := s.Cache.Set(cacheKey, string(val), 10*time.Second)
		if err != nil {
			return nil, err
		}
	}

	return &transparency, nil
}

func (s *GroupService) Species(groupName string) (*[]dtos.Fish, error) {
	ids, err := s.SensorDataRepository.GetLastSensorDataIDsByGroupName(groupName)
	if err != nil {
		return &[]dtos.Fish{}, err
	}
	fishSpecies, err := s.FishSpeciesRepository.GetFishSpeciesBySensorDataIDs(ids)
	if err != nil {
		return &[]dtos.Fish{}, err
	}
	var fish []dtos.Fish
	for _, species := range fishSpecies {
		fish = append(fish, dtos.Fish{Name: species.Name, Count: species.Count})
	}
	return &fish, nil
}

func (s *GroupService) SpeciesTopN(groupName string, N int, timeInterval *dtos.TimeInterval) (*[]dtos.Fish, error) {
	var ids []uuid.UUID
	var err error
	if !timeInterval.StarTime.IsZero() && !timeInterval.EndTime.IsZero() {
		ids, err = s.SensorDataRepository.GetSensorDataIDsInTimeInterval(timeInterval)
		if err != nil {
			return &[]dtos.Fish{}, err
		}
	} else {
		ids, err = s.SensorDataRepository.GetLastSensorDataIDsByGroupName(groupName)
		if err != nil {
			return &[]dtos.Fish{}, err
		}
	}
	fishSpecies, err := s.FishSpeciesRepository.GetTopNFishSpeciesBySensorDataIDs(ids, N)
	if err != nil {
		return &[]dtos.Fish{}, err
	}
	var fish []dtos.Fish
	for _, species := range fishSpecies {
		fish = append(fish, dtos.Fish{Name: species.Name, Count: species.Count})
	}
	return &fish, nil
}
