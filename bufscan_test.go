package bufscan

import (
	"testing"
	"os"
	"bufio"
	"strings"
)

func TestBufScan(t *testing.T) {
	f, err := os.Open("test.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var lines []string
	err = BufScan(reader, func(line string) error {
		lines = append(lines, line)
		return nil
	})

	if err != nil {
		t.Error(err)
	}

	if lines[0] != "test" {
		t.Errorf("got %v\nwant %v", lines[0], "test")
	}

	aArray := make([]string, 10000)
	for i := 0; i < 10000; i++ {
		aArray[i] = "a"
	}
	a10000 := strings.Join(aArray, "")
	if lines[1] != a10000 {
		t.Errorf("got %v\nwant %v", lines[1], a10000)
	}

	a2Array := make([]string, 100000)
	for i := 0; i < 100000; i++ {
		a2Array[i] = "a"
	}
	a100000 := strings.Join(a2Array, "")
	if lines[2] != a100000 {
		t.Errorf("got %v\nwant %v", lines[2], a100000)
	}

	if lines[3] != "test" {
		t.Errorf("got %v\nwant %v", lines[3], "test")
	}
}
