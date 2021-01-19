package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func WriteIndex(path string, line string, isAppend bool) error {
	if err := RotateIndexFile(path); err != nil {
		return err
	}
	flag := os.O_CREATE | os.O_WRONLY
	if isAppend {
		flag = flag | os.O_APPEND
	} else {
		flag = flag | os.O_TRUNC
	}
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}

func ReadLastIndex(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var cursor int64 = 0
	line := ""
	stat, _ := f.Stat()
	fSize := stat.Size()
	if fSize == 0 {
		return "", nil
	}
	for {
		cursor -= 1
		f.Seek(cursor, io.SeekEnd)
		char := make([]byte, 1)
		f.Read(char)
		if (cursor != -1) && (char[0] == 10 || char[0] == 13) {
			break
		}
		line = fmt.Sprintf("%s%s", string(char), line)
		if cursor == -fSize {
			break
		}
	}

	return line, nil
}

func RotateIndexFile(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	modTime := stat.ModTime()
	nowTime := time.Now()
	if modTime.Year() != nowTime.Year() && modTime.YearDay() == nowTime.YearDay() {
		return nil
	}
	os.Rename(path, fmt.Sprintf("%s.%d", path, modTime.Unix()))

	return nil
}

func main() {
	for i := 0; i < 0; i++ {
		if err := WriteIndex("./data.txt", strconv.Itoa(i), true); err != nil {
			fmt.Printf("write error: %v ", err)
			os.Exit(1)
		}
	}

	lastIndex, err := ReadLastIndex("./data.txt")
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	fmt.Print(lastIndex)
}
