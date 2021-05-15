package main

import (
	"log"
	"os"
)

const badgeFlag = "<!---badges-->"
const prdVersionFlag = "<!---prd-version-->"
const stgVersionFlag = "<!---stg-version-->"
const devVersionFlag = "<!---dev-version-->"

func main() {
	log.Println("Generating badges...")
	log.Println(os.Args)
	/*
	if len(os.Args) < 3 {
		panic("expected coverage and report card")
	}
	coverageInput := os.Args[1]
	reportCard := os.Args[2]
	version := strings.Split(os.Args[3], ";")

	b, err := ioutil.ReadFile("README.md")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	coverage, err := strconv.ParseFloat(coverageInput, 64)
	if err != nil {
		return
	}

	coverageBadge := fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/green)"
	if coverage < 80 && coverage >= 70 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
	} else if coverage < 70 && coverage >= 60 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/yellow)"
	} else if coverage < 60 {
		coverageBadge = fmt.Sprintf("![](https://badgen.net/badge/coverage/%s", coverageInput) + "%25/red)"
	}

	reportCardResults := strings.Split(reportCard, "\n")
	reportCardGrade := strings.ReplaceAll(strings.ReplaceAll(strings.Split(reportCardResults[0], ": ")[1], "%", "%25"), " ", "%20")
	reportCardBadge := "![](https://badgen.net/badge/Report%20Card/" + reportCardGrade
	if strings.Contains(reportCard, "A") {
		reportCardBadge = reportCardBadge + "/green)"
	} else if strings.Contains(reportCard, "B") {
		reportCardBadge = reportCardBadge + "/yellow)"
	} else if strings.Contains(reportCard, "C") {
		reportCardBadge = reportCardBadge + "/orange)"
	} else {
		reportCardBadge = reportCardBadge + "/red)"
	}

	productionVersionBadge := ""
	stagingVersionBadge := ""
	developmentVersionBadge := ""
	switch version[0] {
	case "production":
		productionVersionBadge = fmt.Sprintf("![](https://badgen.net/badge/%s%s", strings.Join(version, "/"), "/blue)")
	case "staging":
		stagingVersionBadge = fmt.Sprintf("![](https://badgen.net/badge/%s%s", strings.Join(version, "/"), "/cyan)")
	case "development":
		developmentVersionBadge = fmt.Sprintf("![](https://badgen.net/badge/%s%s", strings.Join(version, "/"), "/grey)")
	}

	startReportCard := 11
	for i, line := range lines {
		startReportCard += 1
		if strings.Contains(line, prdVersionFlag) && productionVersionBadge != "" {
			lines[i] = fmt.Sprintf("%s %s *_Released on %s_\"", productionVersionBadge, prdVersionFlag, time.Now().Local().Format("2006-01-02 3:4:5 PM MST"))
		}
		if strings.Contains(line, stgVersionFlag) && stagingVersionBadge != "" {
			lines[i] = fmt.Sprintf("%s %s *_Released on %s_", stagingVersionBadge, stgVersionFlag, time.Now().Local().Format("2006-01-02 3:4:5 PM MST"))
		}
		if strings.Contains(line, devVersionFlag) && developmentVersionBadge != "" {
			lines[i] = fmt.Sprintf("%s %s *_Released on %s_\"", developmentVersionBadge, devVersionFlag, time.Now().Format("2006-01-02 3:4:5 PM MST"))
		}
		if strings.Contains(line, badgeFlag) {
			lines[i] = fmt.Sprintf("%s %s %s", reportCardBadge, coverageBadge, badgeFlag)
			startReportCard = -2
		}
		if startReportCard >= 0 && startReportCard < len(reportCardResults) {
			lines[i] = reportCardResults[startReportCard]
			startReportCard += 1
		}
	}

	f, err := os.OpenFile("README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = f.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	output := strings.Join(lines, "\n")
	_, err = fmt.Fprintf(f, "%s", output)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Success!")
	 */
}