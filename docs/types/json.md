# _murex_ Language Guide

## Data-Type Reference: `json` (JSON)

> JavaScript Object Notation (JSON) (primitive)

### Description

JSON is a primitive data-type within _murex_.



### Examples

Example JSON document taken from [Wikipedia](https://en.wikipedia.org/wiki/JSON)

    {
      "firstName": "John",
      "lastName": "Smith",
      "isAlive": true,
      "age": 27,
      "address": {
        "streetAddress": "21 2nd Street",
        "city": "New York",
        "state": "NY",
        "postalCode": "10021-3100"
      },
      "phoneNumbers": [
        {
          "type": "home",
          "number": "212 555-1234"
        },
        {
          "type": "office",
          "number": "646 555-4567"
        },
        {
          "type": "mobile",
          "number": "123 456-7890"
        }
      ],
      "children": [],
      "spouse": null
    }

### Default Associations

* MIME: `application/json`
* Extension: `json`


### Supported Hooks

* `Marshaller()`
    Writes minified JSON when no TTY detected and human readable JSON when stdout is a TTY
* `ReadArray()`
    Works with JSON arrays. Maps are converted into arrays
* `ReadIndex()`
    Works against all properties in JSON
* `ReadMap()`
    Works with JSON maps
* `ReadNotIndex()`
    Works against all properties in JSON
* `Unmashaller()`
    Supported
* `WriteArray()`
    Works with JSON arrays

### See Also

* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`cast`](../commands/cast.md):
  Alters the data type of the previous function without altering it's output
* [`hcl` (HCL)](../types/hcl.md):
  HashiCorp Configuration Language (HCL)
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`runtime`](../commands/runtime.md):
  Returns runtime information on the internal state of _murex_
* [`toml` (TOML)](../types/toml.md):
  Tom's Obvious, Minimal Language (TOML)
* [`yaml` (YAML)](../types/yaml.md):
  YAML Ain't Markup Language (YAML)
* [element](../commands/element.md):
  
* [format](../commands/format.md):
  
* [open](../commands/open.md):
  