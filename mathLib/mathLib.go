package mathLib

import (
	"fmt"
	"math"
)

/*
Сравнивает два значения типа float64 с заданной точностью или с самой большой аппаратной точностью, которую можно достичь,
если в аргументе limit было передано отрицательное число (например, -1).
*/
func EqualFloat64(x, y, limit float64) bool {
	if limit <= 0.0 {
		limit = math.SmallestNonzeroFloat64 //limit -> 0+
	}
	return math.Abs(x-y) <= (limit * math.Min(math.Abs(x), math.Abs(y)))
}

/*
Сравнение чисел в виде строк (более медленный способ).
Точность определяется количеством знаков после десятичной точки.
Если числа будут существенно отличаться по величине, соттветственно, будут отличаться и длины строк a и b (например, 12.32 и 592.85),
поэтому сравнение таких чисел будет выполняться быстрее.
*/
func EqualFloat64Prec(x, y float64, decimals int) bool {
	a := fmt.Sprintf("%.*f", decimals, x)
	b := fmt.Sprintf("%.*f", decimals, y)
	return len(a) == len(b) && a == b
}

/*
Округление числа в большую сторону, если дробная часть больше или равна 0.5, в противном случае - в меньшую сторону.
*/
func RoundingFloat64ToInt32(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}

/*
Среднее арифметическое произвольного количества чисел типа float64.
*/
func AverageFloat64(xs... float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}