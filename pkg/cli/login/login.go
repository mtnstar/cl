package login

import (
	"encoding/base64"
	"strings"
)

func NewCmdLogin(loginToken string) {
}

type LoginInfo struct {
	Username string
	Password string
	Host     string
}

func extractLogin(loginToken string) LoginInfo {
	s := strings.Split(loginToken, "@")
	creds, host := s[0], s[1]
	ds, _ := base64.StdEncoding.DecodeString(creds)
	sc := strings.Split(string(ds), ":")
	username, password := sc[0], sc[1]
	return LoginInfo{
		Username: username,
		Password: password,
		Host:     host,
	}
}
