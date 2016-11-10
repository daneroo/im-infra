package dc

import (
	"fmt"
	"log"

	"github.com/docker/go-dockercloud/dockercloud"
	"github.com/spf13/viper"
)

// Try is just a stub
func Try(verbose bool, dryRun bool) {
	fmt.Printf("Try verbose: %v dry-run: %v\n", verbose, dryRun)

	// https: //godoc.org/github.com/docker/go-dockercloud/dockercloud
	// dockercloud.Namespace = "yourOrganizationNamespace"
	dockercloud.User = viper.GetString("docker-cloud-user")
	dockercloud.ApiKey = viper.GetString("docker-cloud-apikey")

	fmt.Printf("User: %v ApiKey %v\n", dockercloud.User, dockercloud.ApiKey)
	fmt.Printf("Region: %v node-type: %v\n", viper.GetString("region"), viper.GetString("node-type"))

	aws, err := dockercloud.GetProvider("aws")
	if err != nil {
		log.Println(err)
		return
	}
	// log.Printf("AWS: %v\n", aws)
	log.Printf("AWS:Available %v\n", aws.Available)
	// log.Printf("AWS:Regions %v\n", aws.Regions)

	// regions, err := dockercloud.ListRegions()
	// log.Printf("Regions: %v\n", regions.Objects)

	// /api/infra/v1/region/aws/us-east-1/
	region, err := dockercloud.GetRegion("/api/infra/v1/region/aws/" + viper.GetString("region") + "/")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Region: %s\n", region.Resource_uri)
	// log.Printf("Region:Nodetypes %s\n", region.Node_types)

	nodetype, err := dockercloud.GetNodeType("aws", "t2.micro")
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("NodeType: %s\n", nodetype.Resource_uri)

	// clusters, err := dockercloud.ListNodeClusters()
	// if err != nil {
	// 	log.Printf("Error: %v\n", err)
	// } else {
	// 	log.Printf("Clusters: %v\n", clusters)
	// }

	var clusterRequest = dockercloud.NodeCreateRequest{
		Name:             "dan",
		Region:           region.Resource_uri,
		NodeType:         nodetype.Resource_uri,
		Target_num_nodes: 1,
	}

	nodecluster, err := dockercloud.CreateNodeCluster(clusterRequest)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Cluster: %v\n", nodecluster)
	if err = nodecluster.Deploy(); err != nil {
		log.Println(err)
		return
	}

	// containers, err := dockercloud.ListContainers()
	// if err != nil {
	// 	log.Printf("Error: %v\n", err)
	// } else {
	// 	log.Printf("Containers: %v\n", containers)
	// }
}
