package plugin

import "testing"

var logger = LogrusPlugin{}

func TestLoad(t *testing.T) {
	if err := logger.Load(); err != nil {
		t.Fatal(err)
	}
}
