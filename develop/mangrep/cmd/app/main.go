package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/golvl2/develop/mangrep/internal/app/managers"
	"github.com/vlasove/golvl2/develop/mangrep/internal/app/mangrep"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "mangrep"
)

var (
	help                     bool
	version                  bool
	fileIn                   string
	fileOut                  string
	isFixedTemplate          bool
	isNumLinePrinted         bool
	isInversion              bool
	isIgnoreCase             bool
	isCounter                bool
	amountLinesPrintedAfter  int
	amountLinesPrintedBefore int
	amountLinesPrintedSides  int
)

func init() {
	flag.IntVar(&amountLinesPrintedSides, "C", 0, "(A+B) Печатать ±N строк вокруг совпадения")
	flag.IntVar(&amountLinesPrintedBefore, "B", 0, "Печатать +N строк до совпадения")
	flag.IntVar(&amountLinesPrintedAfter, "A", 0, "Печатать +N строк после совпадения")
	flag.BoolVar(&isCounter, "c", false, "Количество строк")
	flag.BoolVar(&isIgnoreCase, "i", false, "Игнорировать регистр")
	flag.BoolVar(&isInversion, "v", false, "Вместо совпадения, исключать")
	flag.BoolVar(&isNumLinePrinted, "n", false, "Печатать номер строки")
	flag.BoolVar(&isFixedTemplate, "F", false, "Точное совпадение со строкой template, не паттерн")
	flag.StringVar(&fileIn, "in", "", "Путь до файла входных данных. Если не указан - данные берутся из StdIN")
	flag.StringVar(&fileOut, "out", "", "Путь до файла вывода: по умолчанию - StdOut.")
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")

}

func usage() {
	log.Printf(`ПРОСТАЯ УТИЛИТА ДЛЯ ПОИСКА В СТРОКАХ
ИСПОЛЬЗОВАНИЕ: ./mangrep [OPTIONS] <template>
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

	args := flag.Args()
	if len(args) != 1 {
		showUsageAndExit(1)
	}

	var inputManager, outputManager managers.Manager
	switch fileIn {
	case "":
		inputManager = managers.NewConsoleManager(os.Stdin, os.Stdout)
	default:
		inputManager = managers.NewFileManager(fileIn)
	}

	switch fileOut {
	case "":
		outputManager = managers.NewConsoleManager(os.Stdin, os.Stdout)
	default:
		outputManager = managers.NewFileManager(fileOut)
	}
	options := []mangrep.Option{
		mangrep.SetFixedTemplateOption(isFixedTemplate),
		mangrep.SetNumLinePrintedOption(isNumLinePrinted),
		mangrep.SetInversionOption(isInversion),
		mangrep.SetIgnoreCaseOption(isIgnoreCase),
		mangrep.SetCounterOption(isCounter),
		mangrep.SetAmountLinesPrintedAfterOption(amountLinesPrintedAfter),
		mangrep.SetAmountLinesPrintedBeforeOption(amountLinesPrintedBefore),
		mangrep.SetAmountLinesPrintedSidesOption(amountLinesPrintedSides),
	}
	grep := mangrep.New(inputManager, outputManager, args[0]).ApplyOptions(options...)
	if err := grep.Search(); err != nil {
		log.Fatal(err)
	}

	if err := grep.OutputResult(); err != nil {
		log.Fatal(err)
	}
}
