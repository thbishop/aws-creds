package main

import (
	"fmt"
	"github.com/vaughan0/go-ini"
	"os"
	"os/user"
	"path"
	"sort"
	"strings"
)

func confFile() (ini.File, error) {
	path, err := confFilePath()
	if err != nil {
		return ini.File{}, err
	}

	return ini.LoadFile(path)
}

func confFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(usr.HomeDir, ".aws", "config"), nil
}

func credKeys() []string {
	return []string{
		"aws_access_key_id", "aws_secret_access_key", "aws_session_token",
	}
}

func credItems(profile string, f ini.File) []string {
	items := []string{}

	section := profileName(profile)

	for _, k := range credKeys() {
		if v, ok := f.Get(section, k); ok {
			items = append(items, fmt.Sprintf("%s=%s", strings.ToUpper(k), v))
		}
	}

	sort.Strings(items)
	return items
}

func profileName(profile string) string {
	if profile == "default" {
		return "default"
	}
	return "profile " + profile
}

func profileExists(profile string, f ini.File) bool {
	_, exists := f[profileName(profile)]
	return exists
}

func main() {
	options := parseCliArgs()

	f, err := confFile()
	if err != nil {
		os.Stderr.Write([]byte(fmt.Sprintf("Error with conf file: %s\n", err)))
		os.Exit(1)
	}

	if !profileExists(options.Profile, f) {
		os.Stderr.Write([]byte(fmt.Sprintf("Profile '%s' does not exist\n", options.Profile)))
		os.Exit(1)
	}

	for _, i := range credItems(options.Profile, f) {
		fmt.Printf("export %s\n", i)
	}
}
