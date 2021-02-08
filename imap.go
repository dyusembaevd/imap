package imap

import (
	"strings"

	"github.com/mohito22/tcp"
)

type IMAPConfig struct {
	TCPConfig       *tcp.TCPConfig
	PortIsAvailable bool
	LOGINAuth       bool
	PLAINAuth       bool
}

func CheckPort143(hostname string) IMAPConfig {
	return CheckPort(hostname, "143")
}

func CheckPort993(hostname string) IMAPConfig {
	return CheckPort(hostname, "993")
}

func CheckPort(hostname, port string) IMAPConfig {
	tcp := tcp.NewConfig()
	imap := NewIMAPConfig()
	imap.TCPConfig = tcp
	if err := tcp.Connect(hostname, port); err != nil {
		return imap
	}

	resp := tcp.ReadTCPMessage()
	if len(resp) != 0 {
		imap.PortIsAvailable = true
	}

	if strings.Contains(strings.ToLower(string(resp)), "auth=login") {
		imap.LOGINAuth = true
	}

	if strings.Contains(strings.ToLower(string(resp)), "auth=plain") {
		imap.PLAINAuth = true
	}

	return imap
}

func NewIMAPConfig() IMAPConfig {
	return IMAPConfig{}
}
