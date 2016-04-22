package runscope

import (
	"os"
	"testing"
)

var TOKEN = os.Getenv("RUNSCOPE_TOKEN")

func TestNewClient(t *testing.T) {
	client := NewClient(&Options{
		Token: TOKEN,
	})
	if client == nil {
		t.Fatalf("Client should not be nil")
	}
}
