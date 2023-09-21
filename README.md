# The Aesir Programming Language

Welcome to the Aesir Programming Language! Aesir is a new programming language written in Go. It's currently at version 0.0.5 and is under active development.

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

```aesir let a = 5; a; // returns 5```

### Boolean Expressions

```aesir let a = true; a; // returns true```

### Arithmetic Operations

```aesir let a = 5 * 5; a; // returns 25```

### Variable Assignments

```aesir let a = 5; let b = a; b; // returns 5```

### Complex Expressions

```aesir let a = 5; let b = a; let c = a + b + 5; c; // returns 15```

### Function Literals

```aesir let identity = fn(x) { x; }; identity(5); // returns 5```

### If-Else Expressions

```aesir let result = if (10 > 1) { "yes" } else { "no" }; result; // returns "yes"```

### Return Statements

```aesir let earlyReturn = fn() { return 5; 10; }; earlyReturn(); // returns 5```

### Error Handling

```aesir let error = 5 + true; // returns "type mismatch: INTEGER + BOOLEAN"```


## Getting Started

To get started with Aesir, you'll need to have Go installed on your machine. Once you have Go installed, you can clone this repository and start it with `go run ./`. You will then be able to enter Aesir code into the REPL.

## Contributing

Contributions to the Aesir Programming Language are welcome! Whether it's reporting bugs, improving documentation, or contributing code, I appreciate all the help we can get.

## License

The Aesir Programming Language is open-source software licensed under the GPL-3 License. See [LICENSE](https://github.com/nexentra/aesir/blob/main/LICENSE.md) for more information.

## Authors

* **Towhid Khan** - *I am a Software Developer.I enjoy using my skills to contribute to the exciting technological advances that happen every day.* - [Towhid Khan](https://github.com/KnockOutEZ) - *Main Developer Behind this project*