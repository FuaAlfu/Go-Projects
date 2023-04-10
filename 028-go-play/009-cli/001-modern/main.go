package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"

)

var cfgFile string

func main() {
    rootCmd := &cobra.Command{
        Use:   "myapp",
        Short: "MyApp is a modern CLI application written in Go.",
        Long: `MyApp is a CLI application that demonstrates the use of the Cobra and Viper libraries
for creating a command-line interface with multiple commands and subcommands, and for reading
configuration files.`,
    }

    rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.myapp.yaml)")
    rootCmd.AddCommand(versionCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of MyApp",
    Long:  `All software has versions. This is MyApp's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("MyApp v0.1")
    },
}

func init() {
    cobra.OnInitialize(initConfig)
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    } else {
        home, err := os.UserHomeDir()
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }

        viper.AddConfigPath(home)
        viper.SetConfigType("yaml")
        viper.SetConfigName(".myapp")
    }

    if err := viper.ReadInConfig(); err != nil {
        fmt.Println("No configuration file found.")
    }
}
