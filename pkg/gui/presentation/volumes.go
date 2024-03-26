package presentation

import "github.com/SpicyChickenFLY/lazysql/pkg/commands"

func GetVolumeDisplayStrings(volume *commands.Volume) []string {
	return []string{volume.Volume.Driver, volume.Name}
}
