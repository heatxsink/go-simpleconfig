package simpleconfig

import (
	"os"
	"fmt"
	"testing"
)

var (
	test_filename = "config/test.config"
	test2_filename = "config/test2.config"
)

type SampleAppConfig struct {
	Key string `json:"key"`
	Secret string `json:"secret"`
	Hostname string `json:"hostname"`
	Port int `json:"port"`
	DebugFlag bool `json:"debug_flag"`
}

func TestLoad(t *testing.T) {
	fmt.Println("Load")
	sac := SampleAppConfig {}
	err := Load(test_filename, &sac)
	if err != nil {
		t.Errorf("Config LoadFromFile Error: %g", err)
	}
	fmt.Println(sac)
}

func TestSave(t *testing.T) {
	fmt.Println("Save")
	sac := SampleAppConfig {}
	sac.Key = "abc"
	sac.Secret = "123"
	sac.Hostname = "4.2.2.2"
	sac.Port = 69
	sac.DebugFlag = true
	err := Save(test2_filename, &sac)
	if err != nil {
		t.Errorf("Config SaveToFile Error: %g", err)
	}
	os.Remove(test2_filename)
}
