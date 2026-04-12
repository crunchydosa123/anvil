# 🔨 Anvil

**Anvil** is a minimal programming language built in Go, featuring a lexer, Pratt parser, AST, and interpreter.

It supports variables, arithmetic expressions with operator precedence, and execution via a CLI.

---

## 🚀 Features

* 🧩 Custom **lexer** (tokenizer)
* 🌳 **AST-based parser** using Pratt parsing (handles precedence)
* ⚡ **Interpreter** (evaluates expressions)
* 📦 Variable support (`let`)
* ➕ Arithmetic operations: `+ - * /`
* 🔁 Parentheses support
* 🖨️ `print()` for output
* 💻 CLI tool (`anvil run file.anv`)
* ⚠️ Basic syntax error handling

---

## 🏗️ Architecture

```
Source Code (.anv)
        ↓
     Lexer
        ↓
     Parser (AST)
        ↓
   Evaluator
        ↓
  Environment (variables)
        ↓
      Output
```

---

## 🧠 Example

### Input (`test.anv`)

```js
let x = 2 + 3 * 4;
let y = (x + 10) / 2;
print(y);
```

### Output

```
11
```

---

## 🛠️ Installation & Usage

### 1. Clone the repo

```
git clone https://github.com/<your-username>/anvil.git
cd anvil
```

### 2. Run a file

```
go run cmd/anvil/main.go run test.anv
```

### 3. Build binary (optional)

```
go build -o anvil cmd/anvil/main.go
./anvil run test.anv
```

---

## 📁 Project Structure

```
anvil/
├── lexer/        # Tokenizer
├── parser/       # Pratt parser
├── ast/          # AST definitions
├── evaluator/    # Interpreter
├── cmd/anvil/    # CLI entrypoint
└── token/        # Token definitions
```

---

## ⚠️ Error Handling

Example:

```js
let = 10;
```

Output:

```
Parser errors:
 - expected next token to be IDENT, got = instead
```

---

## 🔥 Future Improvements

* REPL (`anvil repl`)
* Better runtime error handling
* Functions & control flow
* Transpiler (Anvil → Go / LLVM IR)

---

## 🎯 Motivation

Built to understand compiler design fundamentals:

* Lexing
* Parsing
* AST construction
* Expression evaluation

---

## 🧠 Key Learnings

* Implemented Pratt parsing for operator precedence
* Designed a full interpreter pipeline
* Built a CLI-based language runtime in Go

---

## 📜 License

MIT License
