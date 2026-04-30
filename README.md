# ASCII Art Generator (Go)

+ This project is a simple ASCII art generator written in Go. It converts input text into ASCII art using different banner styles.

## Features

Converts text into ASCII art

+ Supports multiple banner styles:
standard
shadow
thinkertoy

+ Handles multi-line input using \n
Validates printable ASCII characters (32–126)
Clean and modular code structure

### How It Works

+ The program reads a banner file (.txt) containing ASCII representations of characters. Each character is defined using 8 lines.

+ Core Components
+ 
+ LoadBanner

Loads a banner file (standard.txt, shadow.txt, or thinkertoy.txt)
Maps ASCII characters to their visual representation
GenerateArt
Processes user input
Handles multi-line text (\n)
Validates characters
Generates final ASCII output
RenderLine
Renders a single line of ASCII art
SplitLine
Splits input into multiple lines
ValidateAll
Ensures characters are within supported ASCII range

### Usage

Basic usage (default banner)
go run . "Hello"
Using a specific banner
go run . "Hello" standard
go run . "Hello" shadow
go run . "Hello" thinkertoy
Multi-line input
go run . "Hello\nWorld" shadow
📂 Project Structure
```
├── main.go
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── banner_loader.go
├── generator.go
├── renderer.go
├── utils.go
└── README.md
```
### Notes
Banner files must be correctly formatted (minimum 855 lines)
Unsupported characters will return an error
Empty input produces no output
If no banner is specified, standard.txt is used by default
🛠 Example

+ Input:

+ go run . "Hi" thinkertoy

+ Output:
```

_   _
| | |
| |_| |
|  _  |
| | | |
|_| |_|
```

(Output varies depending on the selected banner)

📖 Requirements
Go (version 1.18 or higher)
✍️ Author

Onoja217
