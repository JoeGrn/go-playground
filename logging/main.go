package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with microseconds")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my: ", log.LstdFlags)
	myLog.Println("custom logger")

	myLog.SetPrefix(("ohmy: "))
	myLog.Println("custom logger with new prefix")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf:", log.LstdFlags)
	buflog.Println("logging to buffer")

	fmt.Print("from bufLog: ", buf.String())

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	mySlog := slog.New(jsonHandler)
	mySlog.Info("info message")

	mySlog.Info("info message", "key", "value", "age", 25)
}
