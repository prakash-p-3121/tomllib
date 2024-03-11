package tomllib

import (
	toml "github.com/pelletier/go-toml/v2"
	"log"
)

func Serialize(content []byte, modelInstPtr interface{}) error {
	err := toml.Unmarshal(content, modelInstPtr)
	if err != nil {
		log.Println("Error unmarshalling TOML data:", err)
		return err
	}
	return nil
}
