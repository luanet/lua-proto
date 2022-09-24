package proto

import (
	"database/sql/driver"
	"encoding/json"
)

type Service int64

const (
	JoinService      Service = iota
	HeartBeatService Service = iota
)

type Certificate map[string]interface{}

func (a Certificate) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *Certificate) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Ip struct {
	Address   string `json:"ip"`
	Swarm     bool
	Gateway   bool
	Signature string
}

type Stats struct {
	Storage int64   // total disk space usage
	In      int64   // total bytes in since last sync
	Out     int64   // total bytes out since last sync
	Ingress float64 // ingress bytes/second
	Egress  float64 // egress bytes/second
}

type HeartBeatReq struct {
	Stats Stats
}

type HeartBeatRes struct {
	Success bool `json:"success"`
}

type JoinReq struct {
	Address   string `json:"address"`
	Ipv4      Ip     `json:"ipv4"`
	Ipv6      Ip     `json:"ipv6"`
	Signature []byte `json:"signature"`
	Expires   int64  `json:"expires"`
}

type JoinRes struct {
	Success bool   `json:"success"`
	Message string `json:message`
	Certs   map[string]Certificate
}

type Proto struct {
	Service Service     `json:"service"`
	Data    interface{} `json:"data"`
}
