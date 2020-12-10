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
	remoteClipFilePath := "/home/nix/Public/clip.txt"
	// saving the source machine clipboard text to a file on remote machine
	// AND copy source machine clipboard text to remote machine clipboard
	remoteCmd := fmt.Sprintf(
		"sed -i '1 i %s' %s && echo '%s' | tr -d '\n' | DISPLAY=:0 xsel -ib",
		string(clipTxt),
		remoteClipFilePath,
		string(clipTxt),
	)
	cmdSSHClip := exec.Command(
		"ssh",
		remoteMachineAddr,
		remoteCmd,
	)
	err = cmdSSHClip.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
