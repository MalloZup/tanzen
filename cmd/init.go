package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize your SAP-HA monitoring solution",
	Long: ` ==============================================
	        Initialize the server or client Monitoring
	        ==============================================`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var prometheus string
var server = &cobra.Command{
	Use:   "server",
	Short: "initialize the server SAP-HA monitoring solution",
	Long: ` ==============================================
                Initialize the server
                ==============================================`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Initializing server")

		fmt.Println("installing packages")

		fmt.Println("configuring Prometheus server")
		fmt.Printf("Prometheus server ip %s \n", prometheus)
	},
}

var client = &cobra.Command{
	Use:   "client",
	Short: "initialize the client SAP-HA monitoring solution",
	Long: ` ==============================================
                Initialize the client
                ==============================================`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Initializing client")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(client)

	// Here you will define your flags and configuration settings.

	server.Flags().StringVar(&prometheus, "prometheusIP", "", "prometheus server ip")
	server.MarkFlagRequired("prometheusIP")
	initCmd.AddCommand(server)

}
