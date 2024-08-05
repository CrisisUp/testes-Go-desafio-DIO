package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)  // arrange
	resultado := 6          // act
	if teste != resultado { // assert
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 8
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	teste := subtrai(10, 15)
	resultado := -5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldSubIncorrect(t *testing.T) {
	teste := subtrai(10, 15)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultCorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldMultIncorrect(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 42
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	teste := divide(20.0, 2.0)
	resultado := 10.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestShouldDivIncorrect(t *testing.T) {
	teste := divide(20.0, 2.0)
	resultado := 42.0
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}
