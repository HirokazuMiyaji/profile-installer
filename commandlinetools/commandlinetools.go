package commandlinetools

import (
	"fmt"

	"github.com/HirokazuMiyaji/profile-installer/command"
)

func Install() error {
	c := &command.Command{}
	c.Run("Install Command Line Tools", "xcode-select --install")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}
