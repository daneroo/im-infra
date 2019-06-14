package dc

import (
	"log"
)

// Cluster represents the charateristics for the cluster
type Cluster struct {
	Environment string
	AWSRegion   string
	NodeType    string
	// ClusterSize int
	DockerCloudAPIKey string
	DockerCloudFile   string

	// these are derived
	// ClusterName string // calculated
	// Stack       string
}

// Up brings up the nodeclster
func Up() {
	log.Printf("Cluster: up")
}
func (cluster Cluster) up() string {
	log.Printf("Cluster: up")
	return "ok"
}
