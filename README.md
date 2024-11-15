# Wordify

A powerful text manipulation tool built in Go that handles text completion, editing, and auto-correction.

## 🌟 Features

- Convert hexadecimal and binary numbers to decimal
- Smart case transformation for text
- Intelligent punctuation formatting
- Article ("a"/"an") correction
- Context-aware spell checking
- Command-line interface for easy usage

## 🛠️ Installation

```bash
# Clone the repository
git clone https://github.com/Stella-Achar-Oiro/wordify.git

# Navigate to project directory
cd wordify

# Build the executable
go build .
```

## 📖 Usage

Run Wordify by providing input and output file paths:

```bash
./wordify input.txt output.txt
```

### Example:

Input (`sample.txt`):
```text
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```

Output (`result.txt`):
```text
Simply add 66 and 2 and you will see the result is 68.
```

## 💡 Key Features

- Number System Conversion
  - Hex to Decimal
  - Binary to Decimal
- Text Transformation
  - Case modification
  - Punctuation formatting
- Smart Corrections
  - Article usage (a/an)
  - Context-based spell checking

## 🔧 Technology Stack

- Go
- Command-line interface
- Text processing algorithms
- File I/O operations

## License

MIT License
