# beawake

A tiny macOS command-line tool that stops your Mac from falling asleep.<br>
It works kind of like Apple's own `caffeinate`, using a native IOKit power assertion under the hood.<br> 
Run it, leave it be, and your screen won't dim or lock due to inactivity. <br>
Hit `Ctrl+C` whenever to stop this and your Mac will go back to sleeping normally.

## Requirements

- macOS
- [Go](https://go.dev/dl/) 1.20 or newer
- Xcode Command Line Tools, since this compiles some C code under the hood
```bash
xcode-select --install
```

## Installation

```bash
GOBIN=/usr/local/bin go install github.com/guru-sankar-subramanian/beawake@main
```

Check that it installed correctly:

```bash
which beawake
```

If that prints a path like `/usr/local/bin/beawake`, you're good to go.

## Usage

```bash
beawake
```

Leave it running for as long as you want your Mac awake. Press `Ctrl+C` when you're done, and things go back to normal.

## Uninstall

```bash
sudo rm /usr/local/bin/beawake
```