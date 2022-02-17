package config

type Config struct {
	RemoteCluster bool `short:"r" long:"remote" description:"Connect to Remote Cluster instead of Local"`
}
