package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}

}

func main() {
	/* fale("Maria", "Pq vc não falou comigo?", 3)
	fale("João", "Só posso falar depois de vc!", 1)  */

	/* go fale("Maria", "	Ei...", 500)
	go fale("JOão", "Opa...", 500)

	fmt.Println("Fim..")
 */


	go fale("Maria", "Endenti!!", 10)
	fale("João", "Parabńes!", 5)
}
