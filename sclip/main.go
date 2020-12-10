package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Get most recent clipboard text from source machine
	// In this case, source machine is strictly a Mac OSX.
	cmdEchoClip := exec.Command("pbpaste", ">&1")
	clipTxt, err := cmdEchoClip.Output()
	if err != nil {
		log.Fatalln(err)
	}

	remoteMachineAddr := "nix@192.168.1.9"
	remoteClipFilePath := "$HOME/Public/clip.txt"
	// saving the source machine clipboard text to a file on remote machine
	// AND copy source machine clipboard text to remote machine clipboard
	remoteCmd := fmt.Sprintf(
		"sed -i '1 i %s' %s; echo %s | xclip -sel clip",
		string(clipTxt),
		remoteClipFilePath,
		string(clipTxt),
	)
	cmdSshClip := exec.Command(
		"ssh",
		remoteMachineAddr,
		remoteCmd,
	)
	err = cmdSshClip.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
