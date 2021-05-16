package main

import (
	"fmt"
	"github.com/jacobkring/go-assert"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestCoverageBadge(t *testing.T) {
	_, err := coverageBadge("googly")
	assert.NotNil(t, err)
	assert.Equal(t, "strconv.ParseFloat: parsing \"googly\": invalid syntax", err.Error())
	cBadge, err := coverageBadge("90")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/90%25/green)", cBadge)
	cBadge, err = coverageBadge("75")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/75%25/yellow)", cBadge)
	cBadge, err = coverageBadge("65")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/65%25/orange)", cBadge)
	cBadge, err = coverageBadge("50")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/50%25/red)", cBadge)
}

func TestReportCardBadge(t *testing.T) {
	reportCard := "Grade: A (85.3%)\nFiles: 2\nIssues: 1\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 50%\ngolint: 50%\nlicense: 100%\nineffassign: 50%\nmisspell: 100%"
	rcBadge, reportCardResults := reportCardBadge(reportCard)

	assert.Equal(t, "![](https://badgen.net/badge/Report%20Card/A%20(85.3%25)/green)", rcBadge)
	assert.Equal(t, []string{"Grade: A (85.3%)", "Files: 2", "Issues: 1", "gofmt: 100%", "go_vet: 100%", "gocyclo: 50%", "golint: 50%", "license: 100%", "ineffassign: 50%", "misspell: 100%"}, reportCardResults)

	reportCard = "Grade: B (75.3%)\nFiles: 2\nIssues: 1\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 50%\ngolint: 50%\nlicense: 0%\nineffassign: 50%\nmisspell: 100%"
	rcBadge, reportCardResults = reportCardBadge(reportCard)

	assert.Equal(t, "![](https://badgen.net/badge/Report%20Card/B%20(75.3%25)/yellow)", rcBadge)
	assert.Equal(t, []string{"Grade: B (75.3%)", "Files: 2", "Issues: 1", "gofmt: 100%", "go_vet: 100%", "gocyclo: 50%", "golint: 50%", "license: 0%", "ineffassign: 50%", "misspell: 100%"}, reportCardResults)

	reportCard = "Grade: C (75.3%)\nFiles: 2\nIssues: 1\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 50%\ngolint: 50%\nlicense: 0%\nineffassign: 50%\nmisspell: 100%"
	rcBadge, reportCardResults = reportCardBadge(reportCard)

	assert.Equal(t, "![](https://badgen.net/badge/Report%20Card/C%20(75.3%25)/orange)", rcBadge)
	assert.Equal(t, []string{"Grade: C (75.3%)", "Files: 2", "Issues: 1", "gofmt: 100%", "go_vet: 100%", "gocyclo: 50%", "golint: 50%", "license: 0%", "ineffassign: 50%", "misspell: 100%"}, reportCardResults)

	reportCard = "Grade: D (15.3%)\nFiles: 2\nIssues: 1\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 50%\ngolint: 50%\nlicense: 0%\nineffassign: 50%\nmisspell: 100%"
	rcBadge, reportCardResults = reportCardBadge(reportCard)

	assert.Equal(t, "![](https://badgen.net/badge/Report%20Card/D%20(15.3%25)/red)", rcBadge)
	assert.Equal(t, []string{"Grade: D (15.3%)", "Files: 2", "Issues: 1", "gofmt: 100%", "go_vet: 100%", "gocyclo: 50%", "golint: 50%", "license: 0%", "ineffassign: 50%", "misspell: 100%"}, reportCardResults)
}

func TestModifyLines(t *testing.T) {
	// Report Card
	linesInput := []string{
		"<!---go-badges-report-card-->",
	}
	rcBadge := "![](https://badgen.net/badge/Report%20Card/A%20(85.3%25)/green)"
	reportCardResults := []string{"Grade: A (85.3%)", "Files: 2", "Issues: 1", "gofmt: 100%", "go_vet: 100%", "gocyclo: 50%", "golint: 50%", "license: 100%", "ineffassign: 50%", "misspell: 100%"}
	output := modifyLines(linesInput, reportCardResults, "", "", rcBadge)
	assert.Equal(t, 13, len(output))

	linesInput = append(output, "<!---go-badges-coverage-->")

	rcBadge, reportCardResults = reportCardBadge("Grade: A (100.0%)\nFiles: 2\nIssues: 0\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 100%\ngolint: 50%\nlicense: 100%\nineffassign: 100%\nmisspell: 100%")
	output = modifyLines(linesInput, reportCardResults, "", "![](https://badgen.net/badge/coverage/90%25/green)", rcBadge)
	assert.Equal(t, 14, len(output))

	assert.Equal(t, "![](https://badgen.net/badge/Report%20Card/A%20(100.0%25)/green) <!---go-badges-report-card-->", output[0])
	assert.Equal(t, "Grade: A (100.0%)", output[2])
	assert.Equal(t, "Issues: 0", output[4])
	assert.Equal(t, "gocyclo: 100%", output[7])
	assert.Equal(t, "ineffassign: 100%", output[10])

	linesInput[0] = "![](https://badgen.net/badge/Report%20Card/A%20(100.0%25)/green) <!---go-badges-report-card-->"
	linesInput[2] = "Grade: A (100.0%)"
	linesInput[4] = "Issues: 0"
	linesInput[7] = "gocyclo: 100%"
	linesInput[10] = "ineffassign: 100%"

	assert.Equal(t, "![](https://badgen.net/badge/coverage/90%25/green) <!---go-badges-coverage-->", output[13])
	linesInput[13] = "![](https://badgen.net/badge/coverage/90%25/green) <!---go-badges-coverage-->"
	assert.Equal(t, linesInput, output)

	linesInput = append(output, "<!---go-badges-version-->")

	versionBadge := "![](https://badgen.net/badge/release/v1.0.0/blue)"

	output = modifyLines(linesInput, reportCardResults, versionBadge, "![](https://badgen.net/badge/coverage/90%25/green)", rcBadge)
	assert.Equal(t, 15, len(output))

	assert.Equal(t, fmt.Sprintf("![](https://badgen.net/badge/release/v1.0.0/blue) <!---go-badges-version-->"), output[14])
	linesInput[14] = output[14]

	assert.Equal(t, linesInput, output)
}

func TestMainProcess(t *testing.T) {
	reportCard := "Grade: A (85.3%)\nFiles: 2\nIssues: 1\ngofmt: 100%\ngo_vet: 100%\ngocyclo: 50%\ngolint: 50%\nlicense: 100%\nineffassign: 50%\nmisspell: 100%"
	err := os.Setenv("INPUT_REPORT-CARD", reportCard)
	assert.Nil(t, err)
	err = os.Setenv("INPUT_VERSION", "v1.0.0")
	assert.Nil(t, err)

	err = os.Setenv("INPUT_COVERAGE", "90.5")
	assert.Nil(t, err)

	err = os.Setenv("INPUT_README-PATH", "/README.md")
	assert.Nil(t, err)

	readmeBasePath = "github/workspace"

	// copy clean readme
	data, err := ioutil.ReadFile("github/workspace/CLEAN_README.md")
	assert.Nil(t, err)
	// Write data to dst
	err = ioutil.WriteFile("github/workspace/README.md", data, 0644)
	assert.Nil(t, err)

	main()

	output, err := ioutil.ReadFile("github/workspace/README.md")
	if err != nil {
		log.Fatal(err)
	}
	result, err := ioutil.ReadFile("github/workspace/README_RESULT.md")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, string(result), string(output))

	err = os.Remove("github/workspace/README.md")
	assert.Nil(t, err)
}
