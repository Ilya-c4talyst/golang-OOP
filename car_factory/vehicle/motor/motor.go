package motor

import "fmt"

const hyundai = "Hyundai"
const autoVAZ = "Автоваз"
const defaultPower = "1.4"

type Motor struct {
	power           string
	amountCylinders uint8
	isRun           bool
	manufacrurer    string
}

func (m *Motor) GetPower() string {
	return m.power
}

func (m *Motor) RunMotor() {
	fmt.Println("Mototr is run")
	m.isRun = true
}

func (m *Motor) IsRun() bool {
	return m.isRun
}

func (m *Motor) StopMotor() {
	fmt.Println("Motor is stop")
	m.isRun = false
}

func (m *Motor) AmountCylinders() uint8 {
	return m.amountCylinders
}

func (m *Motor) GetManufacturerName() string {
	return m.manufacrurer
}
