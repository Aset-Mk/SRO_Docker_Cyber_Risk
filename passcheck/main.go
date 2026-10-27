// passcheck — консольная проверка стойкости паролей.
// Считает энтропию, ищет пароль в словаре утёкших и оценивает время перебора.
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

var leaked = map[string]bool{
	"123456": true, "password": true, "qwerty": true, "111111": true,
	"12345678": true, "admin": true, "iloveyou": true, "qwerty123": true,
	"1q2w3e4r": true, "password1": true, "123123": true, "letmein": true,
}

func alphabet(p string) int {
	var lower, upper, digit, other bool
	for _, r := range p {
		switch {
		case unicode.IsLower(r):
			lower = true
		case unicode.IsUpper(r):
			upper = true
		case unicode.IsDigit(r):
			digit = true
		default:
			other = true
		}
	}
	n := 0
	if lower {
		n += 26
	}
	if upper {
		n += 26
	}
	if digit {
		n += 10
	}
	if other {
		n += 33
	}
	return n
}

func crackTime(bits float64) string {
	// 10^10 попыток/с — офлайн-перебор быстрого хэша на GPU
	sec := math.Pow(2, bits) / 2 / 1e10
	switch {
	case sec < 1:
		return "мгновенно"
	case sec < 3600:
		return fmt.Sprintf("%.0f сек", sec)
	case sec < 86400:
		return fmt.Sprintf("%.1f ч", sec/3600)
	case sec < 86400*365:
		return fmt.Sprintf("%.0f дней", sec/86400)
	case sec < 86400*365*1e6:
		return fmt.Sprintf("%.0f лет", sec/86400/365)
	default:
		return "больше миллиона лет"
	}
}

func check(p string) {
	length := len([]rune(p))
	bits := float64(length) * math.Log2(float64(max(alphabet(p), 1)))
	verdict := "СТОЙКИЙ"
	switch {
	case leaked[strings.ToLower(p)]:
		verdict, bits = "В СПИСКЕ УТЁКШИХ", 0
	case bits < 40:
		verdict = "СЛАБЫЙ"
	case bits < 60:
		verdict = "СРЕДНИЙ"
	}
	fmt.Printf("%-22s длина=%-3d энтропия=%5.1f бит  перебор: %-20s %s\n",
		p, length, bits, crackTime(bits), verdict)
}

func main() {
	if len(os.Args) > 1 {
		for _, p := range os.Args[1:] {
			check(p)
		}
		return
	}
	// без аргументов — читаем пароли из stdin, по одному на строку
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		if p := strings.TrimSpace(sc.Text()); p != "" {
			check(p)
		}
	}
}
