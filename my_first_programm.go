package my_awesome_project


import (
"bufio"
"fmt"
"os"
"regexp"
"strconv"
"strings"
) // Hi!

func scanAll(result string) []string {
	var arr []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при считывании строки. Программа будет принудительно завершена")
		os.Exit(0)
	}
	result = scanner.Text()
	result = strings.ReplaceAll(result, " ", "")
	//result = strings.ReplaceAll(result, "	", "")
	rome := regexp.MustCompile(`^(I{1,3}|IV|V|VI|VII|VIII|IX|X)(\*|\/|\+|\-)(I{1,3}|IV|V|VI|VII|VIII|IX|X)$`)
	arab := regexp.MustCompile(`^(10|[1-9])(\*|\/|\+|\-)(10|[1-9])$`)
	if arab.MatchString(result) || rome.MatchString(result) {
		arr = strings.Split(result, "")
	} else {
		fmt.Printf("Ввод символов некорректный. Введите выражение формата 1 + 1 или I + I. ")
		os.Exit(0)
	}

	/*!(arab.MatchString(result) || rome.MatchString(result)) {
		fmt.Printf("FAILL:1")
		os.Exit(0)
	}*/
	//arr = strings.Split(result, "")
	return arr

}

func separation(fullExp []string) (string, string, string) {
	i := 0
	for index, value := range fullExp {
		if value == "+" || value == "-" || value == "/" || value == "*" {
			i = index
			break
		}
	}
	operation := fullExp[i]
	first := strings.Join(fullExp[:i], "")
	second := strings.Join(fullExp[i+1:], "")
	return operation, first, second
}
func romenum(rome string) int {
	switch rome {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return 0
	}
}
func resultRome(i int) string {
	rNum := [101]string{
		"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	if i < 1 {
		fmt.Println("Результат расчета в римских цифрах не может быть меньше I")
		os.Exit(0)
	}
	return rNum[i]
}
func operation(str string, x, y int) int {
	switch str {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	}
	return 0
}
func main() {
	var a2, a3 int
	k := false
	//var operandeRome1, operandeRome2 string
	var result string
	var resultNum int
	var preExp []string
	fmt.Printf("Введите выражение:")
	preExp = scanAll(result)

	var i1, i2, i3 = separation(preExp)

	if _, err := strconv.Atoi(i2); err == nil {
		a2, err = strconv.Atoi(i2)
	} else {
		k = true
		a2 = romenum(i2)

	}
	if _, err := strconv.Atoi(i3); err == nil {
		a3, err = strconv.Atoi(i3)
	} else {
		a3 = romenum(i3)
		k = true
	}

	if k == true {
		resultNum = operation(i1, a2, a3)
		fmt.Println("Результат вычислений:", resultRome(resultNum))
	} else {
		resultNum = operation(i1, a2, a3)
		fmt.Println("Результат вычислений:", resultNum)
	}

	//	fmt.Println(i1)
	//fmt.Println(a2)
	//fmt.Println(a3)

	//fmt.Println("asfaf")
}
