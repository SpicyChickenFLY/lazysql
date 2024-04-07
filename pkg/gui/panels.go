package gui

import (
	"github.com/SpicyChickenFLY/lazysql/pkg/gui/panels"
	"github.com/jesseduffield/gocui"
)

func (gui *Gui) intoInterface() panels.IGui {
	return gui
}

func (gui *Gui) FilterString(view *gocui.View) string {
	if gui.State.Filter.panel != nil && gui.State.Filter.panel.GetView() != view {
		return ""
	}

	return gui.State.Filter.needle
}
