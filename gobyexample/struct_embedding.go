package gobyexample

import "fmt"

type structEmbeddingBase struct {
	num int
}

func (b structEmbeddingBase) describe() string {
	return fmt.Sprintf("Base with number: %d", b.num)
}

type structEmbeddingContainer struct {
	structEmbeddingBase
	str string
}

func ShowStructEmbedding() {

	co := structEmbeddingContainer{
		structEmbeddingBase: structEmbeddingBase{
			num: 21,
		},
		str: "YOOOOO!",
	}

	fmt.Printf("Container: %v, String: %v\n", co.num, co.str)

	fmt.Println("Also number is:", co.structEmbeddingBase.num)

	fmt.Println("Describe method from base:", co.structEmbeddingBase.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer interface describe method:", d.describe())
}
