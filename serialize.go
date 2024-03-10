package tomllib

import (
	toml "github.com/pelletier/go-toml/v2"
	"log"
)

func Serialize(content []byte, modelInst interface{}) error {
	err := toml.Unmarshal(content, &modelInst)
	if err != nil {
		log.Println("Error unmarshalling TOML data:", err)
		return err
	}
	return nil
}
