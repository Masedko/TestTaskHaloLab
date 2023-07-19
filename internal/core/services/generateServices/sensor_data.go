package generateServices

import (
	"TestTaskHaloLab/internal/core/models"
	"TestTaskHaloLab/internal/core/services/generateServices/utils"
	"TestTaskHaloLab/internal/data/static"
	"fmt"
	"github.com/google/uuid"
	"log"
	"math/rand"
	"sync"
	"time"
)

type SensorDataGenerateServiceSensorRepository interface {
	GetSensors() ([]models.Sensor, error)
}

type SensorDataGenerateServiceSensorDataRepository interface {
	CreateSensorData(sensorData *models.SensorData) error
}

type SensorDataGenerateServiceFishSpeciesRepository interface {
	CreateFishSpecies(fishSpecies *models.FishSpecies) error
}

type SensorDataGenerateService struct {
	SensorRepository      SensorDataGenerateServiceSensorRepository
	SensorDataRepository  SensorDataGenerateServiceSensorDataRepository
	FishSpeciesRepository SensorDataGenerateServiceFishSpeciesRepository
}

func NewSensorDataGenerateService(sensorRepository SensorDataGenerateServiceSensorRepository,
	sensorDataRepository SensorDataGenerateServiceSensorDataRepository,
	fishSpeciesRepository SensorDataGenerateServiceFishSpeciesRepository) *SensorDataGenerateService {
	return &SensorDataGenerateService{SensorRepository: sensorRepository,
		SensorDataRepository:  sensorDataRepository,
		FishSpeciesRepository: fishSpeciesRepository,
	}
}

func (s *SensorDataGenerateService) GenerateSensorData() {
	var wg sync.WaitGroup
	sensors, err := s.SensorRepository.GetSensors()
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot get sensors: %s", err.Error()))
	}
	wg.Add(len(sensors))
	for _, sensor := range sensors {
		go GenerateSensorData(sensor, s, &wg)
	}
}

func GenerateSensorData(sensor models.Sensor, sensorDataGenerateService *SensorDataGenerateService, wg *sync.WaitGroup) {
	// Generate data for the sensor based on the output rate
	// You can use a ticker to periodically generate the data
	defer wg.Done()
	ticker := time.NewTicker(time.Second * time.Duration(sensor.OutputRate))
	var sensorData models.SensorData
	var fishSpecies models.FishSpecies
	var fishCount int
	var fishSpeciesCount int
	var fishSpeciesNames []string
	var minFishSpeciesCount, maxFishSpeciesCount int
	var sumOfCoords float64
	var maxSumOfCoords float64
	for range ticker.C {
		sensorDataID, err := uuid.NewUUID()
		if err != nil {
			log.Fatal("Cannot create UUID")
		}
		sumOfCoords = sensor.CoordinateX + sensor.CoordinateY + sensor.CoordinateZ
		maxSumOfCoords = static.XMax + static.YMax + static.ZMax

		// From 1 - RandomizeCoefficient to 1 + RandomizeCoefficient
		randomizeTemperatureCoefficient := 1 + static.RandomizeCoefficient - 2*rand.Float32()*static.RandomizeCoefficient
		randomizeTransparencyCoefficient := 1 + static.RandomizeCoefficient - 2*rand.Float64()*static.RandomizeCoefficient

		sensorData = models.SensorData{
			ID:           sensorDataID,
			Temperature:  float32(static.TemperatureWithDepth*sensor.CoordinateZ) * randomizeTemperatureCoefficient,
			Transparency: int((maxSumOfCoords - (sumOfCoords)) / (maxSumOfCoords / 100) * randomizeTransparencyCoefficient),
			SensorName:   sensor.Name,
			SensorIndex:  sensor.Index,
		}
		err = sensorDataGenerateService.SensorDataRepository.CreateSensorData(&sensorData)
		if err != nil {
			log.Fatal("Cannot create sensorData")
		}
		switch {
		case sensor.CoordinateZ < static.ShallowLevel:
			fishCount = rand.Intn(len(static.ShallowFishNames))
			minFishSpeciesCount = static.MinShallowLevelFishSpeciesCount
			maxFishSpeciesCount = static.MaxShallowLevelFishSpeciesCount
			fishSpeciesNames = make([]string, len(static.ShallowFishNames))
			copy(fishSpeciesNames, static.ShallowFishNames)
			utils.ShuffleSliceOfString(fishSpeciesNames)
		case sensor.CoordinateZ < static.MiddleLevel:
			fishCount = rand.Intn(len(static.MiddleFishNames))
			minFishSpeciesCount = static.MinMiddleLevelFishSpeciesCount
			maxFishSpeciesCount = static.MaxMiddleLevelFishSpeciesCount
			fishSpeciesNames = make([]string, len(static.MiddleFishNames))
			copy(fishSpeciesNames, static.MiddleFishNames)
			utils.ShuffleSliceOfString(fishSpeciesNames)
		case sensor.CoordinateZ < static.DeepLevel:
			fishCount = rand.Intn(len(static.DeepFishNames))
			minFishSpeciesCount = static.MinDeepLevelFishSpeciesCount
			maxFishSpeciesCount = static.MaxDeepLevelFishSpeciesCount
			fishSpeciesNames = make([]string, len(static.DeepFishNames))
			copy(fishSpeciesNames, static.DeepFishNames)
			utils.ShuffleSliceOfString(fishSpeciesNames)
		}
		for fish := 0; fish < fishCount-1; fish++ {
			fishSpeciesCount = utils.GenerateIntInInterval(minFishSpeciesCount, maxFishSpeciesCount)
			fishSpecies = models.FishSpecies{Name: fishSpeciesNames[fish], Count: fishSpeciesCount, SensorDataID: &sensorDataID}
			err = sensorDataGenerateService.FishSpeciesRepository.CreateFishSpecies(&fishSpecies)
			if err != nil {
				log.Fatal("Cannot create Fish Species")
			}
		}

	}
}
