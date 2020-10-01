//Package filework needs for opening or creating .json
package filework

import (
	"log"
	"os"
)

//OpenTaskFile opens or creates json file
func OpenTaskFile() *os.File {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		log.Fatal("tasks file does not exist")
		return nil
	}
	file, err := os.Open("tasks.json")
	if err != nil {
		log.Fatal("Can't open this file")
	}
	return file
}
