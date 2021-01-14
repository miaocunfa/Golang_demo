package main

import "fmt"

// 即使底层使用相同类型 float64,
// Celsius Fahrenheit 也不是相同的类型
type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15    //绝对零度
	FreezingC     Celsius = 0          //冰点
	BoilingC      Celsius = 100        //沸点
)

func main() {
	fmt.Printf("%g\n", BoilingC - FreezingC)       // 摄氏 - 摄氏
	boilingF := CToF(BoilingC)                            // 华氏
	fmt.Printf("%g\n", boilingF - CToF(FreezingC)) // 华氏 - 华氏
	//fmt.Printf("%g\n", boilingF - FreezingC)	//类型不匹配   华氏 - 摄氏
}

// 摄氏度转华氏
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// 华氏转摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}