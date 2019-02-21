package salt

import (
	"fmt"
	"log"
	"strings"
)

func DownloadFile(target Target, masterConfig MasterConfig, file string) (string, error) {
	stdout, _, err := Ssh(target, masterConfig, "--no-color", "cmd.run", fmt.Sprintf("'cat %s'", file))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Join(strings.Split(stdout, "\n    ")[1:], "\n"), nil
}