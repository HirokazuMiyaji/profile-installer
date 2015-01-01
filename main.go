package main

import (
	"fmt"
	"os"

	"github.com/HirokazuMiyaji/profile-installer/commandlinetools"
	"github.com/HirokazuMiyaji/profile-installer/dotfiles"
	"github.com/HirokazuMiyaji/profile-installer/formula"
	"github.com/HirokazuMiyaji/profile-installer/homebrew"
	"github.com/HirokazuMiyaji/profile-installer/loginshell"
)

func main() {
	err := commandlinetools.Install()
	if err != nil {
		os.Exit(1)
	}

	err = homebrew.Install()
	if err != nil {
		os.Exit(1)
	}

	err = formula.Install()
	if err != nil {
		os.Exit(1)
	}

	err = dotfiles.Install()
	if err != nil {
		os.Exit(1)
	}

	err = loginshell.Install()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println("[DONE]")
}
