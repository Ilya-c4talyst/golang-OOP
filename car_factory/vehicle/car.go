package vehicle

import (
	"fmt"
	"golang/car_factory/vehicle/motor"
)

type Car struct {
	motor           motor.IMotor
	isLeftHandDrive bool
	brand           string
	isMove          bool
}

func (c *Car) ChangeMotor(motor motor.IMotor) {
	fmt.Println()
	fmt.Println("------Chage Motor-------")
	fmt.Printf("Old motor: %+v\n", c.motor)
	c.motor = motor
	fmt.Printf("New motor: %+v\n", c.motor)
	fmt.Println()
}

func (c *Car) GetBrand() string {
	return c.brand
}

func (c *Car) IsLeftHandDrive() bool {
	return c.isLeftHandDrive
}

func (c *Car) StartMotor() {
	c.motor.RunMotor()
}

func (c *Car) StopMotor() {
	c.motor.StopMotor()
}

func (c *Car) StartMove() {
	if c.motor.IsRun() {
		c.isMove = false
		fmt.Printf("Car %v started moving\n", c.brand)
	} else {
		fmt.Printf("Car %v cant start moving\n", c.brand)
	}
}

func (c *Car) StopMove() {
	if c.isMove {
		c.isMove = false
		fmt.Printf("Car %v stopped\n", c.brand)
	} else {
		fmt.Println("To stop first start moving")
	}

}

func (c *Car) GetMotorData() string {
	return fmt.Sprintf("%+v", c.motor)
}

func (c *Car) GetMotorPower() string {
	return c.motor.GetPower()
}

func (c *Car) IsMove() bool {
	return c.isMove
}

func Newcar(brand string, motor motor.IMotor, isLeftHandDrive bool) *Car {
	return &Car{
		motor:           motor,
		brand:           brand,
		isLeftHandDrive: isLeftHandDrive,
		isMove:          false,
	}
}

func NewDefaultCar() *Car {
	return &Car{
		motor:           motor.NewAutoVazMotor("1.6"),
		brand:           "Lada",
		isLeftHandDrive: false,
		isMove:          false,
	}
}
