package proto

type Service int64

const (
	JoinService      Service = iota
	HeartBeatService Service = iota
)

type Ip struct {
	Address   string `json:"ip"`
	Swarm     bool
	Gateway   bool
	Signature string
}

type Stats struct {
	Storage  int64
	TotalOut int64
	Ingress  float64
	Egress   float64
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
}

type Proto struct {
	Service Service     `json:"service"`
	Data    interface{} `json:"data"`
}
