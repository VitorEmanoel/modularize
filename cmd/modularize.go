package main

import (
	"dev.vitoremanoel.tech/mobolife/modularize/cmd/commands"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{

}

func main() {
	commands.NewCommand(rootCmd)
	err := rootCmd.Execute()
	if err != nil {
		log.Println("Error in starting modularize CLI. Error: ", err.Error())
	}
}
