package simpleconfig

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"encoding/json"
)

func file_exists(path string) (bool, error) {
	_, error := os.Stat(path)
	if error == nil { return true, nil }
	if os.IsNotExist(error) { return false, nil }
	return false, error
}

func Load(filename string, object interface{}) (error) {
	file_exists, err := file_exists(filename)
	if err != nil {
		return err
	}
	if !file_exists {
		message := fmt.Sprintf("Filename %s does not exist.", filename)
		log.Fatal(message)
		return err
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &object)
	if err != nil {
		return err
	}
	return nil
}

func Save(filename string, object interface{}) (error) {
	json_data, err := json.MarshalIndent(&object, "", "\t")
	if err == nil {
		err = ioutil.WriteFile(filename, json_data, 0700)
		return err
	}
	return err
}
