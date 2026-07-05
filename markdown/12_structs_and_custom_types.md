# Structs and Custom Types in Go

A struct groups related values into one type. A custom type gives a meaningful name and its own identity to an existing kind of value.

Together, they let a program describe its data in the language of its problem:

```go
type StudentID int

type Student struct {
	ID    StudentID
	Name  string
	Score int
}
```

Instead of passing an unrelated `int`, `string`, and `int`, the program can work with one `Student` value whose fields explain what each piece of data means.

## Defining a custom type

Use `type` to define a new type from an existing underlying type:

```go
type StudentID int
type EnrollmentStatus string
```

`StudentID` has `int` as its underlying type, but it is a distinct type. `EnrollmentStatus` is similarly distinct from `string`.

This distinction helps prevent unrelated values from being mixed accidentally:

```go
var id StudentID = 1042
var count int = 1042

// id = count // invalid: StudentID and int are different types
id = StudentID(count)
```

The expression `StudentID(count)` explicitly converts the `int` value to `StudentID`. A conversion is allowed when the types have compatible underlying representations.

Custom types can also make a small set of meaningful constants:

```go
type EnrollmentStatus string

const (
	StatusActive   EnrollmentStatus = "active"
	StatusInactive EnrollmentStatus = "inactive"
)
```

The type communicates what the string represents, while the constants provide consistent values. Go does not prevent another `EnrollmentStatus` string from being created, so validation may still be needed when input comes from a user or an external system.

## Defining a struct

A struct declaration lists its fields and their types:

```go
type Student struct {
	ID     StudentID
	Name   string
	Score  int
	Status EnrollmentStatus
}
```

Each field has a name and a type. Fields with the same type can be grouped:

```go
type Point struct {
	X, Y int
}
```

Separate field declarations are often clearer when the fields represent different ideas.

## Creating struct values

A struct literal creates a value and initializes its fields:

```go
student := Student{
	ID:     StudentID(1042),
	Name:   "Amina",
	Score:  84,
	Status: StatusActive,
}
```

Named fields make the code easy to read and allow fields to appear in any order. Fields that are omitted receive their zero values:

```go
student := Student{
	Name: "Amina",
}
```

It is also possible to provide values without field names:

```go
student := Student{StudentID(1042), "Amina", 84, StatusActive}
```

Named fields are usually safer. An unkeyed literal depends on the declaration's exact field order and must provide every field.

## Reading and updating fields

Use dot notation to select a field:

```go
fmt.Println(student.Name)
fmt.Println(student.Score)
```

Use the same notation in an assignment:

```go
student.Score = 91
student.Status = StatusInactive
```

Output:

```text
Amina
91
```

A field can itself have a custom type, a struct type, a slice, a map, or any other valid Go type.

## Struct zero values

The zero value of a struct contains the zero value of every field:

```go
var student Student

fmt.Printf("%q %d %q\n", student.Name, student.Score, student.Status)
```

Output:

```text
"" 0 ""
```

No constructor is required just to obtain a usable struct value. Design fields so their zero values are useful when practical. A struct containing maps, slices, or pointers may still need those fields initialized before certain operations, such as writing to a nil map.

## Nested structs

A struct can contain another struct:

```go
type Address struct {
	City    string
	Country string
}

type Student struct {
	Name    string
	Address Address
}
```

Initialize and access nested fields with the same literal and dot syntax:

```go
student := Student{
	Name: "Amina",
	Address: Address{
		City:    "Dubai",
		Country: "UAE",
	},
}

fmt.Println(student.Address.City)
```

Output:

```text
Dubai
```

Nesting keeps related data together without flattening every detail into one large struct.

## Structs are copied by value

Assigning a struct to another variable copies its fields:

```go
original := Student{Name: "Amina", Score: 84}
copied := original

copied.Name = "Mona"

fmt.Println(original.Name)
fmt.Println(copied.Name)
```

Output:

```text
Amina
Mona
```

The two struct values are independent. However, a field such as a slice, map, or pointer still refers to shared underlying data after being copied. The copy follows the normal behavior of each field type.

Passing a struct to a function also copies it. Pass a pointer when the function must update the caller's struct:

```go
func addBonus(student *Student, points int) {
	student.Score += points
}
```

Go automatically follows the struct pointer for field selection, so `student.Score` is shorthand for `(*student).Score`.

## Adding methods to a type

A method is a function associated with a type. Its receiver appears between `func` and the method name:

```go
func (student Student) Passed() bool {
	return student.Score >= 60
}
```

Call the method with dot notation:

```go
student := Student{Name: "Amina", Score: 84}
fmt.Println(student.Passed())
```

Output:

```text
true
```

A value receiver receives a copy of the value. It is a good fit for a method that reads data or calculates a result.

Use a pointer receiver when the method should change the original value:

```go
func (student *Student) AddBonus(points int) {
	student.Score += points
}

student.AddBonus(5)
fmt.Println(student.Score)
```

Output:

```text
89
```

Go automatically takes the address of an addressable value for this call. `student.AddBonus(5)` works even though `AddBonus` has a `*Student` receiver.

Methods can also be declared on custom types whose underlying type is not a struct:

```go
func (status EnrollmentStatus) IsActive() bool {
	return status == StatusActive
}
```

The receiver type must be defined in the same package as the method.

## Defined types and aliases

These declarations look similar but mean different things:

```go
type StudentID int // a new defined type
type Score = int   // an alias for int
```

`StudentID` is distinct from `int` and can have its own methods. `Score` is only another spelling of `int`; values of the two names are the same type.

Use a defined type when the program needs a distinct concept or type-specific methods. Aliases are mainly useful for compatibility and refactoring, not for creating a new domain concept.

## Comparing structs

A struct can be compared with `==` when every field in it is comparable:

```go
first := Point{X: 2, Y: 3}
second := Point{X: 2, Y: 3}

fmt.Println(first == second)
```

Output:

```text
true
```

A struct containing a slice, map, or function cannot be compared with `==` because those field types are not comparable.

## Common mistakes

- Mixing a defined type with its underlying type without an explicit conversion.
- Using an unkeyed struct literal and placing values in the wrong field order.
- Expecting an omitted field to receive anything other than its zero value.
- Expecting a copied struct or a value receiver to update the original struct.
- Using a pointer receiver for a method that only needs to read a small value.
- Assuming typed string constants prevent every invalid string from being converted.
- Comparing a struct that contains a slice, map, or function.
- Confusing a type alias with a new defined type.

Use structs to group data that belongs together, and use custom types to give important values a clear identity. Add methods when behavior naturally belongs with those types.
