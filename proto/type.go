package proto

import (
	metrics "github.com/libp2p/go-libp2p-core/metrics"
)

type Service int64

const (
	JoinService      Service = iota
	HeartBeatService Service = iota
)

type HeartBeatReq struct {
	Stats metrics.Stats
}

type HeartBeatRes struct {
	Success bool `json:"success"`
}

type JoinReq struct {
	Address   string `json:"address"`
	Signature []byte `json:"signature"`
	Expires   int64  `json:"expires"`
}

type JoinRes struct {
	Success bool   `json:"success"`
	Message string `json:message`
}

type Proto struct {
	Service Service     `json:"service"`
	Data    interface{} `json:"data"`
}
