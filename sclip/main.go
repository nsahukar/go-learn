package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmdCpClip := exec.Command("pbpaste", ">&1")
	clipTxt, err := cmdCpClip.Output()
	if err != nil {
		log.Fatalln(err)
	}

	remoteCmd := fmt.Sprintf("sed -i '1 i %s' $HOME/Public/clip.txt", string(clipTxt))
	cmdNixClipWrite := exec.Command(
		"ssh",
		"nix@192.168.1.9",
		remoteCmd,
	)
	err = cmdNixClipWrite.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
