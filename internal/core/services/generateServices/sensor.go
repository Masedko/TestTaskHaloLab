package generateServices

import (
	"TestTaskHaloLab/internal/core/models"
	"TestTaskHaloLab/internal/core/services/generateServices/utils"
	"TestTaskHaloLab/internal/data/static"
	"fmt"
	"log"
	"math/rand"
)

type SensorGenerateServiceSensorRepository interface {
	CreateSensors(sensors *[]models.Sensor) error
	GetSensors() ([]models.Sensor, error)
}

type SensorGenerateService struct {
	SensorRepository SensorGenerateServiceSensorRepository
}

func NewSensorGenerateService(sensorRepository SensorGenerateServiceSensorRepository) *SensorGenerateService {
	return &SensorGenerateService{SensorRepository: sensorRepository}
}

func (s *SensorGenerateService) GenerateSensors() {
	// Create MinSensorGroupCount-MaxSensorGroupCount sensor groups with 10-50 sensors in each group
	// Generate coordinates from 0 to 1000 (Imagine we have pivot point for each group in (0, 0, 0))
	// OutputRate is random number from 0 to 60 seconds

	// Check for existing sensors
	var sensors []models.Sensor
	sensors, err := s.SensorRepository.GetSensors()
	if len(sensors) != 0 {
		return
	}

	// Generate MinSensorGroupCount-MaxSensorGroupCount sensor groups
	groupCount := utils.GenerateIntInInterval(static.MinSensorGroupCount, static.MaxSensorGroupCount)

	for group := 0; group < groupCount; group++ {
		// Generate MinSensorCount-MaxSensorCount sensors in each group
		sensorsInGroupCount := utils.GenerateIntInInterval(static.MinSensorCount, static.MaxSensorCount)

		for sensorID := 0; sensorID < sensorsInGroupCount; sensorID++ {
			// Generate random coordinates from 0 to MaxCoordinate
			coordinateX := rand.Float64() * float64(utils.GenerateIntInInterval(static.XMin, static.XMax))
			coordinateY := rand.Float64() * float64(utils.GenerateIntInInterval(static.YMin, static.YMax))
			coordinateZ := rand.Float64() * float64(utils.GenerateIntInInterval(static.ZMin, static.ZMax))

			// Generate random output rate from 0 to 60 seconds
			outputRate := utils.GenerateIntInInterval(static.MinOutputRate, static.MaxOutputRate)

			// Create a new sensorID and append to slice of sensors
			sensors = append(sensors, models.Sensor{
				Name:        static.GreekLetters[group],
				Index:       sensorID,
				CoordinateX: coordinateX,
				CoordinateY: coordinateY,
				CoordinateZ: coordinateZ,
				OutputRate:  outputRate,
			})
		}
	}

	err = s.SensorRepository.CreateSensors(&sensors)
	if err != nil {
		log.Println(fmt.Sprintf("Cannot create sensors: %s", err.Error()))
	}
}
