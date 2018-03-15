package fifo

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Writer struct {
}

func NewWriter() *Writer {
	return &Writer{}
}

func (f *Writer) Write(content string) (string, error) {
	tempDir, err := ioutil.TempDir("", "reconfigure-pipeline")
	if err != nil {
		return "", err
	}

	fifoPath := filepath.Join(tempDir, "fifo")
	// err = syscall.Mkfifo(fifoPath, 0600)
	// if err != nil {
	// 	return "", err
	// }

	// go func() {
	fi, err := os.Create(fifoPath)
	defer fi.Close()

	if err != nil {
		log.Fatal(err)
	}

	fi.WriteString(content)
	// }()

	return fifoPath, nil
}
