package api_consul

type ResponseMembers struct {
	Name string `json:"name"`
	Addr string `json:"addr"`
	Port int    `json:"port"`
	Tags struct {
		Build   string `json:"build"`
		DC      string `json:"dc"`
		ID      string `json:"id"`
		Role    string `json:"role"`
		Segment string `json:"segment"`
		Vsn     string `json:"vsn"`
		Vsn_Max string `json:"vsn_max"`
		Vsn_Min string `json:"vsn_min"`
	}
	Status      int `json:"status"`
	ProtocolMin int `json:"protocolMin"`
	ProtocolMax int `json:"protocolMax"`
	ProtocolCur int `json:"protocolCur"`
	DelegateMin int `json:"delegateMin"`
	DelegateMax int `json:"delegateMax"`
	DelegateCur int `json:"delegateCur"`
}

type NodeData struct {
	ID            string `json:"ID"`
	Node          string `json:"Node"`
	Address       string `json:"Address"`
	Datacenter    string `json:"Datacenter"`
	TaggedAddress struct {
		LAN     string `json:"lan"`
		LANIPv4 string `json:"lan_ipv4"`
		WAN     string `json:"wan"`
		WANIPv4 string `json:"wan_ipv4"`
	} `json:"TaggedAddresses"`
	Meta struct {
		ConsulNetworkSegment string `json:"consul-network-segment"`
		ConsulVersion        string `json:"consul-version"`
	} `json:"Meta"`
	CreateIndex int `json:"CreateIndex"`
	ModifyIndex int `json:"ModifyIndex"`
}

type ServiceData struct {
	ID      string      `json:"ID"`
	Service string      `json:"Service"`
	Tags    []string    `json:"Tags"`
	Address string      `json:"Address"`
	Meta    interface{} `json:"Meta"`
	Port    int         `json:"Port"`
	Weights struct {
		Passing int `json:"Passing"`
		Warning int `json:"Warning"`
	} `json:"Weights"`
	EnableTagOverride bool `json:"EnableTagOverride"`
	Proxy             struct {
		Mode        string   `json:"Mode"`
		MeshGateway struct{} `json:"MeshGateway"`
		Expose      struct{} `json:"Expose"`
	} `json:"Proxy"`
	Connect     struct{} `json:"Connect"`
	PeerName    string   `json:"PeerName"`
	CreateIndex int      `json:"CreateIndex"`
	ModifyIndex int      `json:"ModifyIndex"`
}

type Data struct {
	Node     NodeData      `json:"Node"`
	Services []ServiceData `json:"Services"`
}
