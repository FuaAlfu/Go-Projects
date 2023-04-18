package main

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

)

func main() {
    var rootCmd = &cobra.Command{
        Use:   "myapp",
        Short: "A simple CLI application",
        Long:  "A simple CLI application with multiple subcommands",
    }

    var cmd1 = &cobra.Command{
        Use:   "cmd1",
        Short: "A subcommand",
        Long:  "A subcommand that does something",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Command 1 executed")
        },
    }

    var cmd2 = &cobra.Command{
        Use:   "cmd2",
        Short: "Another subcommand",
        Long:  "Another subcommand that does something else",
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Println("Command 2 executed")
        },
    }

    rootCmd.AddCommand(cmd1, cmd2)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
