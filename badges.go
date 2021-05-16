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

const coverageFlag = "<!---go-badges-coverage-->"
const reportCardFlag = "<!---go-badges-report-card-->"
const versionFlag = "<!---go-badges-version-->"

func isGoBadges() bool {
	fmt.Println(os.Getenv("IS_GO_BADGES"))
	return os.Getenv("IS_GO_BADGES") == "true"
}

func maxedBadges(counts map[string]int, badge string) bool {
	fmt.Println(counts[badge] == 1)
	if isGoBadges() {
		return counts[badge] == 1
	} else {
		return false
	}
}

func main() {
	counts := map[string]int{
		"coverage":   0,
		"reportCard": 0,
		"version":    0,
	}
	log.Println("Generating badges...")
	reportCard := os.Getenv("INPUT_REPORT-CARD")
	versionInput := os.Getenv("INPUT_VERSION")
	coverageInput := os.Getenv("INPUT_COVERAGE")
	readmePath := os.Getenv("INPUT_README-PATH")

	b, err := ioutil.ReadFile("/github/workspace" + readmePath)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

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
		fmt.Println(reportCard)
		fmt.Println(reportCard)
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

	versionBadge := fmt.Sprintf("![](https://badgen.net/badge/release/%s%s", versionInput, "/blue)")

	i := 0
	for i < len(lines) {
		line := lines[i]
		if strings.Contains(line, versionFlag) && versionBadge != "" && !maxedBadges(counts, "version") {
			lines[i] = fmt.Sprintf("%s %s *_Released on %s_\"", versionBadge, versionFlag, time.Now().Format("2006-01-02 3:4:5 PM MST"))
			counts["version"] += 1
		}
		if strings.Contains(line, coverageFlag) && !maxedBadges(counts, "coverage") {
			lines[i] = fmt.Sprintf("%s %s", coverageBadge, coverageFlag)
			counts["coverage"] += 1
		}
		if reportCardResults != nil && strings.Contains(line, reportCardFlag) {
			startReportCard := 0
			lines[i] = fmt.Sprintf("%s %s", reportCardBadge, reportCardFlag)
			i += 1
			line = lines[i]
			for startReportCard < len(reportCardResults) {
				if len(lines) > i && strings.Contains(lines[i], "```") {
					// then it seems like we already have a generated badges report card
					lines[i] = "```"
					i += 1
					for startReportCard < len(reportCardResults) {
						lines[i] = reportCardResults[startReportCard]
						startReportCard += 1
						i += 1
					}
					lines[i] = "```"
				} else {
					// seems like we haven't generated a report card yet
					lines = append(lines[:i+1], lines[i:]...)
					lines[i] = "```"
					i += 1
					for startReportCard < len(reportCardResults) {
						lines = append(lines[:i+1], lines[i:]...)
						lines[i] = reportCardResults[startReportCard]
						startReportCard += 1
						i += 1
					}
					lines = append(lines[:i+1], lines[i:]...)
					lines[i] = "```"
				}
				startReportCard += 1
			}
		}
		i += 1
	}

	f, err := os.OpenFile("/github/workspace"+readmePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
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
