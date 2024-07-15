package package_learn

import "testing"

func TestFactorial(t *testing.T) {
	result := Math_Factorial(3)
	expected := 6

	if result != expected {
		t.Errorf("Factorial(3) = %d; expected %d", result, expected)
	}
}

/*
чтобы запустить тест конкретно этого пакет:
go test ./package_learn

чтобы запустить всеобщий тест:
go test ./...
*/
