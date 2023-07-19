package services

import "TestTaskHaloLab/internal/core/dtos"

type RegionServiceSensorRepository interface {
	GetGroupsInCoordinates(coordinates *dtos.Coordinates) ([]dtos.Group, error)
}

type RegionServiceSensorDataRepository interface {
	GetMinimumTemperatureOfGroups(groups []dtos.Group) (dtos.Temperature, error)
	GetMaximumTemperatureOfGroups(groups []dtos.Group) (dtos.Temperature, error)
}

type RegionService struct {
	SensorRepository     RegionServiceSensorRepository
	SensorDataRepository RegionServiceSensorDataRepository
}

func NewRegionService(sensorRepository RegionServiceSensorRepository, sensorDataRepository RegionServiceSensorDataRepository) *RegionService {
	return &RegionService{SensorRepository: sensorRepository, SensorDataRepository: sensorDataRepository}
}

func (s *RegionService) TemperatureMinimumInsideRegion(coordinates *dtos.Coordinates) (*dtos.Temperature, error) {
	groups, err := s.SensorRepository.GetGroupsInCoordinates(coordinates)
	if err != nil {
		return &dtos.Temperature{}, err
	}
	temperature, err := s.SensorDataRepository.GetMinimumTemperatureOfGroups(groups)
	return &temperature, nil
}

func (s *RegionService) TemperatureMaximumInsideRegion(coordinates *dtos.Coordinates) (*dtos.Temperature, error) {
	groups, err := s.SensorRepository.GetGroupsInCoordinates(coordinates)
	if err != nil {
		return &dtos.Temperature{}, err
	}
	temperature, err := s.SensorDataRepository.GetMaximumTemperatureOfGroups(groups)
	return &temperature, nil
}
