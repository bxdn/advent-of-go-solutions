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

### Generating Solution Stubs
In order to stub out a day of Advent of Go, provide the `-g` flag with the `-y` and `-d` flags, to provide a year and day.

So, in order to generate 2020 day 5 stubs to work on, run the following: `go run . -g -y 2020 -d 5`

You will then find those solutions to implement in `solutions/2020/day5` as `pt1.go` and `pt2.go`

As a shorthand, you can also use `-n` (now) if you want to use the current day and year. 

So, if you want to generate today's solution, it's as easy as running `go run . -g -n`

### Printing Solutions
To run and print all solutions, run `go run .`

However, you wouldn't need a framework to accomplish that.

If you, for instance, wanted to print only year 2019 day 5 part 2, you could run `go run . -y 2019 -d 5 -p 2`

Inputs will also be pulled in for each solution as needed.

### Submitting
To submit a solution after implementation, use the `-s` flag. You will need the `-y`, `-d`, and `-p` flags to specify a solution to submit. 

The solution will be run and automatically submitted to Advent of Code, and the message in the response page will be printed.

So, once you have created your solution for year 2023 day 10 part 1, you could submit that solution with `go run . -s -y 2023 -d 10 -p 1`

Again, you could also run `go run . -s -n -p 1` to (s)ubmit (p)art 1's solution for the puzzle corresponding to (n)ow

### Testing Solutions
If you want to test against the answers to see if your solutions are correct, you can run `go run . -t` or to just print the failures, `go run . -q`

So, if you wanted to test only 2023's solutions in quiet mode, you could run `go run . -q -y 2023`

The tests will of course only pass if you have previously submitted the answers for those solutions.

The answers will be pulled in automatically as needed for testing.
