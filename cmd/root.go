package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"wps.ktkt.com/monitor/tool-ssl-renewal/internal/setting"
	"wps.ktkt.com/monitor/tool-ssl-renewal/version"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tool-ssl-renewal command flags",
	Short: "Create or Binding ssl certificate to ucloud ULB",
	Long:  `创建和绑定ssl证书到ucloud的负载均衡`,
	//PersistentPreRunE: func(cmd *cobra.Command, args []string) error{
	//	err := setting.CheckAndParseConf(cfgFile)
	//	if err != nil {
	//		fmt.Fprintf(os.Stdout, "setting.CheckAndParseConf failed : %v", err)
	//		return err
	//	}
	//	return nil
	//},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	rootCmd.DisableSuggestions = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.tool-test.yaml)")
	rootCmd.MarkPersistentFlagRequired("config")

	rootCmd.Version = "\n" + version.BuildInfo.TplVersion()
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		//home, err := homedir.Dir()
		//cobra.CheckErr(err)

		// Search config in home directory with name ".tool-test" (without extension).
		viper.AddConfigPath("./data/")
		viper.SetConfigName("run")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// parse config
	conf := &setting.Config{}
	viper.Unmarshal(conf)
	setting.Conf = conf
}
