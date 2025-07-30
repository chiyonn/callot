# Callot

![callot overview](docs/callot.png)

Callot is a command line application for calculating trading lot sizes based on
predefined risk settings. It stores its configuration in
`~/.config/callot/config.json` and provides a simple interactive interface as
well as a set of subcommands for managing that configuration.

## Installation

```
go install github.com/chiyonn/callot@latest
```

## Usage

Running `callot` without arguments launches the interactive mode. The program
will prompt you for a currency pair, the loss-cut width and the take-profit
ratio and will then display the maximum allowable trading volume.

All configuration can be adjusted via subcommands:

```
callot config show            # display current settings
callot config add-pair USDJPY # register a currency pair
callot config set-margin 40   # set margin amount (40 means 400000 JPY)
callot config set-risk 1.6    # set risk tolerance percentage
callot config set-ratio 2     # set default take-profit ratio
```

## Development

A prebuilt binary is included for convenience, but you can build everything from
source using `go build`.
