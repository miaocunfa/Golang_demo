package tempconv

// 摄氏度转华氏
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c * 9 / 5 + 32)
}

// 华氏转摄氏度
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 开尔文转摄氏度
func KToC(k Kelvins) Celsius {
	return Celsius(k - 273.15)
}