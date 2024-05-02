# Installation

1. Download the repo
```bash
git clone https://github.com/zSnails/minigo.git
```

2. Make sure to download all the dependencies
```bash
go mod tidy
```

3. Build the compiler
```bash
go build cmd/minigo.go
```

After doing the steps above you'll be left with a binary file on the root of this
project, what you do with that is up to you, for the
[editor](https://github.com/Mortalcr/Tronchaeditor.git) to recognize it the binary
must be on your system's path.

If you're on linux you should already know how to add a binary to your path, if you
don't then consider moving back to windows.

If you're on windows there's this tool called the internet, you should check it out.


# Available Types

`bool`, `int`, `float`, `string`, `rune`

# Built in constants

`true`, `false`

# Built in functions

`append`, `len`, `cap`
