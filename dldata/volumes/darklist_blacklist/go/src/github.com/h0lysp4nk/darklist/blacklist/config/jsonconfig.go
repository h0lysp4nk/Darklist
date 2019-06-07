/* Package Type */
package config

/* Imports */
import (
	"fmt"
	"io"
)

/* Config Structure:
	This struct defines the allowed keys/values in the JSON file
*/
type Config struct {
    SMTP struct {
        Host     string `json:"host"`
		Port string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		SSLTLS bool `json:"ssltls"`
		STARTTLS bool `json:"starttls"`
    } `json:"database"`
    RunAt string `json:"runat"`
}

/* LoadConfiguration:
	Function will load a JSON file and return it
*/
func LoadConfiguration(file string) Config {
    var config Config
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        fmt.Println(err.Error())
    }
    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}