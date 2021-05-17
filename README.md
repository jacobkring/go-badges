# Go Badges

Go Badges is a simple docker action intended to be used in Github Actions for Golang projects. Badges are generated using various inputs and comment tags in your README file.
<br/><br/>
![](https://badgen.net/badge/coverage/89.4%25/green) <!---go-badges-coverage-->
![](https://badgen.net/badge/release/v1.10.18/blue) <!---go-badges-version-->
<br/>![](https://badgen.net/badge/license/MIT/blue) <br/>
#### Go Report Card
![](https://badgen.net/badge/Report%20Card/A+%20(100.0%25)/green) <!---go-badges-report-card-->
```
Grade: A+ (100.0%)
Files: 2
Issues: 0
gofmt: 100%
go_vet: 100%
gocyclo: 100%
golint: 100%
license: 100%
ineffassign: 100%
misspell: 100%
```

## Motivation

I contribute to a number of private Golang repos that make it inconvenient or impossible to use 3rd party tools for things like goreportcard.com and generating badges.

## Usage

In your README file you should leave comments on the lines that you want the badges to be generated.
Checkout the raw README file in this project for an example. <br/><br/>You do not need to include any badge links,
they will be generated after the first run.

#### Coverage
```
<!---go-badges-coverage-->
```
#### Version
```
<!---go-badges-version-->
```
#### Report Card
```
<!---go-badges-report-card-->
```

#### Workflow

In your workflow file you should include the inputs you want generated.
```
on: [push]

jobs:
  go_badges:
    runs-on: ubuntu-latest
    name: A job to generate badges
    steps:
      - uses: actions/checkout@v2
      - id: badges
        uses: jacobkring/go-badges@v1.10.4
        with:
          version: ${{ steps.version.outputs.tag }}
          report-card: ${{ steps.goreportcard.outputs.reportCard }}
          coverage: ${{ steps.coverage.outputs.coverage }}
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
