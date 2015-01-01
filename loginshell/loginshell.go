package loginshell

import (
	"fmt"

	"github.com/HirokazuMiyaji/profile-installer/command"
)

func Install() error {
	c := &command.Command{}

	c.Run("Set Zsh", `sudo sh -c "echo /usr/local/bin/zsh >> /etc/shells"`)
	c.Run("Change Zsh", "chsh -s /usr/local/bin/zsh")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}

func Uninstall() error {
	c := &command.Command{}

	c.Run("UnSet Zsh", `sudo sh -c "sed -e /\/usr\/local\/bin\/zsh//g /etc/shells`)
	c.Run("Change Bash", "chsh -s /bin/bash")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}
