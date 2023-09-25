# The Aesir Programming Language

Welcome to the Aesir Programming Language! Aesir is a new programming language written in Go. It's currently at it's beta and is under active development.

## About Aesir

Aesir is designed to be a simple, yet powerful programming language. It's written in Go, making it highly efficient and portable. The language is still in its early stages, but it's already capable of handling a variety of tasks.

## Features

- Integer and Boolean expressions
- Prefix and Infix operators
- Conditional (if-else) expressions
- Return statements
- Error handling
- Let statements
- Function literals

## Examples

Here are some examples of what you can do with Aesir:

### Integer Expressions

```let a = 5; a; // returns 5```

### Boolean Expressions

```let a = true; a; // returns true```

### Arithmetic Operations

```let a = 5 * 5; a; // returns 25```

### Variable Assignments

```let a = 5; let b = a; b; // returns 5```

### Complex Expressions

```let a = 5; let b = a; let c = a + b + 5; c; // returns 15```

### Function Literals

```let identity = fn(x) { x; }; identity(5); // returns 5```

### If-Else Expressions

```let result = if (10 > 1) { return true } else { return false }; result; // returns true```

### Return Statements

```let earlyReturn = fn() { return 5; 10; }; earlyReturn(); // returns 5```

### Error Handling

```let error = 5 + true; // returns "type mismatch: INTEGER + BOOLEAN"```

## More Examples

### Fibonacci Sequence

```let fibonacci = fn(x) { if (x == 0) { 0 } else { if (x == 1) { 1 } else { fibonacci(x - 1) + fibonacci(x - 2) } } }; fibonacci(10);```

### Factorial

```let factorial = fn(x) { if (x == 0) { 1 } else { x * factorial(x - 1) } }; factorial(5);```

## Getting Started

To get started with Aesir, you'll need to have Go installed on your machine. Once you have Go installed, you can clone this repository and start it with `go run ./`. You will then be able to enter Aesir code into the REPL.

## Contributing

Contributions to the Aesir Programming Language are welcome! Whether it's reporting bugs, improving documentation, or contributing code, I appreciate all the help we can get.

## License

The Aesir Programming Language is open-source software licensed under the GPL-3 License. See [LICENSE](https://github.com/nexentra/aesir/blob/main/LICENSE.md) for more information.

## Authors

* **Towhid Khan** - *I am a Software Developer.I enjoy using my skills to contribute to the exciting technological advances that happen every day.* - [Towhid Khan](https://github.com/KnockOutEZ) - *Main Developer Behind this project*