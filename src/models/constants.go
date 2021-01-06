package models

import (
	"context"
	"net"
	"time"
)

// TCP_BUFFER_SIZE is the maximum packet size
const TCP_BUFFER_SIZE = 1024 * 64

// DEFAULT_RELAY is the default relay used (can be set using --relay)
var (
	DEFAULT_RELAY      = "croc.schollz.com"
	DEFAULT_RELAY6     = "croc6.schollz.com"
	DEFAULT_PORT       = "9009"
	DEFAULT_PASSPHRASE = "pass123"
)

func init() {
	var err error
	if err == nil {
		DEFAULT_RELAY += ":" + DEFAULT_PORT
	} else {
		DEFAULT_RELAY = ""
	}
	if err == nil {
		DEFAULT_RELAY6 = "[" + DEFAULT_RELAY6 + "]:" + DEFAULT_PORT
	} else {
		DEFAULT_RELAY6 = ""
	}
}
