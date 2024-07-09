package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const red = "\033[0;31m"
const yellow = "\033[0;33m"
const green = "\033[0;32m"
const blue = "\033[0;34m"
const gray = "\033[0;90m"
const purple = "\033[0;35m"
const cyan = "\033[0;36m"
const boldRed = "\033[1;31m"
const boldYellow = "\033[1;33m"
const boldBlue = "\033[1;34m"
const boldGreen = "\033[1;32m"
const boldPurple = "\033[1;35m"
const boldWhite = "\033[1;37m"
const reset = "\033[0m"

func GetOS() string {
	return runtime.GOOS
}

func GetKernelVersion() string {
	cmd := exec.Command("uname", "-r")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func GetUptime() string {
	cmd := exec.Command("uptime", "-p")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func GetMemoryUsage() string {
	cmd := exec.Command("free", "-m")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	lines := strings.Split(string(output), "\n")
	if len(lines) < 2 {
		return "unknown"
	}
	fields := strings.Fields(lines[1])
	if len(fields) < 7 {
		return "unknown"
	}
	return fmt.Sprintf("%sMiB / %sMiB", fields[2], fields[1])
}

func GetHostName() string {
	cmd := exec.Command("hostname")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func GetShellVersion() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		return "unknown"
	}
	cmd := exec.Command(shell, "--version")
	output, err := cmd.Output()
	if err != nil {
		return "unknown"
	}
	return strings.TrimSpace(string(output))
}

func PrintASCIIArt() string {
	return `⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡤⠚⣷⠀⠀⣀⣤⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡞⣟⢀⡴⠋⠀⠀⣿⠖⠋⢀⡏⠀⠀⠀⡀⡀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢀⡀⡼⠀⢸⡟⡸⠀⠀⠀⠃⠀⠀⢸⡧⠜⠛⠛⣻⠃⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢺⢾⡃⠀⠈⣴⠁⢻⡀⠀⠀⢀⡠⠀⠀⠀⠀⢸⣇⣤⡀   %s⠀
⠀⠀⠀⠀⠀⠀⠀⠸⡜⠂⠀⠀⣟⠀⢸⠑⠀⠰⠁⠀⠀⠀⠀⠀⠛⠉⡼⠁   %s⠀
⠀⠀⠀⠀⠀⠀⠀⠈⣷⣾⣿⣿⣿⣿⣾⣶⣶⣤⣀⡀⢰⠕⠋⠀⠀⠸⠧⣤⡄  %s
⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣦⣔⠈⣤⣶⡚⠁   %s⠀
⠀⠀⠀⠀⣠⣶⡀⢸⡟⠿⡿⠿⡟⠻⠿⣿⣿⣿⣿⣿⣿⣿⠿⣿⠋⠁⠀⠀   %s⠀
⠀⠀⠀⢰⢧⡷⡿⢘⡎⠀⠀⠐⣶⢶⣲⠈⠙⠋⠉⠉⠁⡘⡯⣿⡶⣆⡀⠀   %s⠀
⠀⠀⠀⢾⢈⣼⣿⣤⣿⣶⣶⣶⣿⣿⣧⣤⣄⣀⣀⣤⣾⣿⣿⢯⢇⣿⢳
⠀⠀⠀⠈⠙⠿⢿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣌⣷⣬⠏⠀   %s⠀
⠀⠀⠀⠀⠀⠀⠀⠉⠙⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠻⠟⠋⠁⠀⠀⠀
⠀⠀⠀⠀⢀⣀⣀⡀⣰⣿⣿⣿⣿⣿⣿⣿⡿⠉⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣿⣿⣿⣿⢿⣿⣿⣿⣿⣿⣿⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢀⣤⣶⣴⣿⣿⣿⡧⠀⠉⠙⢿⣿⣿⣿⣿⣾⣶⣿⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠉⠛⠛⠿⣿⣿⡇⠀⠀⠀⠀⠻⣿⣿⣿⡿⠿⣿⣿⣿⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠁⠀⠀⠸⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀`
}

func main() {
	asciiArt := PrintASCIIArt()

	os_sys := GetOS()
	osFormatted := fmt.Sprintf("%sOS: %s%s", blue, reset, os_sys)

	kernel := GetKernelVersion()
	kernelFormatted := fmt.Sprintf("%sKernel: %s%s", blue, reset, kernel)

	uptime := GetUptime()
	uptimeFormatted := fmt.Sprintf("%sUptime: %s%s", blue, reset, uptime)

	memory := GetMemoryUsage()
	memoryFormatted := fmt.Sprintf("%sMemory: %s%s", blue, reset, memory)

	hostname := GetHostName()
	hostnameFormatted := fmt.Sprintf("%sHostname: %s%s", blue, reset, hostname)

	shell := GetShellVersion()
	formattedShell := fmt.Sprintf("%sShell: %s%s", blue, reset, shell)

	formattedColors := fmt.Sprintf("%s███%s███%s███%s███%s███%s███ %s", red, green, yellow, blue, purple, cyan, reset)

	fmt.Println("moefetch - A simple system information tool")
	fmt.Println("-------------------------------------------")

	fmt.Printf(asciiArt, osFormatted, kernelFormatted, uptimeFormatted, memoryFormatted, hostnameFormatted, formattedShell, formattedColors)

	fmt.Println("\n-------------------------------------------")
}
