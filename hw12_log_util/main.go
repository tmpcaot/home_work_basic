package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// appConfig содержит конфигурацию приложения.
type appConfig struct {
	File   string
	Level  string
	Output string
}

// status хранит информацию о логе.
type status struct {
	ip     string
	level  string
	method string
	code   string
	engine string
}

// ConfigFile устанавливает значение поля File конфигурации.
func (cfg *appConfig) ConfigFile(value string) {
	if value == "" {
		cfg.File = os.Getenv("LOG_ANALYZER_FILE")
	} else {
		cfg.File = value
	}
}

// ConfigLevel устанавливает значение поля Level конфигурации.
func (cfg *appConfig) ConfigLevel(value string) {
	if value == "" {
		cfg.Level = os.Getenv("LOG_ANALYZER_LEVEL")
	} else {
		cfg.Level = value
	}
}

// ConfigOutput устанавливает значение поля Output конфигурации.
func (cfg *appConfig) ConfigOutput(value string) {
	if value == "" {
		cfg.Output = os.Getenv("LOG_ANALYZER_OUTPUT")
	} else {
		cfg.Output = value
	}
}

// ReadFile читает файл и возвращает его содержимое как срез строк.
func ReadFile(logfile string) ([]string, error) {
	words := make([]string, 0)
	file, err := os.Open(logfile)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, nil
}

// Add добавляет данные в структуру status.
func (s *status) Add(ip, level, method, code, engine string) {
	s.ip = ip
	s.level = level
	s.method = method
	s.code = code
	s.engine = engine
}

// sort сортирует статус по уровню и возвращает карту с подсчетами.
func sort(sl []status, level string) map[string]map[string]int64 {
	m := map[string]map[string]int64{
		"ip":     {},
		"method": {},
		"code":   {},
		"engine": {},
	}
	for _, s := range sl {
		if s.level != strings.ToUpper(level) {
			continue
		}

		m["ip"][s.ip]++
		m["method"][s.method]++
		m["engine"][s.engine]++
		m["code"][s.code]++
	}
	return m
}

// writeFile записывает результаты в указанный файл.
func writeFile(input map[string]map[string]int64, pathfile string) error {
	file, err := os.Create(pathfile)
	if err != nil {
		return err
	}
	defer file.Close()
	for k, v := range input {
		for kk, vv := range v {
			_, writeErr := file.WriteString(fmt.Sprintf("[%s][%s][%d]\n", k, kk, vv))
			if writeErr != nil {
				return writeErr
			}
		}
	}
	return nil
}

func main() {
	var (
		logAnalyzerFile   string
		logAnalyzerLevel  string
		logAnalyzerOutput string
		showHelp          bool
	)

	// Использование встроенного пакета флагов
	flag.StringVar(&logAnalyzerFile, "file", "", "Входной файл")
	flag.StringVar(&logAnalyzerLevel, "level", "", "Уровень журнала: DEBUG, INFO, ERROR")
	flag.StringVar(&logAnalyzerOutput, "output", "", "Выходной файл")
	flag.BoolVar(&showHelp, "help", false, "Показать сообщение помощи")

	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	c := appConfig{}
	c.ConfigFile(logAnalyzerFile)
	c.ConfigLevel(logAnalyzerLevel)
	c.ConfigOutput(logAnalyzerOutput)

	if c.File == "" {
		panic("не указан входной файл")
	}

	if c.Level == "" {
		c.Level = "INFO"
	}

	// Предварительно выделяем память под срез st
	const initialCapacity = 1024
	st := make([]status, 0, initialCapacity)

	stringlog, err := ReadFile(c.File)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range stringlog {
		splitstring := strings.Fields(line)
		if len(splitstring) < 7 {
			continue // Пропускаем некорректные строки
		}
		ip := splitstring[0]
		level := splitstring[1]
		method := splitstring[4]
		code := splitstring[5]
		engine := splitstring[6]

		var struc status
		struc.Add(ip, level, method, code, engine)
		st = append(st, struc)
	}

	rr := sort(st, c.Level)
	if c.Output == "" {
		for k, v := range rr {
			for kk, vv := range v {
				fmt.Println(k, kk, vv)
			}
		}
	} else {
		wErr := writeFile(rr, c.Output)
		if wErr != nil {
			fmt.Printf("Ошибка записи в файл: %v\n", wErr)
		}
	}
}
