package errors

import (
	stdErr "errors"
	"fmt"
	"io"
	"reflect"
)

// errors 记录堆栈 和 堆叠的错误
type errors struct {
	*stack
	errorArray *[]error
}

func (e *errors) printErrorArray() {
	for _, error := range *e.errorArray {
		fmt.Println(error.Error())
	}
}

func (e *errors) Error() string {
	return fmt.Sprintf("%+v", e)
}

func (e *errors) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Println("------error stack------")
			e.printErrorArray()
			fmt.Println("------function stack------")
			e.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, e.Error())
	}
}

func CreateStackFromMessage(msg string) error {
	errArr := make([]error, 0, 32)
	errArr = append(errArr, stdErr.New(msg))
	return &errors{
		stack:      callers(),
		errorArray: &errArr,
	}
}

func CreateStackFromError(err error) error {
	errArr := make([]error, 0, 32)
	errArr = append(errArr, err)
	return &errors{
		stack:      callers(),
		errorArray: &errArr,
	}
}

func WrapError(payload error, err error) error {
	//判断 payload 必须是 errors类型 否则不允许使用
	if reflect.TypeOf(payload) != reflect.TypeOf(&errors{}) {
		panic("Incorrect payload type , please use CreateStackFromError or CreateStackFromMessage function to create!")
	}

	payloadErrors := payload.(*errors)
	*payloadErrors.errorArray = append(*payloadErrors.errorArray, err)

	return payload
}

func FindError(payload error, err error) error {
	//判断 payload 必须是 errors类型 否则不允许使用
	if reflect.TypeOf(payload) != reflect.TypeOf(&errors{}) {
		panic("Incorrect payload type , please use CreateStackFromError or CreateStackFromMessage function to create!")
	}

	payloadErrors := payload.(*errors)

	for _, error := range *(payloadErrors.errorArray) {
		if reflect.TypeOf(error) == reflect.TypeOf(err) {
			return error
		}
	}
	return nil
}
