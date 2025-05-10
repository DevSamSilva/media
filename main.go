package main

import (
	"fmt"
	"log"
)

func verificaNota(nota float32) bool {
	return nota >= 0 && nota <= 10
}

func verificaAluno(alunos *int) {
	fmt.Println("Digite a quantidade de alunos (máximo 10):")
	fmt.Scan(&*alunos)

	if *alunos <= 0 {
		fmt.Println("Número de alunos inválido")
		log.Fatal()
	} else if *alunos > 10 {
		fmt.Println("Somente até 10 alunos são permitidos")
		log.Fatal()
	}
}

func main() {
	var numeroDeAlunos int
	var nota float32
	var soma float32

	verificaAluno(&numeroDeAlunos)

	for i := 1; i <= numeroDeAlunos; i++ {
		for {
			fmt.Printf("Digite a nota do aluno %d: ", i)
			fmt.Scan(&nota)

			if verificaNota(nota) {
				break
			} else {
				fmt.Println("Nota inválida, tente novamente!")
			}
		}
		soma += nota
	}

	media := soma / float32(numeroDeAlunos)
	fmt.Printf("A média da turma é: %.1f\n", media)
}
