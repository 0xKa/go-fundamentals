package practice

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type fileBill struct {
	id     int
	name   string
	amount float64
}

func (bill fileBill) format() string {
	return fmt.Sprintf("%d: %s - $%.2f", bill.id, bill.name, bill.amount)
}

func SavingFilesEx14() {
	workspace, setupErr := os.MkdirTemp("", "go-saving-files-*")
	if setupErr != nil {
		fmt.Printf("Could not prepare the practice: %v\n", setupErr)
		return
	}
	defer os.RemoveAll(workspace)

	bill := fileBill{id: 1, name: "Electricity", amount: 100}

	fmt.Println("1. Convert a value to bytes")
	fmt.Println("Files store bytes, so format the bill and convert its text.")
	fmt.Println("---")
	fmt.Println(`bill := fileBill{id: 1, name: "Electricity", amount: 100}`)
	fmt.Println("data := []byte(bill.format())")
	data := []byte(bill.format())
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> text = %q\n", string(data))
	fmt.Printf(">> bytes = %d\n", len(data))

	fmt.Println("\n2. Create a folder and save the file")
	fmt.Println("MkdirAll prepares the folder; WriteFile creates or replaces the file.")
	fmt.Println("---")
	fmt.Println(`folder := filepath.Join(workspace, "bills")`)
	fmt.Println("saveErr := os.MkdirAll(folder, 0o755)")
	fmt.Println("timestamp := time.Now().Unix()")
	fmt.Printf("%s\n", `filename := fmt.Sprintf("bill_%d_%d.txt", bill.id, timestamp)`)
	fmt.Println("path := filepath.Join(folder, filename)")
	fmt.Println("saveErr = os.WriteFile(path, data, 0o644)")
	folder := filepath.Join(workspace, "bills")
	saveErr := os.MkdirAll(folder, 0o755)
	if saveErr != nil {
		fmt.Println("---")
		fmt.Println("Output:")
		fmt.Printf(">> error = %v\n", saveErr)
		return
	}
	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("bill_%d_%d.txt", bill.id, timestamp)
	path := filepath.Join(folder, filename)
	saveErr = os.WriteFile(path, data, 0o644)
	fmt.Println("---")
	fmt.Println("Output:")
	if saveErr != nil {
		fmt.Printf(">> error = %v\n", saveErr)
		return
	}
	fmt.Printf(">> saved = %s\n", filename)
	fmt.Println(">> mode used = 0o644")

	fmt.Println("\n3. Read the file to verify it")
	fmt.Println("ReadFile confirms that the saved bytes contain the expected text.")
	fmt.Println("---")
	fmt.Println("savedData, err := os.ReadFile(path)")
	savedData, err := os.ReadFile(path)
	fmt.Println("---")
	fmt.Println("Output:")
	if err != nil {
		fmt.Printf(">> error = %v\n", err)
		return
	}
	fmt.Printf(">> contents = %q\n", string(savedData))
}
