package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {

	type args struct {
		scenario         string
		dbFileName       func()
		expectedFileName string
	}

	tests := []args{
		{
			scenario: "value-provided",
			dbFileName: func() {
				os.Unsetenv("FileName")
				os.Setenv("FileName", "{\"test\"}")

			},
			expectedFileName: "test",
		},
		{
			scenario: "value-not-provided",
			dbFileName: func() {
				os.Unsetenv("FileName")
				os.Setenv("FileName", "{\"\"}")

			},
			expectedFileName: "",
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.dbFileName()

			err := LoadDBConfig()
			if err != nil {
				t.Errorf("%s got unexpected result: got %v want %v", test.scenario,
					FileName, test.expectedFileName)
			}
		})
	}
}
