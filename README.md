# Wordify

### How to Use Wordify: A Simple Text Editing Tool in Go

#### Introduction
Wordify is a simple text completion, editing, and auto-correction tool written in Go. It allows you to make various modifications to a text file, such as converting hexadecimal and binary numbers to decimal, changing the case of words, and formatting punctuation.

#### Prerequisites
Before using Wordify, ensure you have the following:
- Basic knowledge of Go programming language
- Go environment set up on your machine

#### Installation
To use Wordify, you need to clone the repository and build the executable file. Follow these steps:
1. Clone the Wordify repository:
   ```
   git clone https://github.com/Stella-Achar-Oiro/wordify.git
   ```
2. Navigate to the Wordify directory:
   ```
   cd wordify
   ```
3. Build the executable file:
   ```
   go build .
   ```

#### Usage
To use Wordify, follow these steps:
1. Run the Wordify executable with the input file and output file as arguments:
   ```
   ./wordify input.txt output.txt
   ```
   Replace `input.txt` with the path to your input file and `output.txt` with the path where you want to save the modified text.

2. Wordify will process the input file according to the specified modifications and save the result in the output file.

#### Possible Outcomes
- Conversion of hexadecimal and binary numbers to decimal.
- Change in the case of words to uppercase, lowercase, or capitalized.
- Formatting of punctuation marks.
- Correction of articles "a" and "an" based on the following word.
- Auto-correction of misspelled words based on context.

#### Example
For example, given the input file `sample.txt`:
```
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
```
Running Wordify:
```
./wordify sample.txt result.txt
```
Will produce the output file `result.txt`:
```
Simply add 66 and 2 and you will see the result is 68.
```

#### Conclusion
Wordify is a useful tool for text manipulation in Go. It provides a simple yet powerful way to make modifications to text files, making it ideal for apprentice software developers learning Go.