package main

/**
	有关温度的转换函数
 */

func CToF(celsius Celsius) Fahrenheit {
	return Fahrenheit(celsius * 9 / 5 + 32)
}

func FToC(fahrenheit Fahrenheit) Celsius {
	return Celsius((fahrenheit - 32) * 5 / 9)
}