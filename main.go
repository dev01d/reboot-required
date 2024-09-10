package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/speedata/optionparser"
)

var (
	version  string
	osType   string
	reboot   string = "\tReboot required"
	noReboot string = "\tNo reboot required"
)

func getPlatformInfo() {
	// Probe for OS type
	release, err := os.Open("/etc/os-release")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer release.Close()

	scanner := bufio.NewScanner(release)
	scanner.Scan()
	result := scanner.Text()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	if strings.Contains(result, "Ubuntu") {
		osType = "ubuntu"
	} else {
		osType = "rhel"
	}
}

func info() {
	if osType == "ubuntu" {
		ubuntuInfo()
	} else if osType == "rhel" {
		rhelInfo()
	} else {
		fmt.Printf("%s", osType)
	}
}

func ubuntuInfo() {
	if _, err := os.Stat("/var/run/reboot-required"); err == nil {
		color.Yellow(reboot)
	} else {
		color.Green(noReboot)
	}
	return
}

func rhelInfo() {
	newest := exec.Command("rpm -q kernel --qf '%{VERSION}-%{RELEASE}.%{ARCH}\n'| sort -V | tail -1")
	current := exec.Command("uname -r")
	if newest != current {
		color.Yellow(reboot)
	} else {
		color.Green(noReboot)
	}
}

func verbose() {
	if osType == "ubuntu" {
		ubuntuVerbose()
	} else if osType == "rhel" {
		rhelInfo()
	} else {
		fmt.Printf("%s", osType)
	}
}

func ubuntuVerbose() {
	if _, err := os.Stat("/var/run/reboot-required"); err == nil {
		rawfile, err := os.Open("/var/run/reboot-required.pkgs")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = rawfile.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		file, err := io.ReadAll(rawfile)
		if err != nil {
			log.Fatal(err)
		}

		color.Yellow(reboot)
		fmt.Print("Package(s) causing reboot:\n", string(file))
	} else {
		color.Green(noReboot)
	}
	return
}

func rhelVerbose() {
	newest := exec.Command("rpm -q kernel --qf '%{VERSION}-%{RELEASE}.%{ARCH}\n'| sort -v | tail -1")
	current := exec.Command("uname -r")
	if newest != current {
		color.Yellow(reboot)
		fmt.Print("Package(s) causing reboot:\n", newest)
	} else {
		color.Green(noReboot)
	}
}

func getVersion() {
	fmt.Printf("Version: %s\n", version)
}

func main() {
	getPlatformInfo()
	op := optionparser.NewOptionParser()
	op.Banner = "Usage: rr [OPTIONS]\n"
	op.On("-v", "--verbose", "Packages causing the need to reboot", verbose)
	op.On("--version", "Print version", getVersion)

	err := op.Parse()
	if err != nil {
		color.Red("Unknown option: %s\n", os.Args[1])
		op.Help()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		info()
	}
}
