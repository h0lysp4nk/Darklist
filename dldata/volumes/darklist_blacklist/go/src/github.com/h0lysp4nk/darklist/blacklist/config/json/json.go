/* Package Type */
package jsonconf

/* Imports */
import (
	"os"
	"encoding/json"
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
    }
    RunAt string `json:"runat"`
}

/* LoadConfiguration:
	Function will load a JSON file and return it
*/
func LoadConfiguration(file string) (config Config, err error) {
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        return
    }
    dec := json.NewDecoder(configFile)
    err = dec.Decode(&config)
    return
}