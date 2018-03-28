// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv0

import "fmt"

type Celsius float64    //摄氏度
type Fahrenheit float64 //华氏度
type Kelvin float64     //开尔文
const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //冰点温度
	BoilingC      Celsius = 100     //水沸腾温度
)

// 为类型 Celsius 定义 String() 方法
func (c Celsius) String() string {
	return fmt.Sprintf("%gC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gF", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
