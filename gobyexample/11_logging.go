package gobyexample

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func ShowLogging() {
	// normal logger
	log.Println("standard logger")

	// with flags,
	// for example, microseconds
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// with file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// custom logger
	mylog := log.New(os.Stdout, "my-logger:", log.LstdFlags)
	mylog.Println("from mylog")

	// custom prefix
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// custom logger with buffer
	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("hello from buflog")
	buflog.Println("hello")
	// print the buffer content
	fmt.Print("from buflog:", buf.String())

	// using slog (structured logger) to log in JSON format
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

	// can also customize log key-value pairs
	myslog.Info("hello again", "key", "val", "age", 25)
}
