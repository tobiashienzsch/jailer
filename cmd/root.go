package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "jailer",
	Short: "jailer",
	Long:  `jailer https://github.com/tobiashienzsch/jailer`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Some info. Earth is not flat.")
		logrus.Warning("This is a warning")
		logrus.Error("Not fatal. An error. Won't stop execution")

	},
}

// Execute is the main entry point for the cli
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Viper config
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $PWD/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find pwd
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name "config" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("config")
		viper.AutomaticEnv() // read in environment variables that match

	}

	// Set log level
	// TODO: set via config.yml
	if viper.Get("verbose") == true {
		logrus.SetLevel(logrus.InfoLevel)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
