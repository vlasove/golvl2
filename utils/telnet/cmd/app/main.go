package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/vlasove/materials/tasks_2/utils/telnet/internal/app/client"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "go-telnet"
)

var (
	help    bool
	version bool
	timeout time.Duration
)

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "Таймаут для соединения")
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")
}

func usage() {
	log.Printf(`ПРОСТАЯ РЕАЛИЗАЦИЯ TELNET (TCP). 
ИСПОЛЬЗОВАНИЕ : ./go-telnet [COMMAND] <host> <port>
ДОСТУПНЫЕ КОМАНДЫ:`)
	flag.PrintDefaults()
}

func showUsageAndExit(exitCode int) {
	usage()
	os.Exit(exitCode)
}

func showVersionAndExit(exitCode int) {
	log.Printf("%s . Version: '%s' Developed by: %s .",
		AppName,
		Version,
		Author,
	)
	os.Exit(exitCode)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if help {
		showUsageAndExit(0)
	}

	if version {
		showVersionAndExit(0)
	}
	args := flag.Args()
	if len(args) != 2 {
		showUsageAndExit(1)
	}

	host, port := args[0], args[1]
	runner := client.New(host, port, timeout)
	if err := runner.Start(); err != nil {
		log.Fatal(err)
	}
}
