// tempconv 包负责摄氏温度与华氏温度的转换
package tempconv

import "fmt"

type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度
type Kelvins float64    //开尔文

const (
	AbsoluteZeroC Celsius = -273.15    //绝对零度
	FreezingC     Celsius = 0          //冰点
	BoilingC      Celsius = 100        //沸点
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvins) String() string {
	return fmt.Sprintf("%gK", k)
}