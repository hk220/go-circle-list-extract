# go-circle-list-extract

go-circle-list-extract extracts a circle list of Comitia in JSON or CSV format.

## Install from binary

Donwload from [release page](https://github.com/hk220/go-circle-list-extract/releases.)

## Install from source

```bash
go install github.com/hk220/go-circle-list-extract
```

## Usage

```bash
go-circle-list-extract -f [format:json, csv] -o [output filename] [eventname]
```

## Support Event

See config.yaml

- Comitia 134: event name is `comitia134`
- Comitia 136: event name is `comitia136`
