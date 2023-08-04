package dotenv

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// LoadEnvConfig takes data from a given file (f.e. .env) and parses it into a given structure.
// It is supposed to parse only string values
func LoadEnvConfig(filePath string, structure interface{}) (interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error when reading from file", err)
		return nil, err
	}

	envConf := parseEnvConfig(data)

	value := reflect.ValueOf(structure)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return nil, errors.New("structure must be a non-nil pointer")
	}

	value = value.Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		field := value.Field(i)
		fieldType := field.Type()

		if fieldType.Kind() != reflect.String {
			continue
		}

		field.SetString(envConf[fieldName])
	}

	return structure, nil
}

func parseEnvConfig(data []byte) map[string]string {
	splitted := strings.Split(string(data), "\n")
	envConf := make(map[string]string)

	for i := 0; i < len(splitted); i += 1 {
		splittedValue := strings.Split(splitted[i], "=")
		key := splittedValue[0]
		value := splittedValue[1]

		envConf[key] = value
	}

	return envConf
}
