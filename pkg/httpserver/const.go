package httpserver

import "time"

const (
	MAX_REQUEST_BODY_SIZE       = 5 // megabytes
	WRITE_TIMEOUT               = 100 * time.Second
	READ_TIMEOUT                = 100 * time.Second
	IDLE_TIMEOUT                = 100 * time.Second
	TCP_KEEPALIVE               = true
	MAX_CONNECTIONS_PER_IP      = 100
	MAX_REQUESTS_PER_CONNECTION = 100
)
