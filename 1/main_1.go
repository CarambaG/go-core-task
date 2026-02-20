package main

import (
	"crypto/sha256"
	hex2 "encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	dec := 10
	oct := 0o17
	hex := 0x2A
	f := 0.062
	str := "Hello"
	b := false
	c := complex64(complex(1.0, 2.0))

	array := []any{dec, oct, hex, f, str, b, c}

	// Вывод переменных
	for _, v := range array {
		fmt.Printf("Type %v - value %v\n", reflect.TypeOf(v), v)
	}
	fmt.Println()

	strSlice := make([]string, 0, len(array))

	for _, v := range array {
		switch val := v.(type) {
		case int:
			strSlice = append(strSlice, strconv.Itoa(v.(int)))
		case float64:
			strSlice = append(strSlice, strconv.FormatFloat(val, 'f', 2, 64))
		case string:
			strSlice = append(strSlice, v.(string))
		case bool:
			strSlice = append(strSlice, strconv.FormatBool(v.(bool)))
		case complex64:
			strSlice = append(strSlice, strconv.FormatComplex(complex128(val), 'f', 1, 64))
		}
	}

	finalStr := strings.Join(strSlice, "")
	fmt.Printf("Final str: %s\n\n", finalStr)

	runes := []rune(finalStr)
	fmt.Printf("Slice rune: %v\n\n", runes)

	runesSalted := string(runes[:len(runes)/2]) + "go-2024" + string(runes[len(runes)/2:])
	fmt.Printf("Add salt: %s\n\n", runesSalted)

	hash := sha256.New()
	hash.Write([]byte(finalStr))

	fmt.Printf("Полученный хэш: %v\n", hex2.EncodeToString(hash.Sum(nil)))
}
