package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const badgeFlag = "<!---badges-->"
const versionFlag = "<!---dev-version-->"

func main() {
	log.Println("Generating badges...")
	reportCard := os.Getenv("INPUT_REPORT-CARD")
	versionInput := os.Getenv("INPUT_VERSION")
	coverageInput := os.Getenv("INPUT_COVERAGE")
	readmePath := os.Getenv("INPUT_README-PATH")

	b, err := ioutil.ReadFile("/github/workspace/" + readmePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	fmt.Println("Lines: \n", lines)
	coverageBadge := fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/green)"
	if coverageInput != "-1" {
		coverage, err := strconv.ParseFloat(coverageInput, 64)
		if err != nil {
			return
		}
		if coverage < 80 && coverage >= 70 {
			coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
		} else if coverage < 70 && coverage >= 60 {
			coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
		} else if coverage < 60 {
			coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/red)"
		}
	}

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
	versionBadge := fmt.Sprintf("![](https://badgen.net/badge/%s%s", versionInput, "/blue)")

	startReportCard := 11
	for i, line := range lines {
		startReportCard += 1
		if strings.Contains(line, versionFlag) && versionBadge != "" {
			lines[i] = fmt.Sprintf("%s %s *_Released on %s_\"", versionBadge, versionFlag, time.Now().Format("2006-01-02 3:4:5 PM MST"))
		}
		if strings.Contains(line, badgeFlag) {
			lines[i] = fmt.Sprintf("%s %s %s", reportCardBadge, coverageBadge, badgeFlag)
			startReportCard = -2
		}
		if reportCardResults != nil && startReportCard >= 0 && startReportCard < len(reportCardResults) {
			if len(lines) > i && strings.Contains(lines[i], strings.Split(reportCardResults[startReportCard], ":")[0]) {
				lines[i] = reportCardResults[startReportCard]
			} else {
				lines = append(lines[:i+1], lines[i:]...)
				lines[i] = reportCardResults[startReportCard]
			}
			startReportCard += 1
		}
	}

	f, err := os.OpenFile("README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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

	cmd, err := exec.Command("/bin/sh", "commit.sh").Output()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(cmd))

	log.Println("Success! README.md badge changes were committed to the repo.")
}
