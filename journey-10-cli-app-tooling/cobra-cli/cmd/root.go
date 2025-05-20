package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "halo",
    Short: "Aplikasi CLI contoh",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Halo dari Cobra CLI!")
    },
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}
