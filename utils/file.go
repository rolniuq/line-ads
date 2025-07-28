package utils

import (
	"os"

	"github.com/bytedance/sonic"
)

func WriteFile(name string, data any) error {
	b, err := sonic.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(name, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile[R any](fileName string) (*R, error) {
	b, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var r R
	if err := sonic.Unmarshal(b, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
