package main

import (
	"./GraGO"
	"crypto/sha1"
	"flag"
	"fmt"
	"os/exec"

	// "fmt"
	"log"
	"strings"
)

func main() {
	stub := `sed -i "s/%s/%s/g" $(find . -type f)`
	mode := flag.String("mode", "hash", "type of args")
	flag.Parse()
	log.Printf("The mode is %s\n", *mode)
	log.Println(flag.Args())

	if *mode == "hash" {
		for _, hash := range flag.Args() {
			contents := GraGO.Deflate(hash)
			log.Printf("Sucessfully deflated commit %s\n", hash)
			var new_msg string
			new_msg = "new"

			split := strings.Split(contents, "\n")

			old_msg := split[len(split)-2]
			new_contents := strings.Replace(contents, old_msg, new_msg, 1)
			log.Printf("Sucessfully formed new msg\n")
			new_hash := GraGO.ConvertBytesString(sha1.Sum([]byte(new_contents)))
			log.Printf("Hash of new message is %s", new_hash)
			GraGO.WriteCommit(new_hash, new_contents)

			comstr := fmt.Sprintf(stub, hash, new_hash)
			cmd := exec.Command("/bin/sh", "-c", comstr)
			_, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
		}
	}

}
