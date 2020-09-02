package handlers

import (
	"fmt"
	"testing"
)

func TestGreeting(t *testing.T) {
	message := "Code.education Rocks!"

	r := Greeting(message)
	expectedMessage := fmt.Sprintf("<b>%s</b>", message)

	if r != expectedMessage {
		t.Errorf("Greeting(\"%s\") = %s; expected %s", message, r, expectedMessage)
	}
}
