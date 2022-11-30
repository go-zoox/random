package random

import "testing"

func TestClient(t *testing.T) {
	clientID := ClientID()
	clientSecret := ClientSecret()

	if len(clientID) != DEFAULT_CLIENT_ID_LENGTH {
		t.Errorf("Client ID length should be %d, got %d", DEFAULT_CLIENT_ID_LENGTH, len(clientID))
	}

	if len(clientSecret) != DEFAULT_CLIENT_SECRET_LENGTH {
		t.Errorf("Client Secret length should be %d, got %d", DEFAULT_CLIENT_SECRET_LENGTH, len(clientSecret))
	}
}
