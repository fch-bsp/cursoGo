package matematica

const erroPadrao = "Valor esperado %v, mas o resuldado foi %v."

func TestMedia(t *testing.T) {
	ValorEsperado := 7.28
	valor := Media(7.2, 9.9, 5.9)

	if valor != VAlorEsperado {
		t.Errorf(erroPadrao, ValorEsperado, Valor)
	}

}
