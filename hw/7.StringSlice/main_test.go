package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
)

type test struct {
	id     int
	input  string
	output string
}

func TestInput1(t *testing.T) {
	tests := []test{
		{
			id:     1,
			input:  "input/input1.txt",
			output: "3 2 1 8 8 2 1 1 2 3 4 4 7 8 9 1 2",
		},
	}

	for _, tt := range tests {
		file, err := os.Open(tt.input)
		if err != nil {
			t.Fatalf("file %s not found", tt.input)
		}

		data := make([]byte, 64)

		for {
			_, err := file.Read(data)
			if err == io.EOF {
				break
			}
		}

		cmd := exec.Command("go", "run", "main.go")
		stdin, err := cmd.StdinPipe()
		if err != nil {
			log.Fatal(err)
		}
		defer stdin.Close()
		io.WriteString(stdin, string(data))

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		outTrim := strings.TrimSpace(string(out))
		if outTrim != tt.output {
			t.Fatalf("test %d failed: expected %s, but was %s,", tt.id, string(out), tt.output)
		}
	}

}
