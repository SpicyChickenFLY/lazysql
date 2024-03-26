package presentation

import "github.com/SpicyChickenFLY/lazysql/pkg/commands"

func GetNetworkDisplayStrings(network *commands.Network) []string {
	return []string{network.Network.Driver, network.Name}
}
