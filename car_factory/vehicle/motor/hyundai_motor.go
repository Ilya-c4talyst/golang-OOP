package motor

type HyundaiMotor struct {
	Motor
	turboMod bool
}

func (m *HyundaiMotor) IsTurboModOn() bool {
	return m.turboMod
}

func (m *HyundaiMotor) TurboModOn() {
	m.turboMod = true
}

func (m *HyundaiMotor) TurboModOff() {
	m.turboMod = false
}

func NewHyundaiMotor(power string) *HyundaiMotor {
	motor := &HyundaiMotor{
		Motor: Motor{
			manufacrurer: hyundai,
			isRun:        false,
		},
		turboMod: false,
	}
	switch power {
	case "1.8":
		motor.amountCylinders = 6
		motor.power = power
	case "1.6":
		motor.amountCylinders = 4
		motor.power = power
	default:
		motor.amountCylinders = 4
		motor.power = defaultPower
	}
	return motor
}
