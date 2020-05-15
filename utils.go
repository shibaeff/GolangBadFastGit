package GraGO

import (
	"bytes"
	"fmt"
	git "gopkg.in/libgit2/git2go.v27"
	"log"
	"os"
	"os/exec"
	"path"
)

func convertStringBytes(s string) []byte {
	decoded := fmt.Sprintf("%x", s)
	return []byte(decoded)
}

func ConvertBytesString(bytes [20]byte) string {
	return fmt.Sprintf("%x", bytes)
}

func getOidFromPos(label string) (*git.Oid, error) {
	out, err := exec.Command("git", "rev-parse", label).Output()
	if err != nil {
		return nil, err
	}
	s := string(out)
	log.Printf(s)
	var oid *git.Oid
	oid = git.NewOidFromBytes(convertStringBytes(s))
	return oid, nil
}

func getOidFromHash(hash string) *git.Oid {
	return git.NewOidFromBytes(convertStringBytes(hash))
}

func getHashFromS(label string) (*git.Oid, error) {
	out, err := exec.Command("git", "rev-parse", label).Output()
	if err != nil {
		return nil, err
	}
	s := string(out)
	log.Printf(s)
	var oid *git.Oid
	oid = git.NewOidFromBytes(convertStringBytes(s))
	return oid, nil
}

func Deflate(hash string) string {
	prefix := hash[:2]
	filename := hash[2:]
	filepath := path.Join("./.git/objects", prefix, filename)
	command := exec.Command("zlib-flate", "-uncompress")
	command.Stdin, _ = os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	var out bytes.Buffer
	command.Stdout = &out
	command.Run()
	return out.String()
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func WriteCommit(hash string, cont string) {

	dirPath := path.Join(".git/objects", hash[:2])

	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directory for commit \n")
		return
	}
	log.Println("Created dir for commit")
	tempPath := path.Join(dirPath, "temp")
	file, err := os.OpenFile(tempPath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error creating commit object file")
		return
	}
	_, err = file.WriteString(cont)
	if err != nil {
		log.Printf("error wirting to temp file\n")
		return
	}
	err = file.Close()
	if !fileExists(tempPath) {
		log.Println("Temp file was not created!")
		return
	}
	log.Println("Created temp file with commit contents")
	command := exec.Command("zlib-flate", "-compress")
	command.Stdin, err = os.OpenFile(tempPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Println("Error opening temp file")
		log.Println(err)
		return
	}
	command.Stdout, err = os.OpenFile(path.Join(dirPath, hash[2:]), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("Error creating commit object file")
		return
	}
	err = command.Run()
	if err != nil {
		log.Println(err)
		return
	}

	if !fileExists(path.Join(dirPath, hash[2:])) {
		log.Println("Problems creating commit file occurred")
	}
	log.Println("Successfully created commit object file!!!")

}
