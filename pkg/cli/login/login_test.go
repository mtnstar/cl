package login

import "testing"

func TestExtractLoginValid(t *testing.T) {
	validLoginToken := "YWxpY2UtYWJjNDI6dG9rZW40Mg==@cryptopus.example.com"
	loginInfo := extractLogin(validLoginToken)

	if loginInfo.Username != "alice-abc42" {
		t.Errorf("Username was incorrect, got: %s, want: %s", loginInfo.Username, "alice-abc42")
	}
	if loginInfo.Password != "token42" {
		t.Errorf("Password was incorrect, got: %s, want: %s", loginInfo.Password, "token42")
	}
	if loginInfo.Host != "cryptopus.example.com" {
		t.Errorf("Host was incorrect, got: %s, want: %s", loginInfo.Host, "cryptopus.example.com")
	}
}
