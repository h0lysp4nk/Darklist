/* Package Type */
package main

/* Imports:
	Additional Golang stuff
*/
import (
	"github.com/h0lysp4nk/darklist/blacklist/config/json"
	"fmt"
)

/* Debug Mode:
	Enable debugging information
*/
var debug = true

/* Main method:
	Self explanitory
*/
func main() {

	/* Print off the sexy banner */
	fmt.Println("    ____             __   ___      __  ")
	fmt.Println("   / __ \\____ ______/ /__/ (_)____/ /_ ")
	fmt.Println("  / / / / __ `/ ___/ //_/ / / ___/ __/ ")
	fmt.Println(" / /_/ / /_/ / /  / ,< / / (__  ) /_   ")
	fmt.Println("/_____/\\__,_/_/  /_/|_/_/_/____/\\__/   ")
	fmt.Println("                                       ")
	fmt.Println("")
	fmt.Println("===================================================")
	fmt.Println("")
	fmt.Println("")

	/* Load SMTP configuration */
	if debug == true {
		fmt.Println("[!] Loading smtp configuration...")
	}
	smtpConfig, smtpErr := jsonconf.LoadSMTPConfiguration("/go/configs/smtp.json")

	/* Ensure there are no errors */
	if smtpErr == nil {
		if debug == true {
			fmt.Println("[!] Success! Configuration loaded.!")
			fmt.Println(smtpConfig)
		}
	} else {
		fmt.Println("[!] Error! Something has gone wrong!")

		if debug == true {
			fmt.Println(smtpErr)
		}
		
	}

	/* Load BLACKLIST configuration */
	if debug == true {
		fmt.Println("[!] Loading blacklist configuration...")
	}
	blacklistConfig, blacklistErr := jsonconf.LoadBLACKLISTConfiguration("/go/configs/blacklist.json")

	/* Ensure there are no errors */
	if blacklistErr == nil {
		if debug == true {
			fmt.Println("[!] Success! Blacklist configuration loaded!")
			fmt.Println(blacklistConfig)
		}
	} else {
		fmt.Println("[!] Error! Something has gone wrong!")

		if debug == true {
			fmt.Println(blacklistErr)
		}
	}

}