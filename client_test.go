package runscope

import "testing"

func TestNewClient(t *testing.T) {
	client := NewClient(Options{})
	if client == nil {
		t.Fatalf("Client should not be nil")
	}
}
