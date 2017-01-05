package cligol

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

// ServerCmd is the server command handler
type ServerCmd struct{}

// Servers are an array of servers from the config/servers.json file
var Servers []Server

// Server is the type for server command
type Server struct {
	Name  string
	SSH   string
	Sites []Site
}

// Site is the type for server command
type Site struct {
	Name    string
	LogFile string
}

// Commands
var serverCommand *flag.FlagSet
var serverLogPtr *string
var serverLinesPtr *string

// Init is loading configurations and initialize main commands
func (serverCmd *ServerCmd) Init() {
	loadConfigs()
	initCommands()
}

// Parse gets commands from user and parse thems
func (serverCmd *ServerCmd) Parse(cmdArgs []string) {
	serverCommand.Parse(cmdArgs)

	if serverCommand.Parsed() {
		if *serverLogPtr != "" {
			printLogs(*serverLogPtr, serverCommand, serverLinesPtr)
			os.Exit(0)
		}

		// commands not found
		serverCommand.PrintDefaults()
		os.Exit(1)
	}
}

func getServerSite(siteName string) (Server, Site, error) {
	var serverType Server
	var siteType Site

	for _, server := range Servers {
		for _, site := range server.Sites {
			if site.Name == siteName {
				serverType = server
				siteType = site
				return serverType, siteType, nil
			}
		}
	}
	return serverType, siteType, errors.New("Site was not found")
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: \"%s\"\n", color.GreenString(strings.Join(cmd.Args, " ")))
}

func printLogs(siteName string, serverCommand *flag.FlagSet, serverLinesPtr *string) {
	server, site, err := getServerSite(siteName)
	if err != nil {
		serverCommand.PrintDefaults()
		os.Exit(1)
	}
	cmdName := "ssh"
	cmdArgs := strings.Split(server.SSH, " ")

	cmdArgs = append(cmdArgs, "tail", "-n", *serverLinesPtr, site.LogFile)
	cmd := exec.Command(cmdName, cmdArgs...)

	printCommand(cmd)

	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(out[:]))
}

func loadConfigs() {
	var _, currentFilePath, _, _ = runtime.Caller(0)
	var dirpath = path.Dir(currentFilePath)
	file, e := ioutil.ReadFile(dirpath + "/config/servers.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	err := json.Unmarshal(file, &Servers)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
}

func initCommands() {
	msg := "Show logs for project. Available projects:"
	for _, server := range Servers {
		for _, site := range server.Sites {
			msg = msg + "\n         - " + site.Name
		}
	}

	serverCommand = flag.NewFlagSet("server", flag.ExitOnError)
	serverLogPtr = serverCommand.String("log", "", msg)
	serverLinesPtr = serverCommand.String("n", "100", "Show last logs number of lines")
}
