//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	arguments := os.Args
	fmt.Println("Пытаюсь создать директорию ", arguments[1])
	if len(arguments) < 2 {
		fmt.Println("Дорогуша, введите таки уже путь к директории проэкта.")
		os.Exit(1)
	}
	err := os.MkdirAll(arguments[1], 0755)
	if err != nil {
		log.Fatal("Ай какая неприятность! --> ", err)
	}
}
