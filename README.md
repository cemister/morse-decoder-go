# Morse Decoder
This is a simple command-line program written in Go that translates Morse code input into plain text. It can decode letters and numbers, and it also supports spaces between words.

## Requirements
  - Go version 1.16 or higher (because of `bytes.Buffer` type, which was introduced in 1.16)
  - Git (optional)
  
## Installation
1. Install Go by following the instructions on the [official Go website](https://golang.org/doc/install).

2. Install [Git](https://git-scm.com/download/win) (On Linux, Git is most likely already installed, i.e., out of box, so you may skip this step.)

2. Clone the repository from GitHub using Git:
```
git clone https://github.com/sncelta/morse-decoder-go.git
```

Alternatively, you can download the source code as a ZIP file from the GitHub repository.

3. Navigate to the root directory of the project.
```
cd morse-decoder-go/src
```
4. Build the project using the following command:
```
go build
```
This will create an executable file called morse_decoder.

## Usage
To use the program, provide a Morse code string as a command line argument using the -morse flag:
```
./morse-decoder-go -morse ".... . .-.. .-.. --- / .-- --- .-. .-.. -.. / -- --- .-. ... . / -.-. --- -.. ."
```

Note that you should enclose the Morse code string in quotes to prevent the shell from interpreting it as separate arguments. Also note that the forward slash character (/) is used to indicate a space in the input Morse code.

If the input Morse code is valid, the program will output the corresponding plain text. If the input contains invalid codes, the program will output a question mark (?) to indicate the invalid code.

## Contributing
If you would like to contribute to this project, you can do so by forking the repository, making changes, and submitting a pull request.

## License
This project is licensed under the MIT License - see the [LICENSE](https://github.com/sncelta/morse-decoder-go/blob/main/LICENSE) file for details.
