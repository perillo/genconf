# genconf utility

[![GoDoc](https://godoc.org/github.com/perillo/genconf?status.svg)](http://godoc.org/github.com/perillo/genconf)

genconf is a simple utility used to generate multiple configuration files from
a template (in Go text/template syntax) and a data file (in JSON format).

## Usage

    genconf [--data FILE] [FILE]

The template definition is read from the file passed as argument or stdin if
the argument list is empty or it is equal to "-".

The generated configuration file is written to stdout.

The template data is loaded from the file specified in the --data flag, and
unmarshaled in a `map[string]interface{}` value.

## Example

The following example is a systemd service file:

```systemd
; This file is a template (using Go text/template).
; Copy the generated file to ~/.config/systemd/user/ and enable it with
; systemctl --user enable <service>
;
; NOTE:
;   systemd version installed on Debian Jessie does not support quoting in
;   ExecStart: the quote characters will be passed to the application.

[Unit]
Description=Example web application

[Install]
WantedBy=default.target

[Service]
Type=simple
ExecStart={{ .ExecPath }} -addr={{ .Addr }} {{ .DSN }}
Restart=on-failure
WorkingDirectory={{ .WorkingDirectory }}
Environment={{ .Environment }}
```

And the following is a sample data file for local deployment:

```json
{
    "ExecPath": "/home/user/.local/bin/example",
    "WorkingDirectory": "/home/user/code/go/src/example",
    "Environment": "'GOPATH=/home/user/code/go' 'GOTRACEBACK=2'",
    "Addr": ":7070",
    "DSN": "user:password@/database?strict=true&sql_notes=false"
}
```
