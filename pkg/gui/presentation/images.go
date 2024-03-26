package presentation

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/commands"
	"github.com/SpicyChickenFLY/lazysql/pkg/utils"
)

func GetImageDisplayStrings(image *commands.Image) []string {
	return []string{
		image.Name,
		image.Tag,
		utils.FormatDecimalBytes(int(image.Image.Size)),
	}
}
