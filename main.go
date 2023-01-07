package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func readBinaryFile(filePath string) (*bytes.Buffer, error) {
	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new bufio.Reader object
	reader := bufio.NewReader(file)

	// Read the contents of the file into a bytes.Buffer
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func main() {
	fmt.Println("=== FTP File Uploader ===")

	if len(os.Args) < 4 {
		fmt.Println("Usage: ./ftpuploader.exe SERVER_IP LOCAL_FILE_PATH REMOTE_FILE_NAME")
		os.Exit(1)
	}

	// Connect to FTP
	ip := os.Args[1]
	filePath := os.Args[2]
	remoteName := os.Args[3]
	c, err := ftp.Dial(ip+":2121", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	// Login to FTP
	err = c.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	data, err := readBinaryFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T\n", data)

	err = c.Stor(remoteName, data)
	if err != nil {
		log.Fatal(err)
	}
}
