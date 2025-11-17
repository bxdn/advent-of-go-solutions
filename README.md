# Welcome to Advent of Go!

## What is Advent of Go?

Advent of Go is a multi-year framework for [Advent of Code](https://adventofcode.com/), a festive programming advent calendar!
It can house all your solutions in an easy and organized layout, and comes with built-in functionality that allows you to generate solutions stubs, pull input data, print results, submit answers, and pull available answers and test against them.

## Installation

1. Install [Go](https://go.dev/doc/install) and [Git](https://git-scm.com/install/) if you haven't already
1. If you would like to keep your solutions up to date on github, fork the repository first
1. Run `git clone github.com/bxdn/advent-of-go` Or `git clone <Your forked repo>`
1. Run `cd advent-of-go`
1. The first time you run, you will be asked to paste your session cookie from Advent of Code (you can find it in your browser once you've logged in).
1. That's it!

## Usage

### Printing Solutions
The usage of Advent of Go is simple: To run and print all solutions, run `go run .`

However, you wouldn't need a framework to accomplish that.

If you, for instance, wanted to print only year 2019 day 5 part 2, you could run `go run . -y 2019 -d 5 -p 2`

### Generating Solution Stubs
In order to stub out a day of Advent of Go, provide the `-g` flag with the `-y` and `-d` flags, to provide a year and day.

So, in order to generate 2020 day 5 stubs to work on, run the following: `go run . -g -y 2020 -d 5`

You will then find those solutions to implement in `solutions/2020/day5` as `pt1.go` and `pt2.go`

### Pulling Input Data Sets
To make input data available to your solutions when run, provide the `i` flag. Similarly to `-g`, you will need both the `-y` and `-d` flags, to provide a year and day to pull.

So, to both generate stubs for 2021 day 19 and pull in the input data, you could run `go run . -g -i -y 2021 -d 19`

### Submitting
To submit a solution after implementation, use the `-s` flag. You will need the `-y`, `-d`, and `-p` flags to specify a solution to submit. 

The solution will be run and automatically submitted to Advent of Code, and the message in the response page will be printed.

So, once you have created your solution for year 2023 day 10 part 1, you could submit that solution with `go run . -s -y 2023 -d 10 -p 1`

### Pulling Answers to Test Against
Important: You can only pull answers after successfully submitting in the Advent of Code website or through Advent of Go.

To pull answers for a solution after getting the star(s), provide the `a` flag. Similarly to the `-g` and `-i` flags, you will need both the `-y` and `-d` flags, to provide a year and day to pull.

So, to pull the answers for your latest solution for 2022 day 12, you could run `go run . -a -y 2022 -d 12` Then, to test your solution, you could run `go run . -t -y 2022 -d 12`, or `go run . -t` if you have all the answers pulled for all implemented solutions.

### Testing Solutions
If you want to test against the answers to see if your solutions are correct, you can run `go run . -t` or to just print the failures, `go run . -q`

So, if you wanted to test only 2023's solutions in quiet mode, you could run `go run . -q -y 2023`

The tests will of course only pass if you have pulled the answers for those solutions using the `-a` flag.
