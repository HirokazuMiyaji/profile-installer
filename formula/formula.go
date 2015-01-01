package formula

import (
	"fmt"

	"github.com/HirokazuMiyaji/profile-installer/command"
)

var cmdStrings = [...]string{
	"install openssl",
	"install readline",
	"install git",
	"install hub",
	"install hg",
	"install zsh",
	"install zsh-completions",
	"install zsh-syntax-highlighting",
	"install tree",
	"install the_silver_searcher",
	"install go",
	"install pyenv-virtualenv",
	"install rbenv",
	"install rbenv-default-gems",
	"install ruby-build",
	"install vim",
	"tap motemen/ghq",
	"install ghq",
	"peco/peco",
	"install peco",
	"tap caskroom/cask",
	"install brew-cask",
	"cask install atom",
	"cask install gitter",
	"cask install coda",
	"cask install intellij-idea",
	"cask install dropbox",
	"cask install box-sync",
	"cask install google-chrome",
	"cask install google-japanese-ime",
	"cask install iterm2",
	"cask install kobito",
	"cask install sequel-pro",
	"cask install slack",
	"cask install vagrant",
}

func Install() error {
	c := &command.Command{}

	c.Run("Brew Upgrade", "brew upgrade")
	for cmdString := range cmdStrings {
		c.Run("brew install", fmt.Sprintf("brew %s", cmdString))
	}

	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	c.Run("Brew cleanup", "brew cleanup")

	return nil
}

func Uninstall() error {
	c := &command.Command{}

	c.Run("Uninstall Brew Cask List", "brew cask list | xargs brew cask uninstall")
	c.Run("Uninstall Brew List", "brew list | xargs brew uninstall")
	if c.Error() != nil {
		fmt.Println(c.Error())
		return c.Error()
	}

	return nil
}
