package main

import "fmt"

type Vehicle interface {
	EngineStart()
	EngineStop()
	Drive(speed int)
}

type Car struct {
	Brand string
}

type Truck struct {
	Brand string
}

type Bike struct {
	Brand string // "HERO"
}

func (c Car) EngineStart() {
	fmt.Println("Car engine started")
}

func (c Car) EngineStop() {
	fmt.Println("Car engine stopped")
}

func (c Car) Drive(speed int) {
	fmt.Printf("Car is driving at speed %d\n", speed)
}

func (t Truck) EngineStart() {
	fmt.Println("Truck engine started")
}

func (t Truck) EngineStop() {
	fmt.Println("Truck engine stopped")
}

func (t Truck) Drive(speed int) {
	fmt.Printf("Truck is driving at speed %d\n", speed)
}
