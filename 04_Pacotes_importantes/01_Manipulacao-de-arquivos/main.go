package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/// Criar e escrever em arquivo
	fmt.Println("-- Criar e escrever em arquivo --")
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Hello, World!")
	size, err := f.Write([]byte("Hello, World! Em bytes"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! tamanho: %d bytes\n", size)
	err = f.Close()
	if err != nil {
		panic(err)
	}

	/// Leitura de arquivo
	fmt.Println("-- Leitura de arquivo --")
	file, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Conteudo dentro de arquivo.txt: %v\n", string(file))

	/// Leitura parte por parte (para menor uso de memoria)
	fmt.Println("-- Leitura parte por parte  --")
	fileOpen, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(fileOpen)
	buffer := make([]byte, 11)
	for {
		size, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:size]))
	}

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}

}
