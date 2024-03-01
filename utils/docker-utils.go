package utils

import (
	"bufio"
	"os"
	"strings"
)




func IsRunningInDocker() bool {

	if _, exists := os.LookupEnv("REMOTE_CONTAINERS_IPC"); exists {
		return true
	}

	if _, exists := os.LookupEnv("REMOTE_CONTAINERS_SOCKETS"); exists {
		return true
	}

	if _, exists := os.LookupEnv("REMOTE_CONTAINERS_DISPLAY_SOCK"); exists {
		return true
	}

	if _, exists := os.LookupEnv("REMOTE_CONTAINERS"); exists {
		return true
	}



	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}



	file, err := os.Open("/proc/1/cgroup")
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "docker") {
			return true
		}
	}

	return false
}
