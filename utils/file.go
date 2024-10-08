package utils

import (
	"encoding/json"
	"os"
)

func WriteFile(name string, data any) {
	b, _ := json.Marshal(data)
	_ = os.WriteFile(name, b, 0644)
}
