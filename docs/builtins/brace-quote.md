# _murex_ Language Guide

## Command reference: brace-quote

> Write a string to the STDOUT without new line

### Description

Write parameters to STDOUT (does not include a new line)

### Usage

    (string to write) -> <stdout>

### Examples

    » (Hello, World!)
    Hello, World!

    » (Hello, World!)
    Hello, World!

### Detail

The `(` function performs exactly like the `(` token for quoting so you do not
need to escape other tokens (eg single / double quotes, `'`/`"`, nor curly
braces, `{}`). However the braces are nestable so you will need to escape those
characters if you don't want them nested.

#### ANSI Constants

`(` supports ANSI constants.

### Synonyms

* (

### See also

* [`>`](>.md): Writes STDIN to disk - overwriting contents if file already exists
* [`>>`](>>.md): Writes STDIN to disk - appending contents if file already exists
* `cast`
* [`err`](err.md): `echo` a string to the STDERR
* [`out`](out.md): `echo` a string to the STDOUT with a trailing new line character
* [`pt`](pt.md): Pipe telemetry. Writes data-types and bytes written
* `sprintf`
* [`tout`](tout.md): `echo` a string to the STDOUT and set it's data-type
* [`ttyfd`](ttyfd.md): Returns the TTY device of the parent.