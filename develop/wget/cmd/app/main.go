package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/golvl2/develop/wget/internal/app/managers"
	"github.com/vlasove/golvl2/develop/wget/internal/app/wget"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "wget"
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
	log.Printf(`ПРОСТАЯ УТИЛИТА ДЛЯ КОПИРОВАНИЯ HTML-КОДА ВЕБ СТРАНИЦ
ИСПОЛЬЗОВАНИЕ: ./wget <url>`)
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
	if len(args) != 1 {
		showUsageAndExit(1)
	}

	url := args[0]
	dm := managers.New(url)
	wget := wget.New(url, dm)

	res, err := wget.Parse()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	size, err := dm.WriteResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("wget: downloaded %d bytes", size)

}
