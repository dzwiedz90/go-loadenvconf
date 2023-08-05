package loadenvconf

import (
	"reflect"
	"testing"
)

type Config struct {
	USERNAME string
	PASSWORD string
}

func TestLoadEnvConfig(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		filePath    string
		expectedErr bool
		expectedCfg Config
	}{
		{
			name:     "ValidData",
			data:     "USERNAME=johndoe\nPASSWORD=secretpass",
			filePath: "testfiles/.test_env",
			expectedCfg: Config{
				USERNAME: "JohnDoe",
				PASSWORD: "p@$$w0rD",
			},
		},
		{
			name:        "Empty data",
			data:        "",
			filePath:    "testfiles/.empty_test_env",
			expectedErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			config := Config{}

			_, err := LoadEnvConfig(test.filePath, &config)

			if (err != nil) != test.expectedErr {
				t.Fatalf("Expected error: %v, but got: %v", test.expectedErr, err)
			}

			if !reflect.DeepEqual(config, test.expectedCfg) {
				t.Errorf("Expected %v, but got %v", test.expectedCfg, config)
			}
		})
	}
}

func TestParseEnvConfig(t *testing.T) {
	tests := []struct {
		name            string
		data            string
		expectedEnvConf map[string]string
	}{
		{
			name: "ValidData",
			data: "USERNAME=johndoe\nPASSWORD=secretpass",
			expectedEnvConf: map[string]string{
				"USERNAME": "johndoe",
				"PASSWORD": "secretpass",
			},
		},
		// Dodaj inne przypadki testowe, jeśli są wymagane
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			envConf := parseEnvConfig([]byte(test.data))

			if !reflect.DeepEqual(envConf, test.expectedEnvConf) {
				t.Errorf("Expected %v, but got %v", test.expectedEnvConf, envConf)
			}
		})
	}
}
