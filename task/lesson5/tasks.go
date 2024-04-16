package lesson5

import "errors"

func Calcuclitor(num1, num2 int, sign string) (res int, err error) {
	switch sign {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		return num1 / num2, nil
	default:
		return 0, errors.New("don't know this operator")
	}
}
func Shot(dist float64) (res string) {
	switch {
	case dist >= 28. && dist <= 30.:
		return "ПОПА_Л"
	case dist > 30:
		return "ПЕРЕЛЕТ"
	case dist > 0 && dist < 28:
		return "НЕДОЛЕТ"
	case dist < 0:
		return "НЕ БЕЙ ПО СВОИМ"
	}

	return
}

func Mark(mark int) (strMark, isPassed string) {
	switch {
	case mark == 5:
		return "отлично", "экзамен сдан"
	case mark == 4:
		return "хорошо", "экзамен сдан"
	case mark == 3:
		return "удовлетворительно", ""
	case mark == 2:
		return "неудовлетворительно", ""

	}
	return
}
