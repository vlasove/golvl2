package main

import (
	"flag"
	"log"
	"os"

	"github.com/vlasove/materials/tasks_2/utils/mansort/internal/app/managers"
	"github.com/vlasove/materials/tasks_2/utils/mansort/internal/app/mansort"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "mansort"
)

var (
	fileIn            string
	fileOut           string
	help              bool
	version           bool
	keyColumnNum      int
	numericSortColNum int
	monthColNum       int
	reverseFlag       bool
	uniqueFlag        bool
	alreadySorted     bool
	ignoreTails       bool
)

func init() {
	flag.StringVar(&fileIn, "in", "", "Путь до файла входных данных. Если не указан - данные берутся из StdIN")
	flag.IntVar(&keyColumnNum, "k", -1, "Номер столбца по которому сортируем. По умолчанию - вся строка")
	flag.IntVar(&numericSortColNum, "n", -1, "Номер колонки с числовым значением для сортировки. По умолчанию - отсутствует")
	flag.IntVar(&monthColNum, "M", -1, "Номер колонки с месяцем. Валидны только кириллические.")
	flag.StringVar(&fileOut, "out", "", "Путь до файла вывода: по умолчанию - StdOut.")
	flag.BoolVar(&reverseFlag, "r", false, "Сортировка в обратном порядке")
	flag.BoolVar(&ignoreTails, "b", false, "Игнорировать хвостовые пробелы")
	flag.BoolVar(&uniqueFlag, "u", false, "Оставить только уникальные строки")
	flag.BoolVar(&alreadySorted, "c", false, "Проверяет, отсортированы ли данные и выходит")
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")
}

func usage() {
	log.Printf(`ПРОСТАЯ УТИЛИТА ДЛЯ СОРТИРОВКИ СТРОК
ИСПОЛЬЗОВАНИЕ: ./mansort [OPTIONS]
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

	options := []mansort.Option{
		mansort.KeyColumnNum(keyColumnNum),
		mansort.ReverseNeeded(reverseFlag),
		mansort.OnlyUnique(uniqueFlag),
		mansort.AlreadySorted(alreadySorted),
		mansort.IgnoreTails(ignoreTails),
		mansort.NumColSort(numericSortColNum),
		mansort.MonthColSort(monthColNum),
	}
	mansort := mansort.New(inputManager, outputManager).ApplyOptions(options...)
	if err := mansort.Sort(); err != nil {
		log.Fatal(err)
	}
	if err := mansort.OutputResult(); err != nil {
		log.Fatal(err)
	}
}
