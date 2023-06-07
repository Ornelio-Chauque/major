package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	major "github.com/Ornelio-Chauque/major/internal"
	"github.com/spf13/cobra"
)

var command string
var username string

func init() {
	majorCmd.Flags().StringVarP(&command, "command", "c", "", "the command and its arguments the you want to run")
	majorCmd.Flags().StringVarP(&username, "user", "u", "root", "The user to run the process")
	majorCmd.MarkFlagRequired("command")
}

var majorCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		command:= parsingCommand(command)
		major.Run(command, username)
	},

	Use:   "major",
	Long:  "major a sudo like tool",
	Short: "Run command as root or any other user",
}

func Execute() {
	if err := majorCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}


func parsingCommand(s string) []string{
	s = strings.TrimSpace(s)
	if len(s) <= 0{
		fmt.Println("Invalid command flag")
	}
	regex := regexp.MustCompile(` +`)
	newCommand := regex.Split(s, -1)
	fmt.Println(newCommand)
	return newCommand;
}
