package config

import (
	"fmt"
	"testing"
)

func TestInitConf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestInitConf"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConf()
			conf := GetConf()
			fmt.Printf("%+v\n", conf)
			if conf.App.PageSize != 10 {
				t.Errorf("InitConf() error")
			}
		})
	}
}
