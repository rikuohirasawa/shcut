# shcut

A simple command-line tool for managing shell shortcuts and aliases.

## Installation

### Using Go
```bash
go install github.com/rikuohirasawa/shcut@latest
```

### Manual Installation
```bash
git clone https://github.com/rikuohirasawa/shcut.git
cd shcut
go build -o shcut
sudo cp shcut /usr/local/bin/
```

## Usage

### Add a shortcut
```bash
# Interactive mode
shcut add

# Direct mode
shcut add my-alias "echo 'Hello World'"
```

### List all shortcuts
```bash
shcut ls
```

### Run a shortcut
```bash
shcut run my-alias
```

### Browse shortcuts interactively
```bash
shcut browse
```

### Remove a shortcut
```bash
shcut rm my-alias
```

## Configuration

Shortcuts are stored in `~/.shcut/config.json` by default.

