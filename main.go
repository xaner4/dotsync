package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	HOME       = os.Getenv("HOME")
	INSTALLDIR = filepath.Join(HOME, "")
	DOTDIR     = filepath.Join(HOME, ".dotfiles")
	BACKUP     = filepath.Join(DOTDIR, ".backup")
	HOSTNAME   = getHostname()
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("error getting hostname: %v", err)
	}
	return hostname
}

func getProfiles() []string {
	profiles := []string{}
	files, err := ioutil.ReadDir(DOTDIR)
	if err != nil {
		log.Fatalf("error reading profiles: %v", err)
	}
	for _, f := range files {
		if f.IsDir() {
			if !strings.HasPrefix(f.Name(), ".") {
				profiles = append(profiles, f.Name())
			}
		}
	}
	return profiles
}

func askProfileInstall(profiles []string) []string {
	install := []string{}
	profilePaths := []string{}
	for i, p := range profiles {
		profilePaths = append(profilePaths, fmt.Sprintf("%s/%s", DOTDIR, p))
		fmt.Printf(" [%d]\t %s:\t\t %s/%s \n", i, p, DOTDIR, p)
	}
	fmt.Println("What profiles do you want to install (Seperate with comma): ")
	var input string
	fmt.Scanln(&input)
	installstr := strings.Split(input, ",")
	for _, i := range installstr {
		i, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal("Something went wrong getting profile")
		}
		install = append(install, profilePaths[i])
	}
	return install
}

func install(path string) (bool, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return false, err
	}
	for _, f := range files {

	return true, nil
}

func main() {
	fmt.Printf("Current Hostname: %s\n", HOSTNAME)
	p := getProfiles()
	i := askProfileInstall(p)
	fmt.Println(i)
	for _, profile := range i {
		install(profile)
	}

}
