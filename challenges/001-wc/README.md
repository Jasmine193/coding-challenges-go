# 001 — wc

An implementation of the Unix `wc` command written in Go, as part of the [John Crickett Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc).

The goal of this challenge is to build a command-line utility that can count **bytes, lines, words, and characters** in a file.

## Challenge

The Unix `wc` command provides different options for counting the contents of a file.

This implementation supports:

| Option | Description      |
| ------ | ---------------- |
| `-c`   | Count bytes      |
| `-l`   | Count lines      |
| `-w`   | Count words      |
| `-m`   | Count characters |

The executable is named `ccwc` to distinguish it from the system's existing `wc` command.

## Requirements

The implementation should support:

* Reading a file from the command line
* Counting bytes using `-c`
* Counting lines using `-l`
* Counting words using `-w`
* Counting characters using `-m`
* Running without an option and displaying the default counts
* Handling invalid or missing file paths
* Supporting UTF-8 characters for character counting

## Project Structure

```text
001-wc/
├── README.md
├── ccwc.go
├── ccwc_test.go
└── cmd/
    └── ccwc/
        └── main.go
```

### `cmd/ccwc/main.go`

The CLI entry point.

It is responsible for:

* Reading command-line arguments
* Identifying the requested option
* Reading the input file
* Calling the appropriate counting function
* Printing the result

### `ccwc.go`

Contains the core counting logic:

* `CountBytes`
* `CountLines`
* `CountWords`
* `CountCharacters`

Keeping this logic separate from the CLI makes the functions easier to test independently.

### `ccwc_test.go`

Contains unit tests for the counting functions, including edge cases for word counting.

## Implementation

### Byte counting

The `-c` option counts the number of bytes in the file.

The implementation uses:

```go
len(data)
```

Since the file is read using `os.ReadFile`, the data is represented as a `[]byte`, making its length the number of bytes.

This is different from character counting when the file contains multibyte UTF-8 characters.

### Line counting

The `-l` option counts newline characters (`\n`) in the file.

For example:

```text
Hello
World
Go
```

contains three newline-terminated lines.

### Word counting

The `-w` option counts whitespace-delimited words.

The implementation uses an `inWord` state to determine when a new word begins.

Conceptually:

```text
outside a word
      │
      │ non-whitespace
      ▼
 inside a word
      │
      │ whitespace
      ▼
outside a word
```

A word is counted only when a non-whitespace character is encountered while currently outside a word.

This prevents multiple consecutive whitespace characters from being counted as multiple words.

For example:

```text
Hello     World
```

contains:

```text
2 words
```

rather than 6.

### Character counting

The `-m` option counts characters rather than bytes.

This distinction becomes important with UTF-8.

For example:

```text
😊
```

occupies:

```text
4 bytes
1 Unicode code point
```

The implementation uses:

```go
utf8.RuneCount(data)
```

to count Unicode code points.

Therefore, for multibyte UTF-8 text:

```text
byte count != character count
```

## Usage

### Build

From the `001-wc` directory:

```bash
go build -o ccwc ./cmd/ccwc
```

### Count bytes

```bash
./ccwc -c test.txt
```

### Count lines

```bash
./ccwc -l test.txt
```

### Count words

```bash
./ccwc -w test.txt
```

### Count characters

```bash
./ccwc -m test.txt
```

### Default behavior

The command can also be run without an option:

```bash
./ccwc test.txt
```

This outputs the line, word, and byte counts along with the filename.

### Installing locally

The command can be installed into Go's binary directory:

```bash
go install ./cmd/ccwc
```

Once the Go binary directory is included in the system `PATH`, the command can be run directly:

```bash
ccwc -c test.txt
```

## Comparing With Unix `wc`

The implementation can be compared with the system's `wc` command.

For example:

```bash
wc -c test.txt
ccwc -c test.txt
```

```bash
wc -l test.txt
ccwc -l test.txt
```

```bash
wc -w test.txt
ccwc -w test.txt
```

```bash
wc -m test.txt
ccwc -m test.txt
```

Comparing the output is particularly useful for the `-m` option because character counting can depend on the current locale and multibyte character support.

## Testing

Run all tests from the `001-wc` directory:

```bash
go test ./...
```

For verbose output:

```bash
go test -v ./...
```

The tests cover:

* Byte counting
* Line counting
* Character counting
* Multiple words
* Multiple spaces
* Newlines
* Tabs
* Leading whitespace
* Trailing whitespace
* Empty input
* Input containing only whitespace

## Key Learnings

This challenge helped me understand several Go concepts:

### `byte` vs `rune`

A `byte` represents one byte of data, while a `rune` represents a Unicode code point.

For ASCII text, characters generally occupy one byte. With UTF-8, a character can occupy multiple bytes.

For example:

```text
A    → 1 byte
é    → 2 bytes
😊   → 4 bytes
```

This makes the distinction between:

```go
len(data)
```

and:

```go
utf8.RuneCount(data)
```

important when implementing text-processing tools.

### Command-line arguments

The CLI uses:

```go
os.Args
```

to access arguments supplied from the terminal.

For example:

```bash
ccwc -w test.txt
```

provides:

```text
os.Args[0] → ccwc
os.Args[1] → -w
os.Args[2] → test.txt
```

### State-based parsing

The word counter uses an `inWord` state to determine when a new word begins.

This is a simple example of state-based parsing and is useful for understanding how text-processing utilities can be implemented without relying on complex libraries.

### Table-driven tests

The word-count tests use Go's table-driven testing pattern, allowing multiple input/output cases to be tested using the same test logic.

## Future Improvements

Potential improvements to the current implementation include:

* More complete handling of whitespace characters
* More robust command-line argument parsing
* Support for multiple input files
* Support for reading from standard input
* Better error handling and usage messages
* Matching the output formatting of Unix `wc` more closely
* Additional tests using different Unicode characters and locales

## Status

**Completed**

Implemented:

* [x] `-c` byte counting
* [x] `-l` line counting
* [x] `-w` word counting
* [x] `-m` character counting
* [x] Command-line interface
* [x] Unit tests
* [x] Comparison with Unix `wc`
