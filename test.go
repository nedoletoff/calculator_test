package main

import (
    "fmt"
    "strconv"
    "errors"
    "strings"
)

type Nums struct {
    n1 int
    n2 int
}

func calculate(n1, n2, o string) string {
    _, err1 := strconv.ParseInt(n1, 10, 12)
    if err1 == nil {
        return strconv.Itoa(num_operation(string_to_arab(n1, n2), get_operation(o)))
    }
    nn1, _ := rome_to_arab_10(n1)
    nn2, _ := rome_to_arab_10(n2)
    nums := Nums{nn1, nn2}

    r := num_operation(nums, get_operation(o))
    if r < 1 {
        return "Римские числа не могут быть отрицательными"
    }
    return arab_to_rome(r)
}

func num_operation(nums Nums, operation int) int {
    var r, n1, n2 int
    n1, n2 = nums.n1, nums.n2
    switch operation {
    case 0:
        r = n1 + n2
    case 1:
        r = n1 - n2
    case 2:
        r = n1 * n2
    case 3:
        r = n1 / n2
    }
    return r
}

func string_to_arab(n1, n2 string) Nums {
    nn1, err1 := strconv.ParseInt(n1, 10, 12)
    if err1 != nil {
        fmt.Println(err1)
    }
    nn2, err2 := strconv.ParseInt(n2, 10, 12)
    if err2 != nil {
        fmt.Println(err2)
    }
    return Nums{int(nn1), int(nn2)}
}

func get_operation(o string) int {
    switch o {
    case "+":
        return 0
    case "-":
        return 1
    case "*":
        return 2
    case "/":
        return 3
    default:
        return 4
    }
}

func arab_to_rome(a int) string {
    var s []string
    if a >= 100 {
        a -= 100
        s = append(s, "C")
    }
    for a >= 90 {
        a -= 90
        s = append(s, "XC")
    }
    for a >= 50 {
        a -= 50
        s = append(s, "L")
    }
    for a >= 40 {
        a -= 40
        s = append(s, "XL")
    }
    for a >= 10 {
        a -= 10
        s = append(s, "X")
    }
    for a >= 9 {
        a -= 9
        s = append(s, "IX")
    }
    for a >= 5 {
        a -= 5
        s = append(s, "V")
    }
    for a >= 4 {
        a -= 4
        s = append(s, "IV")
    }
    for a >= 1 {
        a -= 1
        s = append(s, "I")
    }
    return strings.Join(s, "")
}

func rome_to_arab_10(r string) (int, error) {
    rome := [10]string {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

    for i := 0; i < 10; i++ {
        if r == rome[i] {
            return i + 1, nil
        }
    }
    return 0, errors.New("Неправильный ввод чисел")
}

func check_enter(n1, n2 string) bool {
    nn1, err1 := strconv.ParseInt(n1, 10, 12)
    nn2, err2 := strconv.ParseInt(n2, 10, 12)
    if err1 == nil && err2 == nil {
        if nn1 > 0 && nn1 < 11 && nn2 > 0 && nn2 < 11 {
            return false
        }
        return true
    }

    _, err1 = rome_to_arab_10(n1)
    _, err2 = rome_to_arab_10(n2)
    if err1 == nil && err2 == nil {
        return false
    }
    return true
}

func check_operator(o string) bool {
    if get_operation(o) == 4 {
        return true
    }
    return false
}


func main() {
    var n1, o, n2, c string

    fmt.Println("Enter string")
    fmt.Scanln(&n1, &o, &n2, &c)

    if c !=  "" {
        fmt.Println("Неверный формат ввода")
        return
    }

    if check_enter(n1, n2) {
        fmt.Println("Неправильный ввод чисел")
        return
    }
    if check_operator(o) {
        fmt.Println("Неправильный оператор")
        return
    }

    fmt.Println(calculate(n1, n2, o))
}
