package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/materials/tasks_2/utils/netcat/internal/app/netcater"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "netcat"
)

var (
	help    bool
	version bool
)

func init() {
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")
}

func usage() {
	log.Printf(`ПРОСТАЯ РЕАЛИЗАЦИЯ КЛИЕНТА NETCAT. 
ИСПОЛЬЗОВАНИЕ : ./netcat [COMMAND] <host> <port> <protocol>
ДОСТУПНЫЕ КОМАНДЫ:
А нету их :)`)
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
	if len(args) != 3 {
		showUsageAndExit(1)
	}

	host, port, protocol := args[0], args[1], args[2]
	nc := netcater.New(host, port, protocol)
	if err := nc.Start(); err != nil {
		log.Fatal(err)
	}
}
