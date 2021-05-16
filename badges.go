package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const coverageFlag = "<!---go-badges-coverage-->"
const reportCardFlag = "<!---go-badges-report-card-->"
const versionFlag = "<!---go-badges-version-->"

func isGoBadges() bool {
	return os.Getenv("IS_GO_BADGES") == "true"
}

func maxedBadges(counts map[string]int, badge string) bool {
	if isGoBadges() {
		return counts[badge] == 1
	}
	return false
}

// CoverageBadge generates a badge with the given coverage percentage (without the % symbol)
func CoverageBadge(coverageInput string) (string, error) {
	coverageBadge := fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/green)"

	coverage, err := strconv.ParseFloat(coverageInput, 64)
	if err != nil {
		return coverageBadge, err
	}
	if coverage < 80 && coverage >= 70 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
	} else if coverage < 70 && coverage >= 60 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/orange)"
	} else if coverage < 60 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/red)"
	}
	return coverageBadge, err
}

func ReportCardBadge(reportCard string) (string, []string) {
	reportCardBadge := "![](https://badgen.net/badge/Report%20Card/"
	var reportCardResults []string
	if reportCard != "" {
		reportCardResults = strings.Split(reportCard, "\n")
		reportCardGrade := strings.ReplaceAll(strings.ReplaceAll(strings.Split(reportCardResults[0], ": ")[1], "%", "%25"), " ", "%20")
		reportCardBadge = reportCardBadge + reportCardGrade
		if strings.Contains(reportCard, "A") {
			reportCardBadge = reportCardBadge + "/green)"
		} else if strings.Contains(reportCard, "B") {
			reportCardBadge = reportCardBadge + "/yellow)"
		} else if strings.Contains(reportCard, "C") {
			reportCardBadge = reportCardBadge + "/orange)"
		} else {
			reportCardBadge = reportCardBadge + "/red)"
		}
	}
	return reportCardBadge, reportCardResults
}

func ModifyLines(lines []string, reportCardResults []string, versionBadge string, coverageBadge string, reportCardBadge string) []string {
	counts := map[string]int{
		"coverage":   0,
		"reportCard": 0,
		"version":    0,
	}

	i := 0
	var newLines []string
	for i < len(lines) {
		line := lines[i]
		if strings.Contains(line, versionFlag) && versionBadge != "" && !maxedBadges(counts, "version") {
			newLines = append(newLines, fmt.Sprintf("%s %s", versionBadge, versionFlag))
			counts["version"] += 1
		} else if strings.Contains(line, coverageFlag) && !maxedBadges(counts, "coverage") {
			newLines = append(newLines, fmt.Sprintf("%s %s", coverageBadge, coverageFlag))
			counts["coverage"] += 1
		} else if reportCardResults != nil && strings.Contains(line, reportCardFlag) {
			reportCardSlice := append([]string{fmt.Sprintf("%s %s", reportCardBadge, reportCardFlag), "```"}, reportCardResults...)
			reportCardSlice = append(reportCardSlice, "```")
			if len(lines) > i+1 && strings.Contains(lines[i+1], "```") && strings.Contains(lines[i+2], "Grade") {
				i += 12
			}
			newLines = append(newLines, reportCardSlice...)
			counts["reportCard"] += 1
		} else {
			newLines = append(newLines, line)
		}
		i += 1
	}
	return newLines
}

var readmeBasePath = "/github/workspace"

func main() {
	log.Println("Generating badges...")
	reportCard := os.Getenv("INPUT_REPORT-CARD")
	versionInput := os.Getenv("INPUT_VERSION")
	coverageInput := os.Getenv("INPUT_COVERAGE")
	readmePath := os.Getenv("INPUT_README-PATH")

	b, err := ioutil.ReadFile(readmeBasePath + readmePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	coverageBadge, err := CoverageBadge(coverageInput)
	if err != nil {
		log.Fatal(err)
	}

	reportCardBadge, reportCardResults := ReportCardBadge(reportCard)

	versionBadge := fmt.Sprintf("![](https://badgen.net/badge/release/%s%s", versionInput, "/blue)")

	lines = ModifyLines(lines, reportCardResults, versionBadge, coverageBadge, reportCardBadge)

	f, err := os.OpenFile(readmeBasePath+readmePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	output := strings.Join(lines, "\n")
	_, err = fmt.Fprintf(f, "%s", output)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}

	_, err = exec.Command("/bin/sh", "commit.sh").Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success! README.md badge changes were committed to the repo.")
}
