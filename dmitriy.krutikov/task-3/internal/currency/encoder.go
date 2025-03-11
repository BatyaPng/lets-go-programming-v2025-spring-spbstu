package currency

import (
	"encoding/json"
	"os"
)

func SaveToJSON(currencies []Currency, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(currencies)
}