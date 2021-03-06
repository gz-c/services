package common

type Config struct {
	Secret SecretConfig `json:"secret"`
	Public PublicConfig `json:"public"`
}

type SecretConfig struct {
	MasterSecKey     string `json:"masterPrivateKey"`
	GenesisSignature string `json:"genesisSignature"`
}

type PublicConfig struct {
	MasterPubKey string `json:"masterPublicKey"`

	GenesisBlock GenesisBlockConfig `json:"genesisBlock"`

	CoinCode string `json:"coinCode"`

	Port             int `json:"port"`
	WebInterfacePort int `json:"webInterfacePort"`
	RPCInterfacePort int `json:"rpcInterfacePort"`

	DataDirectory string `json:"dataDirectory"`
	LogFmt        string `json:"logFmt"`
}

type GenesisBlockConfig struct {
	Address    string `json:"address"`
	CoinVolume uint64 `json:"coins"`
	Timestamp  uint64 `json:"timestamp"`
	BodyHash   string `json:"bodyHash"`
	HeaderHash string `json:"headerHash"`
}
