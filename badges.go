package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const coverageFlag = "<!---go-badges-coverage-->"
const reportCardFlag = "<!---go-badges-report-card-->"
const versionFlag = "<!---go-badges-version-->"

func isGoBadges() bool {
	return os.Getenv("INPUT_IS_GO_BADGES") == "true"
}

func maxedBadges(counts map[string]int, badge string) bool {
	if isGoBadges() {
		return counts[badge] == 1
	}
	return false
}

func coverageBadge(coverageInput string) (string, error) {
	cBadge := fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/green)"

	coverage, err := strconv.ParseFloat(coverageInput, 64)
	if err != nil {
		return cBadge, err
	}
	if coverage < 80 && coverage >= 70 {
		cBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
	} else if coverage < 70 && coverage >= 60 {
		cBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/orange)"
	} else if coverage < 60 {
		cBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/red)"
	}
	return cBadge, err
}

func reportCardBadge(reportCard string) (string, []string) {
	rcBadge := "![](https://badgen.net/badge/Report%20Card/"
	var reportCardResults []string
	if reportCard != "" {
		reportCardResults = strings.Split(reportCard, "\n")
		reportCardGrade := strings.ReplaceAll(strings.ReplaceAll(strings.Split(reportCardResults[0], ": ")[1], "%", "%25"), " ", "%20")
		rcBadge = rcBadge + reportCardGrade
		if strings.Contains(reportCard, "A") {
			rcBadge = rcBadge + "/green)"
		} else if strings.Contains(reportCard, "B") {
			rcBadge = rcBadge + "/yellow)"
		} else if strings.Contains(reportCard, "C") {
			rcBadge = rcBadge + "/orange)"
		} else {
			rcBadge = rcBadge + "/red)"
		}
	}
	return rcBadge, reportCardResults
}

func modifyLines(lines []string, reportCardResults []string, versionBadge string, cBadge string, rcBadge string) []string {
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
			newLines = append(newLines, fmt.Sprintf("%s %s", cBadge, coverageFlag))
			counts["coverage"] += 1
		} else if reportCardResults != nil && strings.Contains(line, reportCardFlag) && !maxedBadges(counts, "reportCard") {
			reportCardSlice := append([]string{fmt.Sprintf("%s %s", rcBadge, reportCardFlag), "```"}, reportCardResults...)
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

	log.Println(reportCard)

	b, err := ioutil.ReadFile(readmeBasePath + readmePath)
	if err != nil {
		log.Fatal("read existing readme", err)
	}

	lines := strings.Split(string(b), "\n")

	cBadge, err := coverageBadge(coverageInput)
	if err != nil {
		log.Fatal("coverage badge", err)
	}

	rcBadge, reportCardResults := reportCardBadge(reportCard)

	versionBadge := fmt.Sprintf("![](https://badgen.net/badge/release/%s%s", versionInput, "/blue)")

	lines = modifyLines(lines, reportCardResults, versionBadge, cBadge, rcBadge)

	f, err := os.OpenFile(readmeBasePath+readmePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("open file for writing", err)
	}

	err = f.Truncate(0)
	if err != nil {
		log.Fatal("truncate file", err)
	}

	output := strings.Join(lines, "\n")
	_, err = fmt.Fprintf(f, "%s", output)
	if err != nil {
		log.Fatal("join lines", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal("close file", err)
	}

	/*
		cmd, err := exec.Command("/bin/sh", "commit.sh").Output()
		if err != nil {
			log.Println("result commit", string(cmd))
			log.Fatal("run commit", err)
		}

		log.Println("result commit", string(cmd))

	*/

	log.Println("Success!")
}
