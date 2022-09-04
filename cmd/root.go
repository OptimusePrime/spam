package cmd

import (
	"errors"
	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "spam",
	Short: "Spam is a simple CLI tool for sending messages to your friends.",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("requires at least 3 arguments")
		}
		if args[0] == "" {
			return errors.New("you must specify a message to send")
		}
		if _, err := strconv.Atoi(args[1]); err != nil {
			return errors.New("you must specify a number of times to send the message")
		}
		if _, err := strconv.Atoi(args[2]); err != nil {
			return errors.New("you must specify a delay between messages")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		message := args[0]
		times, _ := strconv.Atoi(args[1])
		delay, _ := strconv.Atoi(args[2])

		time.Sleep(5 * time.Second)
		for i := 0; i < times; i++ {
			robotgo.TypeStr(message)
			robotgo.KeyTap("enter")
			robotgo.MilliSleep(delay * 1000)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}
