package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	tasks := loadTasksFromFile("tasks.txt")

	for {
		fmt.Println("\n=== To-Do App ===")
		fmt.Println("Pilih aksi:")
		fmt.Println("1. Tambah Task")
		fmt.Println("2. Lihat task")
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		if !scanner.Scan() {
			break
		}

		pilihan := strings.TrimSpace(scanner.Text())

		switch pilihan {
		case "1":
			tasks = tambahTask(scanner, tasks)
			saveTasksToFile("tasks.txt", tasks)
		case "2":
			tampilkanTasks(tasks)
		case "3":
			fmt.Println("Terima kasih, sampai jumpa")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahTask(scanner *bufio.Scanner, tasks []string) []string {
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

	return tasks
}

func tampilkanTasks(tasks []string) {
	fmt.Println("\n📋 Daftar To-Do:")
	if len(tasks) == 0 {
		fmt.Println("- Belum ada task.")
		return
	}
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func saveTasksToFile(filename string, tasks []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Gagal menyimpan file:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		file.WriteString(task + "\n")
	}
}

func loadTasksFromFile(filename string) []string {
	var tasks []string

	file, err := os.Open(filename)
	if err != nil {
		return tasks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		task := strings.TrimSpace(scanner.Text())
		if task != "" {
			tasks = append(tasks, task)
		}
	}

	return tasks
}