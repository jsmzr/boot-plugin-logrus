package plugin

import "testing"

var logger = LogrusPlugin{}

func TestLoad(t *testing.T) {
	if err := logger.Load(); err != nil {
		t.Fatal(err)
	}
}

func TestOrder(t *testing.T) {
	if logger.Order() > 0 {
		t.Fatal("log order should be < 0")
	}
}

func TestEnabled(t *testing.T) {
	if !logger.Enabled() {
		t.Fatal("log enabled must be true")
	}
}
