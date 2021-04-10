package config

type NetworkConfig struct {
	GRPCServerAddress string
	GRPCPort          string
	HTTPServerAddress string
	HTTPPort          string
}

func GetNetworkConfig() *NetworkConfig {
	return &NetworkConfig{
		GRPCServerAddress: "localhost",
		GRPCPort:          "50051",
		HTTPServerAddress: "localhost",
		HTTPPort:          "8080",
	}
}
