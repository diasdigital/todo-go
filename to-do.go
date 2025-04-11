package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var tasks []string


func tambahTask(scanner *bufio.Scanner) {
	fmt.Print("Masukkan task baru: ")
	if scanner.Scan() {
		task := strings.TrimSpace(scanner.Text())
		if task != "" {
			tasks = append(tasks, task)
			fmt.Println("Task berhasil ditambahkan!")
		} else {
			fmt.Println("Task tidak boleh kosong.")
		}
	}
}

func tampilkanTasks() {
	fmt.Println("\n📋 Daftar To-Do:")
	if len(tasks) == 0 {
		fmt.Println("- Belum ada task.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n=== To-Do App ===")
		fmt.Println("1. Tambah Task")
		fmt.Println("2. Lihat Task")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		if !scanner.Scan() {
			break
		}

		pilihan := strings.TrimSpace(scanner.Text())

		switch pilihan {
		case "1":
			tambahTask(scanner)
		case "2":
			tampilkanTasks()
		case "3":
			fmt.Println("Terima kasih, sampai jumpa 👋")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}