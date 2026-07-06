package main

import (
	"fmt"
	"go-fundamentals/internal/cli"
	"log"
	"os"
	"time"
)

type bill struct {
	id     int
	name   string
	amount float64
	paid   bool
}

func (b bill) format() string {
	return fmt.Sprintf("%v: %v - $%.2f", b.id, b.name, b.amount)
}

func (b bill) saveToFile() {
	fmt.Printf("Saving bill %d to file...\n", b.id)

	// timestamp to add to the file name
	timestamp := time.Now().Unix()

	// convert the bill to a byte slice
	data := []byte(b.format())

	folderName := "bills"

	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(folderName+"/"+fmt.Sprintf("bill_%d_%d.txt", b.id, timestamp), data, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Bill saved to file successfully.")
	}
}

func main() {

	b := bill{
		id:     1,
		name:   "Electricity",
		amount: 100.00,
		paid:   false,
	}

	b2 := bill{
		id:     2,
		name:   "Water",
		amount: 30.00,
		paid:   true,
	}

	fmt.Println(b.format())
	fmt.Println(b2.format())

	b.saveToFile()
	b2.saveToFile()

	cli.Start()
}
