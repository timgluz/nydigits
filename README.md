# Little solver of the NY/Times digits

[![Go](https://github.com/timgluz/nydigits/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/timgluz/nydigits/actions/workflows/go.yml)

Small proof-of-concept how [NY/Times Digits](https://www.nytimes.com/games/digits)
could be solved programmatically.

## TODO

- [x] fix use of duplicate values
- [x] Show operations
- [x] Tests
- [x] Taskfile
- [ ] Goreleaser
- [ ] Spin
- [ ] Look for optimal solution

## Usage

```bash
go run main.go --target 456 3 13 19 20 23 25

Solving NYDigits
Found solution:  456
Target:   456
Distance: 0
----------------------------------
Operations:
        1:   3 +  23 =  26
        2:  26 *  19 = 494
        3: 494 -  13 = 481
        4: 481 -  25 = 456
```

### Compiling from source

If you have [Taskfile](https://taskfile.dev/) installed, then you can compile a binary
and use the compiled binary to avoid rebuilding it all the time;

```bash
task build

./bin/nydigits --target 93 5 7 9 10 15 25
```
