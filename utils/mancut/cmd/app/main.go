package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/materials/tasks_2/utils/mancut/internal/app/managers"
	"github.com/vlasove/materials/tasks_2/utils/mancut/internal/app/mancut"
)

// SOLID + конфиги хендлеров
const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "mancut"
)

var (
	fields      string
	delimeter   string
	isSeparated bool
	help        bool
	version     bool
)

func init() {
	flag.StringVar(&fields, "f", "", "Выбрать поля (колонки). Целые положительные числа, через запятую в кавычках")
	flag.StringVar(&delimeter, "d", "\t", "Использовать другой разделитель")
	flag.BoolVar(&isSeparated, "s", false, "Выводить только строки с разделителем")
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")
}

func usage() {
	log.Printf(`ПРОСТАЯ УТИЛИТА ДЛЯ ОБРЕЗКИ СТРОК
ИСПОЛЬЗОВАНИЕ: COMMAND | ./mancut [OPTIONS]
ДОСТУПНЫЕ ОПЦИИ:`)
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

	manager := managers.NewConsoleManager(os.Stdin, os.Stdout)

	options := []mancut.Option{
		mancut.SetFieldsOption(fields),
		mancut.SetDelimeterOption(delimeter),
		mancut.SetSeparatedOption(isSeparated),
	}

	mancut := mancut.New(manager).ApplyOptions(options...)
	if err := mancut.Cut(); err != nil {
		log.Fatal(err)
	}

	if err := mancut.OutputResult(); err != nil {
		log.Fatal(err)
	}
}
