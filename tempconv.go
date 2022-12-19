// вычисление температур для цельсия и фаренгейта
package tempconv


type Celsius    float64
type Fahrenheit float64

const (
    AbsoluteZeroC   Celsius = -273.15
    FreezingC       Celsius = 0
    BoilingC        Celsius = 100
)

// T(x) приводит значение x к типу T

// Большая буква экспортируется
// маленькая буква не портуется
