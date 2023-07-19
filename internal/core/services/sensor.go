package services

import "TestTaskHaloLab/internal/core/dtos"

type SensorServiceSensorDataRepository interface {
	GetAverageTemperatureByGroupWithTimeInterval(group *dtos.Group, timeInterval *dtos.TimeInterval) (dtos.Temperature, error)
}

type SensorService struct {
	SensorDataRepository SensorServiceSensorDataRepository
}

func NewSensorService(sensorDataRepository SensorServiceSensorDataRepository) *SensorService {
	return &SensorService{SensorDataRepository: sensorDataRepository}
}

func (s *SensorService) AverageTemperatureInInterval(group *dtos.Group, timeInterval *dtos.TimeInterval) (*dtos.Temperature, error) {
	temperature, err := s.SensorDataRepository.GetAverageTemperatureByGroupWithTimeInterval(group, timeInterval)
	if err != nil {
		return &dtos.Temperature{}, err
	}
	return &temperature, nil
}
