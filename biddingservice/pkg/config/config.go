package config

type KaleidoAPIGateway struct {
	Port                   int
	KaleidoRestGatewayURL  string
	Gateway                string
	KaleidoAuthUsername    string
	KaleidoAuthPassword    string
	FromAddress            string
	ContractMainSourceFile string
	ContractClassName      string
}

func APIGateway() *KaleidoAPIGateway {
	return &KaleidoAPIGateway{
		Port:                   4000,
		KaleidoRestGatewayURL:  "u0wej9gatr-u0j7u0eb4f-connect.us0-aws.kaleido.io",
		Gateway:                "u0k4mie8a8",
		KaleidoAuthUsername:    "u0qltsa87h",
		KaleidoAuthPassword:    "awxwAmXq42HAloNdhWzXbrkLGcMjsZ2ME6wNIML4Ihs",
		FromAddress:            "0xaa3347224b6ca9098db1dcdbc799a2f876d8fdc5",
		ContractMainSourceFile: "simplestorage.sol",
		ContractClassName:      "simplestorage",
	}
}
