package errors

import (
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	err := ReadConfig()
	if err != nil {
		fmt.Println(err.Error())

		if findErr := FindError(err, &FileNotFoundError{}); findErr != nil {
			fmt.Println("Found FileNotFoundError , content : ", findErr.Error())
		}

		if findErr := FindError(err, &ReadFileError{}); findErr != nil {
			fmt.Println("Found ReadFileError , content : ", findErr.Error())
		}

		if findErr := FindError(err, &ReadConfigError{}); findErr != nil {
			fmt.Println("Found ReadConfigError , content : ", findErr.Error())
		}
	}
}

type ReadConfigError struct {
}

func (r ReadConfigError) Error() string {
	return "Read config error"
}

type ReadFileError struct {
}

func (r ReadFileError) Error() string {
	return "Read file error"
}

type FileNotFoundError struct {
	Filename string
}

func (c *FileNotFoundError) Error() string {
	return "Not found file  :" + c.Filename
}

func FindFile() error {
	return CreateStackFromError(&FileNotFoundError{Filename: "1.json"})
}

func ReadFile() error {
	err := FindFile()
	if err == nil {
		return nil
	} else {
		return WrapError(err, &ReadFileError{})
	}
}

func ReadConfig() error {
	err := ReadFile()
	if err == nil {
		return nil
	} else {
		return WrapError(err, &ReadConfigError{})
	}
}
