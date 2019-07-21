package docs

func init() {

	Definition["swivel-datatype"] = "# _murex_ Language Guide\n\n## Command Reference: `swivel-datatype`\n\n> Converts tabulated data into a map of values for serialised data-types such as JSON and YAML\n\n### Description\n\n`swivel-datatype` rotates a table by 90 degrees then exports the output as a\nseries of maps to be marshalled by a serialised datatype such as JSON or YAML.\n\n### Usage\n\n    <stdin> -> swivel-datatype: data-type -> <stdout>\n\n### Examples\n\nLets take the first 5 entries from `ps`:\n\n    » ps: aux -> head: -n5 -> format: csv\n    \"USER\",\"PID\",\"%CPU\",\"%MEM\",\"VSZ\",\"RSS\",\"TTY\",\"STAT\",\"START\",\"TIME\",\"COMMAND\"\n    \"root\",\"1\",\"0.0\",\"0.1\",\"233996\",\"8736\",\"?\",\"Ss\",\"Feb19\",\"0:02\",\"/sbin/init\"\n    \"root\",\"2\",\"0.0\",\"0.0\",\"0\",\"0\",\"?\",\"S\",\"Feb19\",\"0:00\",\"[kthreadd]\"\n    \"root\",\"4\",\"0.0\",\"0.0\",\"0\",\"0\",\"?\",\"I<\",\"Feb19\",\"0:00\",\"[kworker/0:0H]\"\n    \"root\",\"6\",\"0.0\",\"0.0\",\"0\",\"0\",\"?\",\"I<\",\"Feb19\",\"0:00\",\"[mm_percpu_wq]\"\n    \nThat data swivelled would look like the following:\n\n    » ps: aux -> head: -n5 -> format: csv -> swivel-datatype: yaml\n    '%CPU':\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    '%MEM':\n    - \"0.1\"\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    COMMAND:\n    - /sbin/init\n    - '[kthreadd]'\n    - '[kworker/0:0H]'\n    - '[mm_percpu_wq]'\n    PID:\n    - \"1\"\n    - \"2\"\n    - \"4\"\n    - \"6\"\n    RSS:\n    - \"8736\"\n    - \"0\"\n    - \"0\"\n    - \"0\"\n    START:\n    - Feb19\n    - Feb19\n    - Feb19\n    - Feb19\n    STAT:\n    - Ss\n    - S\n    - I<\n    - I<\n    TIME:\n    - \"0:02\"\n    - \"0:00\"\n    - \"0:00\"\n    - \"0:00\"\n    TTY:\n    - '?'\n    - '?'\n    - '?'\n    - '?'\n    USER:\n    - root\n    - root\n    - root\n    - root\n    VSZ:\n    - \"233996\"\n    - \"0\"\n    - \"0\"\n    - \"0\"\n    \nPlease note that for input data-types whose table doesn't define titles (such as\nthe `generic` datatype), the map keys are defaulted to column numbers:\n\n    » ps: aux -> head: -n5 -> swivel-datatype: yaml\n    \"0\":\n    - USER\n    - root\n    - root\n    - root\n    - root\n    \"1\":\n    - PID\n    - \"1\"\n    - \"2\"\n    - \"4\"\n    - \"6\"\n    \"2\":\n    - '%CPU'\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    \"3\":\n    - '%MEM'\n    - \"0.1\"\n    - \"0.0\"\n    - \"0.0\"\n    - \"0.0\"\n    ...\n\n### Detail\n\nYou can check what output data-types are available via the `runtime` command:\n\n    runtime --marshallers\n    \nMarshallers are enabled at compile time from the `builtins/data-types` directory.\n\n### See Also\n\n* [`alter`](../commands/alter.md):\n  Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [`swivel-table`](../commands/swivel-table.md):\n  Rotates a table by 90 degrees\n* [format](../commands/format.md):\n  \n* [square-bracket-open](../commands/square-bracket-open.md):\n  "

}
