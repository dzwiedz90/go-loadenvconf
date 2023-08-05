package loadenvconf

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
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: %v", filePath)
		}
		return nil, fmt.Errorf("error when getting file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		return nil, errors.New("file is empty")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.New("Error when reading from file" + err.Error())
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
	if len(splitted) == 0 {
		return nil
	}
	envConf := make(map[string]string)

	for i := 0; i < len(splitted); i += 1 {
		splittedValue := strings.Split(splitted[i], "=")
		key := splittedValue[0]
		value := splittedValue[1]

		envConf[key] = value
	}

	return envConf
}
