package generateControllers

type SensorGenerateService interface {
	GenerateSensors()
}

type SensorGenerateController struct {
	SensorGenerateService SensorGenerateService
}

func NewSensorGenerateController(sensorGenerateService SensorGenerateService) *SensorGenerateController {
	return &SensorGenerateController{SensorGenerateService: sensorGenerateService}
}

func (sc *SensorGenerateController) GenerateSensors() {
	sc.SensorGenerateService.GenerateSensors()
}
