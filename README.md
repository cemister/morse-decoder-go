# Morse Decoder
This program is a command-line tool for decoding Morse code. Given a string of Morse code as input, it will decode the message and output the corresponding letters and numbers.

## Prerequisites
1. [Go](https://go.dev/doc/install) version 1.16 or higher (because of `bytes.Buffer` type, which was introduced in 1.16)
2. [Git](https://git-scm.com/download/) (optional, may have already installed it if on Linux, i.e., out of box)

## Installation
1. Clone this repository to your local machine.
```
git clone https://github.com/sncelta/morse-decoder-go.git
```
Alternatively, you can download the source code as a ZIP file from the GitHub repository.

2. Navigate to the project directory.

```
cd morse-decoder-go
```

3. Build the program using the  `go build` command.
```
go build
```
### Result
This will create an executable file called morse_decoder in the current directory.

## Usage
To use the program, run the morse-code-decoder executable with the Morse code you want to decode as a command-line argument. For example:

```
./morse_decoder --morse ".--. ..- -.-. -.- / -- --- .-. ... . / -.-. --- -.. ." --warn
```
Alternatively, you can run it via source code:
```
go run main.go --morse ".--. ..- -.-. -.- / -- --- .-. ... . / -.-. --- -.. ." --warn
```
### Result
This will decode the Morse code ".--. ..- -.-. -.- / -- --- .-. ... . / -.-. --- -.. ." and output the corresponding letters and numbers. The `--warn` flag is optional and will cause the program to warn you if there's invalid codes in your input.

## Flags
The following flags are available:

  - `--morse` - The Morse code to decode. Required.
  - `--warn` - Toggle warnings for invalid codes. If this flag is present, the program will output a warning message if there's any invalid code in the input. True by default.

## Contributing
If you find a bug or have a suggestion for this program, please feel free to open an issue or submit a pull request on GitHub.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/sncelta/morse-decoder-go/blob/main/LICENSE) file for details.
