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
func Average(numbers ...float64) float64 {
	total := float64(0)
	for _, x := range numbers {
		total += x
	}
	return total / float64(len(numbers))
}

/*
Вычисляет медиану последовательности чисел: если их нечетное количество - возвращается значение серединного элемента последовательности.
Если их четное количество - возвращается среднее арифметическое двух серединных элементов последовательности.
Функция Median не сортирует передаваемую последовательность. Если вам нужен результат медианы отсортированной последовательности - передавайте в функцию заранее отсортированный слайс
*/
func Median(numbers []float64) float64 {
	middle := len(numbers) / 2
	result := numbers[middle]
	if len(numbers)%2 == 0 {
		result = (result + numbers[middle-1]) / 2
	}
	return result
}

/*
Сумма значений произвольной последовательности чисел типа float64
*/
func Sum(numbers ...float64) (sum float64) {
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
Факториал числа
*/
func Fact(number int) (result int) {
	result = 1
	for i := 1; i <= number; i++ {
		result *= i
	}
	return result
}

/*
Произведение значений произвольной последовательности чисел типа float64
*/
func Multiplication(numbers ...float64) float64 {
	mul := float64(1)
	for _, number := range numbers {
		mul *= number
	}
	return mul
}

/*
Площадь окружности, по передаваемому радиусу типа float64
*/
func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

/*
Вычисляет длину результирующего вектора по координатам двух векторов из трехмерного пространства.
Первый вектор - вектор, у которого начало в точке A, конец в точке B.
Второй вектор - вектор, у которого начало в точке C, конец в точке D.
Передавать в функцию координаты точек A, B, C, D в следующем порядке:
(xA, xB, yA, yB, zA, zB, xC, xD, yC, yD, zC, zD), где x, y, z - значения координат осей x,y,z соотвественно. A, B, C, D - точки, к которым принадлежат
координаты x, y, z.
Например, xC - значение координаты x точки C, yD - значение координаты y точки D.
*/
func ResultingVector(xA, xB, yA, yB, zA, zB, xC, xD, yC, yD, zC, zD float64) float64 {
	return math.Sqrt((xB-xA+xD-xC)*(xB-xA+xD-xC) + (yB-yA+yD-yC)*(yB-yA+yD-yC) + (zB-zA+zD-zC)*(zB-zA+zD-zC))
}

/*
Вычисляет скалярное произведение двух векторов из трехмерного пространства.
Первый вектор - вектор, у которого начало в точке A, конец в точке B.
Второй вектор - вектор, у которого начало в точке C, конец в точке D.
Передавать в функцию координаты точек A, B, C, D в следующем порядке:
(xA, xB, yA, yB, zA, zB, xC, xD, yC, yD, zC, zD), где x, y, z - значения координат осей x,y,z соотвественно. A, B, C, D - точки, к которым принадлежат
координаты x, y, z.
Например, xC - значение координаты x точки C, yD - значение координаты y точки D.
*/
func ScalarProduct(xA, xB, yA, yB, zA, zB, xC, xD, yC, yD, zC, zD float64) float64 {
	return (xB-xA)*(xD-xC) + (yB-yA)*(yD-yC) + (zB-zA)*(zD-zC)
}

/*
Транспонирует матрицу размера 3 на 3
последовательно принимает значения (a11, a12, a13, a21, a22, a23, a31, a32, a33) исходной матрицы (по строкам), возвращает слайс
значений транспонированной матрицы (первые 3 значения слайса - 1ая строка транспонированной матрицы, следющие 3 значения слайса - 2ая строка транспонированной матрицы и тд)
*/
func TransposedMatrix3x3(a11, a12, a13, a21, a22, a23, a31, a32, a33 float64) []float64 {
	return []float64{a11, a21, a31, a12, a22, a32, a13, a23, a33}
}

/*
Площадь прямоугольника, по его сторонам, передаваемым в функцию (типа float64)
*/
func RectangleArea(a, b float64) float64 {
	return a * b
}

/*
Вычисляет минимальное значение и его позиции в последовательности (отсчет с нуля) по передаваемой в функцию последовательности чисел типа float 64.
Возвращает 2 аргумента - минимальное значение и слайс позиций минимальных значений в последовательности (их может быть несколько)
*/
func Min(numbers ...float64) (float64, []int) {
	min := numbers[0]
	var minPositions []int
	for i, number := range numbers {
		if number == min {
			minPositions = append(minPositions, i)
		} else if number < min {
			min = number
			minPositions = make([]int, 0)
			minPositions = append(minPositions, i)
		}
	}
	return min, minPositions
}

/*
Вычисляет максимальное значение и его позиции в последовательности (отсчет с нуля) по передаваемой в функцию последовательности чисел типа float 64.
Возвращает 2 аргумента - максимальное значение и слайс позиций максимальных значений в последовательности (их может быть несколько)
*/
func Max(numbers ...float64) (float64, []int) {
	max := numbers[0]
	var maxPositions []int
	for i, number := range numbers {
		if number == max {
			maxPositions = append(maxPositions, i)
		} else if number > max {
			max = number
			maxPositions = make([]int, 1)
			maxPositions = append(maxPositions, i)
		}
	}
	return max, maxPositions
}

/*
1ым аргументом принимает число, наличие которого нужно установить в последовательности. Далее - последовательность чисел, в которой требуется определить наличие искомого числа
Возвращает true, если число найдено в последовательности, иначе - false
*/
func SequenceHaveNumber(number float64, numbers ...float64) bool {
	for _, val := range numbers {
		if val == number {
			return true
		}
	}
	return false
}

/*
Принимает строку типа string - число в двоичном представлении little endian. Возвращает соответствующее ему значение в десятичной системе типа int
*/
func BinaryToDecimal(binary string) (result int) {
	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}


