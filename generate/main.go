package main

import (
	"aoc25/aoc"
	"aoc25/idea"
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
)

var fileTemplate = "package day_%d\n\nfunc Part%d(input string) {\n}\n"
var inputsPath = "inputs"

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func generate(year, day int) {
	client := aoc.CreateClient()
	inputFileName := path.Join(inputsPath, fmt.Sprintf("input-%d.txt", day))
	packageName := fmt.Sprintf("day_%d", day)
	file, _ := os.ReadFile("main/main.go")
	mainSrc := string(file)

	var isMainSetup = func() bool {
		exp, _ := regexp.Compile(fmt.Sprintf("%d:\\s+{\\s*day_%d.Part1,\\s*day_%d.Part2\\s*},", day, day, day))
		return exp.MatchString(mainSrc)
	}

	var updateImports = func() {
		exp, _ := regexp.Compile("(?s)import \\(.*?\\)\n")
		importSrc := exp.FindString(mainSrc)
		lines := strings.Split(importSrc, "\n")
		imports := lines[1 : len(lines)-2]
		imports = append(imports, fmt.Sprintf("\t\"aoc25/day_%d\"", day))
		slices.Sort(imports)
		importSrcUpdated := fmt.Sprintf("import (\n%s\n)\n", strings.Join(imports, "\n"))
		mainSrc = strings.Replace(mainSrc, importSrc, importSrcUpdated, -1)
	}

	var updateMap = func() {
		exp, _ := regexp.Compile("(?s)var advents =.*?},\\s+}\\s")
		mapSrc := exp.FindString(mainSrc)
		lines := strings.Split(mapSrc, "\n")
		days := lines[1 : len(lines)-2]
		days = append(days, fmt.Sprintf("\t%d: {day_%d.Part1, day_%d.Part2},", day, day, day))
		mapSrcUpdated := fmt.Sprintf("var advents = map[int][]func(input string){\n%s\n}\n", strings.Join(days, "\n"))
		mainSrc = strings.Replace(mainSrc, mapSrc, mapSrcUpdated, -1)
	}

	if !exists(packageName) {
		_ = os.Mkdir(packageName, 0755)
		_ = os.WriteFile(path.Join(packageName, "part-1.go"), []byte(fmt.Sprintf(fileTemplate, day, 1)), 0755)
		_ = os.WriteFile(path.Join(packageName, "part-2.go"), []byte(fmt.Sprintf(fileTemplate, day, 2)), 0755)
	}

	if !exists(inputFileName) {
		input := client.LoadInput(year, day)
		_ = os.WriteFile(path.Join(inputsPath, fmt.Sprintf("input-%d.txt", day)), []byte(input), 0755)
	}

	if !isMainSetup() {
		updateImports()
		updateMap()
		_ = os.WriteFile("main/main.go", []byte(mainSrc), 0755)
	}

	configs := createRunConfigurations(day)

	for _, config := range configs {
		if idea.HasRunConfiguration(config.Name) {
			continue
		}

		println(fmt.Sprintf("add config for %d", day))
		idea.AddRunConfiguration(&config)
	}

}

func createBaseRunConfiguration() idea.RunConfiguration {
	return idea.RunConfiguration{
		Type:             "GoApplicationRunConfiguration",
		FactoryName:      "Go Application",
		Module:           "advent2025",
		WorkingDirectory: "$PROJECT_DIR$",
		Kind:             "PACKAGE",
		Package:          "aoc25/main",
		Directory:        "$PROJECT_DIR$",
		FilePath:         "$PROJECT_DIR$/main/main.go",
		Method:           2,
	}
}

func createRunConfigurations(day int) []idea.RunConfiguration {
	var configurations []idea.RunConfiguration
	for i := range 2 {
		part := i + 1
		config := createBaseRunConfiguration()
		config.Name = fmt.Sprintf("Day %d, part %d", day, part)
		config.Parameters = []string{strconv.Itoa(day), strconv.Itoa(part)}
		config.FolderName = fmt.Sprintf("Day %d", day)
		configurations = append(configurations, config)

		testConfig := createBaseRunConfiguration()
		testConfig.Name = fmt.Sprintf("Day %d, part %d (test)", day, part)
		testConfig.Parameters = []string{strconv.Itoa(day), strconv.Itoa(part), "test"}
		testConfig.FolderName = fmt.Sprintf("Day %d", day)
		configurations = append(configurations, testConfig)
	}
	return configurations
}

func main() {
	date := time.Now().Local()
	year := date.Year()

	for day := range lo.Min([]int{date.Day(), 25}) {
		generate(year, day+1)
	}
}
