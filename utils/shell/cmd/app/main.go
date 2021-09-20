package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/materials/tasks_2/utils/shell/internal/app/command"
	"github.com/vlasove/materials/tasks_2/utils/shell/internal/app/shell"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "shell"
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
	log.Printf(`ПРОСТАЯ ОБОЛОЧКА. ДЛЯ ВЫХОДА ИСПОЛЬЗУЙТЕ \exit.`)
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

	s := shell.New()
	commander := command.New(s, os.Stdout, os.Stdin)
	if err := commander.Start(); err != nil {
		log.Fatal(err)
	}

}
