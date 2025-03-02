package main

import (
	"fmt"
	"os"
)

func SaveData1(path string, date []byte) error {
	fp, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer fp.Close()

	_, err = fp.Write(date)
	if err != nil {
		return err
	}
	return fp.Sync()
}
func SaveData2(path string, data []byte) error {
	tmp := fmt.Sprintf("%s.tmp", path)

	fp, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0664)

	if err != nil {
		return err
	}

	defer func() {
		if closeErr := fp.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
		if err != nil {
			os.Remove(tmp)
		}
	}()

	if _, err = fp.Write(data); err != nil {
		return err
	}

	if err = fp.Sync(); err != nil {
		return err
	}

	return os.Rename(tmp, path)
}
func main() {
	fmt.Println("Hello, world")
	SaveData1("/Users/nkudinov/Work/sundb/file", []byte("Hello"))
}
