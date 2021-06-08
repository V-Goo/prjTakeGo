//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Дорогуша, введите таки уже путь к директории проэкта.")
		os.Exit(1)
	}

	dirTarget := arguments[1]
	fileTarget := arguments[2]
	fmt.Println("Пытаюсь создать директорию ", dirTarget)
	err := os.MkdirAll(dirTarget, 0755)
	if err != nil {
		log.Fatal("Ай какая неприятность! --> ", err)
	}

	chdirErr := os.Chdir(dirTarget)
	if chdirErr != nil {
		log.Fatal("Что-то пошло не так! --> ", err)
	}
	currentPath, _ := os.Getwd()
	fmt.Println("Рабочая директория изменена на: ", currentPath)

	// Создать файл если он не существует, дописать если существует
	f, err := os.OpenFile(fileTarget, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("//usr/bin/env go run \"$0\" \"$@\"; exit\n")); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("code", fileTarget)
	log.Printf("Running %v command and waiting for it to finish...", cmd)
	// err := cmd.Start()
	errno := cmd.Run()
	log.Printf("Command finished with error? %v", errno)
}
