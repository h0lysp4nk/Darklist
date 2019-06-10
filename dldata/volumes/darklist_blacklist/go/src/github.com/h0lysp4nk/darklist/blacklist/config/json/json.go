/* Package Type */
package jsonconf

/* Imports */
import (
	"os"
	"encoding/json"
)

/* BLACKLISTConfig Structure:
	This struct defines the allowed keys/values in the JSON file
*/
type BLACKLISTConfig struct {
    Blacklists []string `json:"blacklists"`
}

/* SMTPConfig Structure:
	This struct defines the allowed keys/values in the JSON file
*/
type SMTPConfig struct {
    SMTP struct {
        Host     string `json:"host"`
		Port string `json:"port"`
		Username string `json:"username"`
        Password string `json:"password"`
		SSLTLS bool `json:"ssltls"`
        STARTTLS bool `json:"starttls"`
    }
    Interval int `json:"interval"`
}

/* LoadSMTPConfiguration:
	Function will load a JSON file and return it
*/
func LoadSMTPConfiguration(file string) (config SMTPConfig, err error) {
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        return
    }
    dec := json.NewDecoder(configFile)
    err = dec.Decode(&config)
    return
}

/* LoadBLACKLISTConfiguration:
	Function will load a JSON file and return it
*/
func LoadBLACKLISTConfiguration(file string) (config BLACKLISTConfig, err error) {
    configFile, err := os.Open(file)
    defer configFile.Close()
    if err != nil {
        return
    }
    dec := json.NewDecoder(configFile)
    err = dec.Decode(&config)
    return
}