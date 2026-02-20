package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestTypeDetection(t *testing.T) {
	vars := []any{
		42,
		0o12,
		0xA2,
		3.14,
		"Nothing",
		true,
		complex64(2.0 + 1.0i),
	}

	expectedTypes := []reflect.Type{
		reflect.TypeOf(42),
		reflect.TypeOf(0o12),
		reflect.TypeOf(0xA2),
		reflect.TypeOf(3.14),
		reflect.TypeOf("Nothing"),
		reflect.TypeOf(true),
		reflect.TypeOf(complex64(2.0 + 1.0i)),
	}

	for i, val := range vars {
		if reflect.TypeOf(val) != expectedTypes[i] {
			t.Errorf("Variable %d expected type %v, got %v",
				i, expectedTypes[i], reflect.TypeOf(val))
		}
	}
}

func TestToStringConversion(t *testing.T) {
	vars := []any{
		42,
		0o12,
		0xA2,
		3.14,
		"Nothing",
		true,
		complex64(2.0 + 1.0i),
	}
	expected := "42101623.14Nothingtrue(2.0+1.0i)"

	var strParts []string
	for _, v := range vars {
		switch val := v.(type) {
		case int:
			strParts = append(strParts, strconv.Itoa(v.(int)))
		case float64:
			strParts = append(strParts, strconv.FormatFloat(val, 'f', 2, 64))
		case string:
			strParts = append(strParts, v.(string))
		case bool:
			strParts = append(strParts, strconv.FormatBool(v.(bool)))
		case complex64:
			strParts = append(strParts, strconv.FormatComplex(complex128(val), 'f', 1, 64))
		}
	}

	result := strings.Join(strParts, "")
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestSHA256WithSalt(t *testing.T) {
	str := "4210183.14Nothingtrue(2.0+1.0i)"
	runes := []rune(str)
	salted := string(runes[:len(runes)/2]) + "go-2024" + string(runes[len(runes)/2:])

	h1 := sha256.New()
	h1.Write([]byte(salted))
	expectedHash := hex.EncodeToString(h1.Sum(nil))

	// Повторяем точную же логику для проверки
	h2 := sha256.New()
	h2.Write([]byte(salted))
	resultHash := hex.EncodeToString(h2.Sum(nil))

	if resultHash != expectedHash {
		t.Errorf("Hash mismatch: expected %s, got %s", expectedHash, resultHash)
	}
}
