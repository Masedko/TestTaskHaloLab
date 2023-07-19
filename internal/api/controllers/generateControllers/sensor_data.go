package generateControllers

type SensorDataGenerateService interface {
	GenerateSensorData()
}

type SensorDataGenerateController struct {
	SensorDataGenerateService SensorDataGenerateService
}

func NewSensorDataGenerateController(sensorDataGenerateService SensorDataGenerateService) *SensorDataGenerateController {
	return &SensorDataGenerateController{SensorDataGenerateService: sensorDataGenerateService}
}

func (sc *SensorDataGenerateController) GenerateSensorData() {
	sc.SensorDataGenerateService.GenerateSensorData()
}
