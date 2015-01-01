package dotfiles

import (
	"fmt"

	"github.com/HirokazuMiyaji/profile-installer/command"
)

func Install() error {
	c := command.Command{}

	dir := "~/src/github.com/HirokazuMiyaji"
	c.Run("Make Directory", fmt.Sprintf("[ ! -d %s ] && mkdir -p %s", dir, dir))
	c.Run("Git Clone", fmt.Sprintf("git clone https://github.com/HirokazuMiyaji/profiles.git %s/profiles", dir))
	c.Run("Git Pull/Submodule Update", fmt.Sprintf("cd %s/profiles && git pull && git submodule sync && git submodule update -i", dir))
	c.Run("Link", "cd %s/profiles && rake")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}
