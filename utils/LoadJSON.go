package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadJSON(filepath string, v any) ([]byte, error) {

	// Open our jsonFile
	jsonFile, err := os.Open(filepath)
	// if we os.Open returns an error then handle it
	if err != nil {
		// t.Log(err)
		return nil, err
	}

	// t.Log("Successfully Opened" + filepath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened file as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	errJ := json.Unmarshal(byteValue, v)

	if errJ != nil {
		// if error is not nil
		// print error
		fmt.Println(errJ)
	}

	return byteValue, nil
}
