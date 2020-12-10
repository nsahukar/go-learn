package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Get trailing bit of the ip address from command line arguments
	// If does not exists, use default.
	ipTrailBit := 9
	if len(os.Args) > 1 {
		ipTrailBit = int(os.Args[1])
	}
	remoteMachineAddr := fmt.Sprintf("nix@192.168.1.%d", ipTrailBit)
	fmt.Fprintln("SSHing to %s ...", remoteMachineAddr)

	// Get most recent clipboard text from source machine
	// In this case, source machine is strictly a Mac OSX.
	cmdEchoClip := exec.Command("pbpaste", ">&1")
	clipTxt, err := cmdEchoClip.Output()
	if err != nil {
		log.Fatalln(err)
	}

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
