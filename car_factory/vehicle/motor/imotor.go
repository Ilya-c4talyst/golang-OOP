package motor

type IMotor interface {
	GetPower() string
	RunMotor()
	StopMotor()
	IsRun() bool
	AmountCylinders() uint8
	GetManufacturerName() string
}
