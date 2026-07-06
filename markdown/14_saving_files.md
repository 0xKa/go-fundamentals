# Saving Files in Go

Saving a file means turning data into bytes and writing those bytes to a path.

The simplest function for this is `os.WriteFile`:

```go
data := []byte("Learn Go")
err := os.WriteFile("notes.txt", data, 0o644)
if err != nil {
	fmt.Println("Could not save file:", err)
}
```

`os.WriteFile` creates the file when it does not exist. If the file already exists, it replaces its contents.

## Convert values to bytes

`os.WriteFile` accepts a `[]byte`. Convert a string with `[]byte(text)`:

```go
type Bill struct {
	ID     int
	Name   string
	Amount float64
}

func (bill Bill) Format() string {
	return fmt.Sprintf("%d: %s - $%.2f", bill.ID, bill.Name, bill.Amount)
}

bill := Bill{ID: 1, Name: "Electricity", Amount: 100}
data := []byte(bill.Format())
```

The bytes contain the formatted bill text that will be stored in the file.

## Create the destination folder

Writing a file does not create missing parent directories. Create them first with `os.MkdirAll`:

```go
folder := "bills"

if err := os.MkdirAll(folder, 0o755); err != nil {
	return err
}
```

`MkdirAll` creates every missing directory in the path. It also succeeds when the directory already exists.

## Build file paths safely

Use `filepath.Join` instead of manually joining path parts with `/` or `\`:

```go
filename := "bill_1.txt"
path := filepath.Join("bills", filename)
```

This uses the correct path separator for the operating system.

A timestamp can make each filename unique:

```go
timestamp := time.Now().Unix()
filename := fmt.Sprintf("bill_%d_%d.txt", bill.ID, timestamp)
```

Unix time is the number of seconds since January 1, 1970 UTC.

## Save a value with a method

A save method can perform the complete operation and return errors to its caller:

```go
func (bill Bill) SaveToFile(folder string) (string, error) {
	if err := os.MkdirAll(folder, 0o755); err != nil {
		return "", fmt.Errorf("create bills folder: %w", err)
	}

	timestamp := time.Now().Unix()
	filename := fmt.Sprintf("bill_%d_%d.txt", bill.ID, timestamp)
	path := filepath.Join(folder, filename)
	data := []byte(bill.Format())

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return "", fmt.Errorf("write bill file: %w", err)
	}

	return path, nil
}
```

Call the method and handle its error:

```go
path, err := bill.SaveToFile("bills")
if err != nil {
	log.Println(err)
	return
}

fmt.Println("Saved:", path)
```

Returning an error lets the caller decide whether to display it, retry, or stop the program. `log.Fatal` is generally better kept near the top-level application code because it immediately terminates the program.

## File permissions

The final argument to `MkdirAll` and `WriteFile` is a permission mode:

```go
0o755 // directory: owner can write; everyone can read and enter
0o644 // file: owner can write; everyone can read
```

The `0o` prefix marks an octal number. Permission behavior varies by operating system, but these are common defaults for directories and ordinary files.

## Appending instead of replacing

Use `os.OpenFile` when new data should be added to the end:

```go
file, err := os.OpenFile(
	"notes.txt",
	os.O_APPEND|os.O_WRONLY|os.O_CREATE,
	0o644,
)
if err != nil {
	return err
}
defer file.Close()

_, err = file.WriteString("Another line\n")
return err
```

Files opened manually should be closed. Closing releases the operating-system resource and can report a final write error.

## Common mistakes

- Passing a string directly to `os.WriteFile` instead of converting it to `[]byte`.
- Forgetting to create the parent directory.
- Joining paths manually instead of using `filepath.Join`.
- Ignoring errors from directory and file operations.
- Expecting `os.WriteFile` to append instead of replace existing contents.
- Opening a file and forgetting to close it.
- Calling `log.Fatal` deep inside reusable code instead of returning an error.

For a small complete save operation: format the value, convert it to bytes, create its folder, build the path, call `os.WriteFile`, and handle the returned error.
