# csv

Small and badly implemented CSV tool

## Installation

```
go install github.com/jamillosantos/csv/cli/csv@latest
```

## Usage

```
csv [flags] [file.csv]
```

### Examples

The following command will read `file.csv` and output the columns 1, 2 and 3.
```
csv --columns 1,2,3 file.csv
```

You can also use the stdin to read the CSV file. The following command will do the same as above but using the STDIN.
```
cat file.csv | csv --columns 1,2,3
```

### Flags

```
      --columns string     Output format (example: 1,2,3)
      --separator string   CSV separator (default ",")
      --skip-headers       Skip headers
  -h, --help               help for csv
```