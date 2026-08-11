# 002 — JSON Parser

A JSON parser implemented from scratch in Go, as part of the [John Crickett Coding Challenges](https://codingchallenges.fyi/challenges/challenge-json-parser).

The goal of this challenge is to build a JSON parser without relying on Go's built-in `encoding/json` package for parsing. The challenge provides an opportunity to understand **lexical analysis, tokenisation, parsing, grammars, and recursive data structures**.

## Challenge

JSON (JavaScript Object Notation) is a lightweight data-interchange format commonly used for exchanging structured data.

The parser is being built incrementally, starting with the simplest valid JSON object and gradually adding support for strings, primitive values, objects, arrays, and more complex JSON structures.

The original challenge is divided into five steps.

## Requirements

The parser will progressively support:

* Parsing an empty JSON object: `{}`
* Parsing objects containing string keys and values
* Parsing string, numeric, boolean, and null values
* Parsing nested objects
* Parsing arrays
* Detecting invalid JSON
* Returning useful parsing errors
* Testing the parser against valid and invalid JSON inputs

## Project Structure

```text
002-json-parser/
├── README.md
├── parser.go
├── parser_test.go
└── cmd/
    └── json-parser/
        └── main.go
```

### `cmd/json-parser/main.go`

The CLI entry point.

It is responsible for:

* Reading the input
* Invoking the parser
* Reporting whether the JSON is valid
* Returning an appropriate exit code

### `parser.go`

Contains the core parsing implementation.

The parser is responsible for:

* Reading the input
* Identifying JSON tokens
* Validating the JSON structure
* Building or validating the parsed representation
* Reporting parsing errors

### `parser_test.go`

Contains unit tests for the parser.

Tests cover both valid and invalid JSON and will be expanded as new parser functionality is implemented.

## Implementation Approach

The parser is being developed incrementally rather than trying to support the complete JSON specification from the beginning.

The general parsing pipeline is:

```text
JSON input
    │
    ▼
Lexical analysis
    │
    ▼
Tokens
    │
    ▼
Parser
    │
    ▼
JSON structure
    │
    ▼
Valid / Invalid
```

### Lexical Analysis

Lexical analysis involves breaking the input into meaningful pieces, or **tokens**.

For JSON, these can include:

```text
{
}
[
]
:
,
string
number
true
false
null
```

The lexer focuses on answering:

> "What are the meaningful pieces of this input?"

### Parsing

The parser takes those tokens and determines whether they form a valid JSON structure.

For example:

```json
{}
```

is valid because an object can contain zero members.

Whereas:

```json
{
```

is invalid because the object is never closed.

The parser therefore focuses on answering:

> "Do these tokens form a valid JSON structure?"

## Challenge Progress

### Step 1 — Empty Object

Support the simplest possible JSON object:

```json
{}
```

The parser should:

* Accept `{}` as valid JSON
* Reject invalid JSON
* Return exit code `0` for valid input
* Return exit code `1` for invalid input

Status:

* [ ] Implemented
* [ ] Tests added

### Step 2 — String Values

Support objects containing string keys and string values:

```json
{
  "key": "value"
}
```

The parser needs to understand:

* Object delimiters `{` and `}`
* String keys
* String values
* The `:` separator

Status:

* [ ] Implemented
* [ ] Tests added

### Step 3 — Primitive Values

Extend the parser to support:

```json
{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}
```

This introduces the following JSON value types:

* String
* Number
* Boolean
* Null

Status:

* [ ] Implemented
* [ ] Tests added

### Step 4 — Objects and Arrays

Support objects and arrays as JSON values:

```json
{
  "key": "value",
  "key-n": 101,
  "key-o": {},
  "key-l": []
}
```

This introduces recursive structures.

For example:

```json
{
  "user": {
    "name": "Jasmine"
  },
  "skills": [
    "Go",
    "Backend"
  ]
}
```

At this stage, the parser needs to handle nested objects and arrays.

Status:

* [ ] Implemented
* [ ] Tests added

### Step 5 — Additional Testing

Add tests beyond the provided challenge cases.

The goal is to ensure the parser:

* Accepts valid JSON
* Rejects malformed JSON
* Produces useful errors
* Handles nested structures
* Handles edge cases

The original challenge also suggests testing against the JSON checker test suite once the parser is complete.

Status:

* [ ] Implemented
* [ ] Tests added

## Usage

Build the CLI:

```bash
go build -o json-parser ./cmd/json-parser
```

Run it against a JSON file:

```bash
./json-parser test.json
```

For valid JSON:

```text
Valid JSON
```

For invalid JSON:

```text
Invalid JSON
```

The CLI should also return an appropriate process exit code:

```text
0 → valid JSON
1 → invalid JSON
```

This makes the parser usable in shell scripts and automated tests.

## Testing

Run all tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

As the parser progresses through each step, new test cases are added to verify the corresponding functionality.

## Key Learnings

This challenge focuses on several important programming concepts.

### Lexical Analysis

Understanding how a raw input string can be broken into meaningful tokens.

### Parsing

Understanding how tokens can be validated against a grammar to determine whether an input is structurally valid.

### Recursive Parsing

JSON naturally contains recursive structures:

```json
{
  "user": {
    "address": {
      "city": "Bengaluru"
    }
  }
}
```

Objects can contain objects, arrays can contain objects, and arrays can contain other arrays.

This makes JSON a good practical example of recursive parsing.

### State and Lookahead

The parser needs to keep track of where it is in the input and determine what token or structure is expected next.

For example, after:

```json
{
  "name"
```

the parser should expect:

```text
:
```

followed by a JSON value.

### Error Handling

A parser isn't only responsible for accepting valid input.

It also needs to reject invalid input and ideally provide enough information to understand what went wrong.

## Design Decisions

The parser is intentionally implemented from scratch rather than using:

```go
encoding/json
```

The purpose of the challenge is to understand how parsing works internally rather than simply deserialize JSON into Go structs.

The implementation will therefore favor clarity and explicit parsing logic over using a high-level JSON library.

## Future Improvements

Potential improvements after completing the challenge include:

* More comprehensive Unicode string handling
* JSON escape sequences
* Better number parsing
* More descriptive parse errors
* Line and column information in errors
* Additional malformed JSON test cases
* Running the parser against a larger JSON compliance test suite
* Separating lexer and parser into independent components
* Building an AST or generic JSON value representation

## Status

**In Progress**

The parser is being implemented incrementally following the stages of the Coding Challenges specification.

* [ ] Step 1 — Empty object
* [ ] Step 2 — String values
* [ ] Step 3 — Primitive values
* [ ] Step 4 — Objects and arrays
* [ ] Step 5 — Additional tests
