# Advent of Code 2021

My solutions for the [Advent of Code 2021](https://adventofcode.com/2021) challenges written in Go.

## Generating the markdown for the problems

First install the [aoc-to-markdown](https://github.com/antonio-ramadas/aoc-to-markdown) Python package:

```console
pip install aoc-to-markdown
```

Login to adventofcode.com and get the `session` cookie content.

Call the `aoc-markdown` script with the cookie content and the year and day.

Windows:

```console
aoc-markdown.bat SESSION_ID -y 2021 -d 1 -s
```

Linux (not tested):

```console
aoc-markdown.sh SESSION_ID -y 2021 -d 1 -s
```
