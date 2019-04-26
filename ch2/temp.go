package main

import "fmt"

/**
	对温度对象的定义
 */

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC	Celsius	= -273.15
	FreezingC		Celsius = 0
	BoilingC 		Celsius = 100
)

func (celsius Celsius) String() string {
	return fmt.Sprintf("%g", celsius)
}

func (fahrenheit Fahrenheit) String() string {
	return fmt.Sprintf("%g", fahrenheit)
}
