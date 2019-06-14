package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Options : Global options
type Options struct {
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

var defaultOptions = &Options{
	Environment:       "dev",
	AWSRegion:         "us-east-1",
	NodeType:          "t2.micro",
	DockerCloudAPIKey: "",
	DockerCloudFile:   "docker-cloud.yml",
}

var cfgFile string
var dryRun bool
var verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "im-infra",
	Short: "A tool for managing docker-cloud infrastucture provisioning and deployment",
	Long: `A longer description that explains managing
docker-cloud infrastucture provisioning and deployment`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
		json, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
		fmt.Printf("Root:PreRun Viper: %v\n", string(json))

		// dc.Try(viper.GetBool("verbose"), viper.GetBool("dry-run"))
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is  [./|$HOME]/.im-infra.[json|toml|yaml])")

	RootCmd.PersistentFlags().BoolP("dry-run", "", false, "Inhibit all actions")
	viper.BindPFlag("dry-run", RootCmd.PersistentFlags().Lookup("dry-run"))

	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	viper.BindPFlag("verbose", RootCmd.PersistentFlags().Lookup("verbose"))

	RootCmd.PersistentFlags().StringP("environment", "e", defaultOptions.Environment, "environment name")
	viper.BindPFlag("environment", RootCmd.PersistentFlags().Lookup("environment"))

	RootCmd.PersistentFlags().StringP("region", "r", defaultOptions.AWSRegion, "AWS Region")
	viper.BindPFlag("region", RootCmd.PersistentFlags().Lookup("region"))

	RootCmd.PersistentFlags().StringP("node-type", "t", defaultOptions.NodeType, "AWS Instance type")
	viper.BindPFlag("node-type", RootCmd.PersistentFlags().Lookup("node-type"))

	RootCmd.PersistentFlags().StringP("docker-cloud-file", "f", defaultOptions.DockerCloudFile, "docker cloud stack file")
	viper.BindPFlag("docker-cloud-file", RootCmd.PersistentFlags().Lookup("docker-cloud-file"))

	RootCmd.PersistentFlags().String("docker-cloud-apikey", defaultOptions.DockerCloudFile, "docker cloud stack file")
	viper.BindPFlag("docker-cloud-apikey", RootCmd.PersistentFlags().Lookup("docker-cloud-apikey"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	// see https://sourcegraph.com/github.com/spf13/cobra/-/blob/cobra/cmd/root.go
	// see https://sourcegraph.com/github.com/TheThingsNetwork/ttn/-/blob/ttnctl/cmd/root.go#L0
	viper.SetConfigName(".im-infra") // name of config file (without extension)
	// Path order is ordered...
	viper.AddConfigPath(".")     // optionally look for config in the working directory
	viper.AddConfigPath("$HOME") // adding home directory as first search path
	viper.AutomaticEnv()         // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
