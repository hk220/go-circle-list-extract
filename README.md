# go-circle-list-extract

go-circle-list-extract extracts a circle list of Comitia in JSON or CSV format.

## Install from binary

Donwload a binary according to your OS and config.yaml from [release page](https://github.com/hk220/go-circle-list-extract/releases.)

## Install from source

```bash
go install github.com/hk220/go-circle-list-extract
```

## Usage

```bash
Usage:
  go-circle-list-extract [event name] [flags]

Flags:
  -c, --config string   config file (default is config.yaml) (default "config.yaml")
  -f, --format string   output format (support: json, csv) (default "csv")
  -h, --help            help for go-circle-list-extract
  -o, --output string   output file name (default "circles.csv")
```

## Support Event

See config.yaml

- Comitia 134: event name is `comitia134`
- Comitia 136: event name is `comitia136`
- Comitia 137: event name is `comitia137`

To see the old events, look at [SupportEvent.md](SupportEvent.md)
