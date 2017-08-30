package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"gopkg.in/cheggaaa/pb.v1"
)

var packerurl = "https://releases.hashicorp.com/packer/0.10.0/packer_0.10.0_linux_amd64.zip"
var gitlabrepo = "http://gitlab.cisco.com/nvermand/packer-vesxi.git"

func downloadFile(filepath string, url string) error {

	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return err
		// exit if not ok
	}
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	fmt.Println("Downloading " + url)
	//Get file size
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	filesize := int(size)

	//Start new bar
	bar := pb.New(filesize).SetUnits(pb.U_BYTES)
	bar.Start()
	//Download file
	resp, err = http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//Create proxy reader
	rd := bar.NewProxyReader(resp.Body)
	// Writer the body to file
	_, err = io.Copy(out, rd)
	if err != nil {
		return err
	}
	bar.FinishPrint("Done")
	return nil
}

//ZipReader unzips file
func ZipReader(src string) ([]string, error) {

	filename := []string{}
	dir, _ := filepath.Split(src)
	// Open a zip archive for reading.
	fmt.Println("Opening " + src)
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	// Iterate through the files in the archive,
	for _, f := range r.File {
		src, err := f.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()
		dst, err := os.Create(dir + f.Name)
		if err != nil {
			return nil, err
		}
		defer dst.Close()
		fmt.Println("Uncompressing to " + dst.Name())
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, err
		}
		filename = append(filename, dst.Name())
	}
	return filename, nil

}

func execCmd(cmdName string, cmdArgs []string) (string, error)  {

	var (
		cmdOut []byte
		err error
	)
	if cmdOut, err = exec.Command(cmdname, cmdargs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error when executing command: ", err)
		os.Exit(1)

}

func main() {

	optionPtr := flag.String("vars", "vars.json", "Path for configuration file")

	flag.Parse()

	var (
		aexec  os.FileMode = 0755
	)

	dir, err := ioutil.TempDir("/tmp", "helper")
	if err != nil {
		log.Fatal(err)
	}
	packerpath := dir + "/packer.zip"
	err = downloadFile(packerpath, packerurl)
	if err != nil {
		log.Fatal(err)
	}
	filename, err := ZipReader(packerpath)
	if err != nil {
		log.Fatal(err)
	}

	os.Chmod(filename[0], aexec)

	cmdname := "git"
	cmdargs := []string{"clone", gitlabrepo}
	fmt.Println("Cloning repo", gitlabrepo)
	if cmdout, err = exec.Command(cmdname, cmdargs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "Error when executing command: ", err)
		os.Exit(1)
	}

	res := string(cmdout)
	fmt.Println(res)



}
