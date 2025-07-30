package cmd

import (
  "fmt"
  "os"

  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "callot",
  Short: "Callot is a CLI tool for margin and lot calculation",
  Long:  "Callot is a command-line tool that helps you manage margin settings and currency pairs for lot calculation.",
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

