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
	Data    map[string]struct {
		Site struct {
			Subject  string   `json:"subject"`
			Altnames []string `json:"altnames"`
			RenewAt  int      `json:"renewAt"`
		} `json:"site"`
		Pems struct {
			Cert      string   `json:"cert"`
			Chain     string   `json:"chain"`
			Privkey   string   `json:"privkey"`
			Subject   string   `json:"subject"`
			Altnames  []string `json:"altnames"`
			IssuedAt  int64    `json:"issuedAt"`
			ExpiresAt int64    `json:"expiresAt"`
		} `json:"pems"`
	} `json:"data"`
}

type Proto struct {
	Service Service     `json:"service"`
	Data    interface{} `json:"data"`
}
