// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv0

import "fmt"

type Celsius float64    //摄氏度
type Fahrenheit float64 //华氏度

const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //冰点温度
	BoilingC      Celsius = 100     //水沸腾温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 为类型 Celsius 定义 String() 方法
func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}
