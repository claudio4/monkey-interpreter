# Monkey Interpreter

This repository contains my implementation of the Monkey programming language interpreter, developed while following the book _Writing An Interpreter In Go_ by Thorsten Ball.

## Overview

Monkey is a simple yet powerful interpreted programming language. This interpreter includes support for:

- Primitive data types: integers, booleans, strings
- Arithmetic and boolean expressions
- Variable bindings and environments
- First-class functions with closures
- Built-in functions (e.g., `len`, `first`, `last`)
- Conditionals and loops
- Metaprogramming and macros.
- A basic REPL for interactive use

## Getting Started

### Prerequisites

To build and run the interpreter, you need:

- [Go](https://golang.org/dl/) (1.18 or later recommended)

### Installation

Clone the repository:

```sh
git clone https://github.com/yourusername/monkey-interpreter.git
cd monkey-interpreter
```

Build the interpreter:

```sh
go build -o monkey
```

Run the REPL:

```sh
./monkey
```

## Usage

In the REPL, you can enter Monkey code interactively:

```monkey
>> let add = fn(a, b) { a + b; };
>> add(2, 3);
5
```

You can also run Monkey scripts from a file:

```sh
./monkey script.monkey
```

## Project Structure

```
monkey/
├── ast/         # Abstract Syntax Tree definitions
├── evaluator/   # Expression and statement evaluation
├── lexer/       # Lexer implementation
├── object/      # Object system for runtime values
├── parser/      # Parsing logic
├── repl/        # Read-Eval-Print Loop (REPL)
├── token/       # Token definitions
└── main.go      # Entry point
```

## Acknowledgments

This project is based on _Writing An Interpreter In Go_ by Thorsten Ball. Highly recommended for anyone interested in language implementation!

## License

This project is licensed under the MIT License.

---

Happy coding! 🚀
