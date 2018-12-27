package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"../store"
)

var cfgFile string
var dataDir string
var jsonOutput bool

var Store store.Store

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&dataDir, "data-dir", "", "data directory (default is $HOME/.gotime)")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotime.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&jsonOutput, "json", "j", false, "output JSON instead of plain text")

	rootCmd.AddCommand(cmdProjects)
	rootCmd.AddCommand(cmdCreate)
	rootCmd.AddCommand(cmdReset)

	viper.BindPFlag("data-dir", rootCmd.PersistentFlags().Lookup("data-dir"))
}

func fatal(err ...interface{}) {
	fmt.Println(err...)
	os.Exit(1)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(home)
		viper.SetConfigName(".gotime")
	}

	viper.SetDefault("data-dir", filepath.Join(home, ".gotime"))

	if err := viper.ReadInConfig(); err != nil {
		// fmt.Println("Can't read config:", err)
		// os.Exit(1)
	}

	dataDir := viper.GetString("data-dir")
	os.MkdirAll(dataDir, 0700)
	Store, err = store.NewStore(dataDir)
	if err != nil {
		fatal(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gotime",
	Short: "gotime is a command line application to track time.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
