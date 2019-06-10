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
var debug = false

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

	/* Load configuration */
	if debug == true {
		fmt.Println("[!] Loading configuration...")
	}
	config, err := jsonconf.LoadConfiguration("/go/config.json")

	/* Ensure there are no errors */
	if err == nil {
		if debug == true {
			fmt.Println("[!] Success! Configuration loaded...")
			fmt.Println(config)
		}
	} else {
		fmt.Println("[!] Error! Something has gone wrong!")

		if debug == true {
			fmt.Println(err)
		}
		
	}

}