package utils

import (
	"encoding/json"
	"os"
)

func WriteFile(name string, data any) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(name, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
