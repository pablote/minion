package lib

import (
	"github.com/spf13/viper"
)

func GetPaths(args []string) []string {
	inventory := args[0]
	paths := viper.GetStringSlice(inventory)
	return paths
}
