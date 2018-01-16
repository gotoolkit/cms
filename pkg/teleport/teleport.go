package teleport

import (
	"fmt"
	"net/http"
	"strings"
)

var teleportURL string

// Setup ...
func Setup(source string) {
	teleportURL = source
}

// SendMessage ...
func SendMessage(msg string) error {
	data := fmt.Sprintf(`{"chatId": "-304626054", "message": %q }`, msg)
	_, err := http.Post(teleportURL, "application/json", strings.NewReader(data))
	return err
}
