package practice

import "fmt"

type structStudentID int

type structEnrollmentStatus string

const (
	structStatusActive   structEnrollmentStatus = "active"
	structStatusInactive structEnrollmentStatus = "inactive"
)

type structAddress struct {
	City    string
	Country string
}

type structStudent struct {
	ID      structStudentID
	Name    string
	Score   int
	Status  structEnrollmentStatus
	Address structAddress
}

func StructsAndCustomTypesEx12() {
	fmt.Println("1. Define a custom type")
	fmt.Println("A defined type gives an underlying type a distinct identity.")
	fmt.Println("Code: type structStudentID int; id := structStudentID(1042)")
	id := structStudentID(1042)
	fmt.Printf("Result: id = %d, type = %T\n", id, id)

	fmt.Println("\n2. Group fields in a struct")
	fmt.Println("Named fields describe the parts of one value.")
	fmt.Println("Code: student := structStudent{ID: id, Name: \"Amina\", Score: 84, Status: structStatusActive}")
	student := structStudent{
		ID:     id,
		Name:   "Amina",
		Score:  84,
		Status: structStatusActive,
	}
	fmt.Printf("Result: ID = %d, Name = %q, Score = %d, Status = %q\n", student.ID, student.Name, student.Score, student.Status)

	fmt.Println("\n3. Read and update fields")
	fmt.Println("Dot notation selects a field for reading or assignment.")
	fmt.Println("Code: student.Score = 91")
	student.Score = 91
	fmt.Printf("Result: student.Score = %d\n", student.Score)

	fmt.Println("\n4. Use struct zero values")
	fmt.Println("Every field starts with the zero value of its type.")
	fmt.Println("Code: var guest structStudent")
	var guest structStudent
	fmt.Printf("Result: Name = %q, Score = %d, Status = %q\n", guest.Name, guest.Score, guest.Status)

	fmt.Println("\n5. Nest one struct inside another")
	fmt.Println("A field can hold another struct value.")
	fmt.Println(`Code: student.Address = structAddress{City: "Dubai", Country: "UAE"}`)
	student.Address = structAddress{City: "Dubai", Country: "UAE"}
	fmt.Printf("Result: student.Address.City = %q\n", student.Address.City)

	fmt.Println("\n6. Copy a struct value")
	fmt.Println("Changing a copied struct does not change the original struct.")
	fmt.Println(`Code: copied := student; copied.Name = "Mona"`)
	copied := student
	copied.Name = "Mona"
	fmt.Printf("Result: student.Name = %q, copied.Name = %q\n", student.Name, copied.Name)

	fmt.Println("\n7. Call a value-receiver method")
	fmt.Println("A value receiver is useful for reading or calculating.")
	fmt.Println("Code: func (student structStudent) passed() bool { return student.Score >= 60 }")
	fmt.Println("Code: student.passed()")
	fmt.Printf("Result: passed = %t\n", student.passed())

	fmt.Println("\n8. Call a pointer-receiver method")
	fmt.Println("A pointer receiver can update the original struct.")
	fmt.Println("Code: func (student *structStudent) addBonus(points int) { student.Score += points }")
	fmt.Println("Code: student.addBonus(5)")
	student.addBonus(5)
	fmt.Printf("Result: student.Score = %d\n", student.Score)

	fmt.Println("\n9. Use typed constants")
	fmt.Println("Typed constants give repeated domain values consistent names.")
	fmt.Println("Code: student.Status = structStatusInactive")
	student.Status = structStatusInactive
	fmt.Printf("Result: student.Status = %q, type = %T\n", student.Status, student.Status)
}

func (student structStudent) passed() bool {
	return student.Score >= 60
}

func (student *structStudent) addBonus(points int) {
	student.Score += points
}
