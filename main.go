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

func armazenaNota(aluno int, notas float32, somas *float32) {
	for i := 1; i <= aluno; i++ {
		for {
			fmt.Printf("Digite a nota do aluno %d: ", i)
			fmt.Scan(&notas)

			if verificaNota(notas) {
				break
			} else {
				fmt.Println("Nota inválida, tente novamente!")
			}
		}
		*somas += notas
	}

}

func Media(med float32, somas float32, alunos int) {
	med = somas / float32(alunos)
	fmt.Printf("A média da turma é: %.1f\n", med)

}

func main() {
	var numeroDeAlunos int
	var nota float32
	var soma float32
	var media float32

	verificaAluno(&numeroDeAlunos)

	armazenaNota(numeroDeAlunos, nota, &soma)

	Media(media, soma, numeroDeAlunos)

}
