# Go Badges

Go Badges is a simple docker action intended to be used in Github Actions for Golang projects. Badges are generated using various inputs and comment tags in your README file.
<br/>
![](https://badgen.net/badge/coverage/90.4%25/green) <!---go-badges-coverage-->
![](https://badgen.net/badge/release/v1.10.2/blue) <!---go-badges-version-->
<br/>![](https://badgen.net/badge/license/MIT/blue) <br/>
#### Go Report Card
![](https://badgen.net/badge/Report%20Card/A+%20(94.1%25)/green) <!---go-badges-report-card-->
```
Grade: A+ (94.1%)
```
## Motivation

I contribute to a number of private Golang repos that make it inconvenient or impossible to use 3rd party tools for things like goreportcard.com and generating badges.

## Usage

In your README file you should leave comments on the lines that you want the badges to be generated.
Checkout the raw README file in this project for an example. <br/><br/>You do not need to include any badge links,
they will be generated after the first run.

#### Coverage
```
![](https://badgen.net/badge/coverage/90.4%25/green) <!---go-badges-coverage-->
```
#### Report Card
```
![](https://badgen.net/badge/Report%20Card/A+%20(94.1%25)/green) <!---go-badges-report-card-->
```
Grade: A+ (94.1%)
```
jobs:
  go_badges:
    runs-on: ubuntu-latest
    name: A job to generate badges
    steps:
      - uses: actions/checkout@v2
      - id: badges
        uses: jacobkring/go-badges@v1
        with:
          readmePath: 'README.md' # default is README.md
          coverage: '90.5'
          reportCard: 'true'
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
