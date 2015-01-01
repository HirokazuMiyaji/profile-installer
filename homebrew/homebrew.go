package homebrew

import (
	"fmt"

	"github.com/HirokazuMiyaji/profile-installer/command"
)

func Install() error {
	c := &command.Command{}

	c.Run("Create Cache Folder", "mkdir ~/Library/Caches/Homebrew")
	c.Run("Install Brew", `ruby -e "$(curl -fsSL https://raw.github.com/mxcl/homebrew/go)"`)
	c.Run("Brew Update", "brew update")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}

func Uninstall() error {
	c := &command.Command{}

	c.Run("Brew prune", "brew prune")
	c.Run("Delete Brew Install Folder", "(ls -A1 $(brew --prefix) | xargs rm -rf)")
	c.Run("Delete Cache Folder", "rm -rf ~/Library/Caches/Homebrew ~/Library/Logs/Homebrew")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}
