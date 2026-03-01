package main

import "fmt"

type ValueError struct {
	message string
}

func (e ValueError) Error() string {
	return e.message
}

func checkBoolString(boolString string) error {
	if boolString == "true" || boolString == "false" {
		return nil
	}

	return ValueError{
		fmt.Sprintf("'%s' is not a boolean string.", boolString)}
}

func main() {
	values := []string{"true", "false", "hello"}
	for _, value := range values {
		err := checkBoolString(value)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
