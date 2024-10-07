# dtoconv

is a simple command-line tool written in Go that converts domain names into their respective organization names.

## Features

- Convert domain names (e.g., `tesla.com`) to organization names (e.g., `Tesla`).
- Supports input from standard input or files.
- Verbose output option for detailed logging.

## Installation

`go install -v github.com/Vulnpire/dtoconv@latest`

## Usage

You can use dtoconv in two ways:

## From Standard Input

```
$ echo "tesla.com" | dtoconv

Tesla
```

## From a File

You can also provide a file containing a list of domain names:

```
$ tac urls.txt | dtoconv

Tesla
Asana Inc.
Fisglobal
Meraki LLC
```
