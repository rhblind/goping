package lib

type CLICategory struct {
	Name string
}

type IPAddress struct {
	State   string `json:"state"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Netmask string `json:"netmask"`
	Network string `json:"network"`
}
