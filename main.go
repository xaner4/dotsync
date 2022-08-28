package main

import (
	"errors"
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
	// var input string
	if err != nil {
		return false, err
	}
	for _, f := range files {
		// rp = realpath
		var rp string
		src := filepath.Join(path, f.Name())
		dst := filepath.Join(INSTALLDIR, f.Name())
		s, err := os.Stat(dst)
		Exists := !errors.Is(err, os.ErrNotExist)
		if err != nil {
			fmt.Println(err)
		}

		if Exists {
			rp, _ = os.Readlink(dst)
		}

		if src == rp {
			fmt.Printf("%s is already installed to %s \n", src, dst)
			// TODO: Check if src is folder, if folder, check if ther exists any files inside and call install on them! (WARNING: Recursion may happen)
			continue
		}

		if dst != rp {
			fmt.Printf("s: %s\n\n", s)
			fmt.Printf("%s is a symlink to %s \n", dst, rp)

			// fmt.Printf("Do you want to overwrite link %s [y/N]: \n", dst)
			// fmt.Scanln(&input)
			// input = strings.TrimSpace(input)
			// input = strings.ToLower(input)
			// switch input {
			// case "y", "yes", "ye", "yess":
			// 	fmt.Printf("Unlinking %s \n", dst)
			// 	os.Remove(dst)
			// }
		}
	}
	return true, nil
}

func backup(dst string) error {
	bd := filepath.Join(BACKUP, HOSTNAME, dst)
	if _, err := os.Stat(bd); os.IsNotExist(err) {
		err := os.Mkdir(bd, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func main() {
	fmt.Printf("Current Hostname: %s\n", HOSTNAME)
	p := getProfiles()
	i := askProfileInstall(p)
	for _, profile := range i {
		install(profile)
	}

}
