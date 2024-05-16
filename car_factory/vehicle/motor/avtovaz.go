package motor

type AvtoVazMotor struct {
	Motor
}

func NewAutoVazMotor(power string) *AvtoVazMotor {
	motor := &AvtoVazMotor{
		Motor: Motor{
			manufacrurer: autoVAZ,
			isRun:        false,
		},
	}
	switch power {
	case "1.6":
		motor.amountCylinders = 4
		motor.power = power
	default:
		motor.amountCylinders = 4
		motor.power = defaultPower
	}
	return motor
}
