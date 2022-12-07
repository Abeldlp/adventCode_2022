package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Elf struct {
	ElfIndex int
	Carry    int
}

func main() {
	file, err := os.Open("./calories.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	elfIndex := 0
	currentAmmount := 0
	strongestElf := Elf{}

	for fileScanner.Scan() {
		calories, err := strconv.Atoi(fileScanner.Text())

		currentAmmount += int(calories)

		if err != nil {
			currentElf := Elf{ElfIndex: elfIndex, Carry: currentAmmount}

			if currentElf.Carry > strongestElf.Carry {
				strongestElf = currentElf
			}

			elfIndex++
			currentAmmount = 0
		}
	}

	fmt.Println(strongestElf)
}
