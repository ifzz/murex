package docs

func init() {

	Definition["a"] = "# _murex_ Language Guide\n\n## Command Reference: `a` (make array)\n\n> A sophisticated yet simple way to build an array or list\n\n### Description\n\n_murex_ has a pretty sophisticated builtin for generating arrays. Think\nlike bash's `{1..9}` syntax:\n\n    a: [1..9]\n\n### Usage\n\n    a: [start..end] -> <stdout>\n    a: [start..end.base] -> <stdout>\n    a: [start..end,start..end] -> <stdout>\n    a: [start..end][start..end] -> <stdout>\n    \nAll usages also work with `ja` and `ta` as well:\n\n    ja: [start..end] -> <stdout>\n    ta: data-type [start..end] -> <stdout>\n\n### Examples\n\n    » a: [1..3]\n    1\n    2\n    3\n    \n    » a: [3..1]\n    3\n    2\n    1\n    \n    » a: [01..03]\n    01\n    02\n    03\n\n### Detail\n\n#### Alternative Number Bases\n\nYou can also specify an alternative number base by using an `x` or `.`\nin the end range:\n\n    a: [00..ffx16]\n    a: [00..ff.16]\n    \nAll number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.\nPlease note that the start and end range are written in the target base\nwhile the base identifier is written in decimal: `[hex..hex.dec]`\n\nAlso note that the additional zeros denotes padding (ie the results will\nstart at `00`, `01`, etc rather than `0`, `1`...)\n\n#### Character arrays\n\nYou can select a range of letters (a to z):\n\n    » a: [a..z]\n    » a: [z..a]\n    » a: [A..Z]\n    » a: [Z..A]\n    \n...or any characters within that range.\n\n#### Special ranges\n\nUnlike bash, _murex_ also supports some special ranges:\n\n  \n    » a: [mon..sun]\n    » a: [monday..sunday]\n    » a: [jan..dec]\n    » a: [january..december]\n    » a: [spring..winter]\n    \nIt is also case aware. If the ranges are uppercase then the return will\nbe uppercase. If the ranges are title case (capital first letter) then\nthe return will be in title case:\n\n    » a: [Monday..Sunday]\n    Monday\n    Tuesday\n    Wednesday\n    Thursday\n    Friday\n    Saturday\n    Sunday\n    \nWhere the special ranges differ from a regular range is they cannot\ncannot down. eg `a: [3..1]` would output\n\n    » a: [3..1]\n    3\n    2\n    1\n    \nhowever a negative range in special ranges will cycle through to the end\nof the range and then loop back from the start:\n\n    » a: [Thursday..Wednesday]\n    Thursday\n    Friday\n    Saturday\n    Sunday\n    Monday\n    Tuesday\n    Wednesday\n    \nThis decision was made because generally with ranges of this type, you\nwould more often prefer to cycle through values rather than iterate\nbackwards through the list.\n\nIf you did want to reverse then just pipe the output into another tool:\n\n    » a: [Monday..Friday] -> mtac\n    Friday\n    Thursday\n    Wednesday\n    Tuesday\n    Monday\n    \nThere are other UNIX tools which aren't data type aware but would work in\nthis specific scenario:\n* `tac` (Linux),\n* `tail -r` (BSD / OS X)\n* `perl -e \"print reverse <>\"` (Multi-platform but requires Perl installed)\n\n#### Advanced Array Syntax\n\nThe syntax for `a` is a comma separated list of parameters with expansions\nstored in square brackets. You can have an expansion embedded inside a\nparameter or as it's own parameter. Expansions can also have multiple\nparameters.\n\n    » a: 01,02,03,05,06,07\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: 0[1..3],0[5..7]\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: 0[1..3,5..7]\n    01\n    02\n    03\n    05\n    06\n    07\n    \n    » a: b[o,i]b\n    bob\n    bib\n    \nYou can also have multiple expansion blocks in a single parameter:\n\n    » a: a[1..3]b[5..7]\n    a1b5\n    a1b6\n    a1b7\n    a2b5\n    a2b6\n    a2b7\n    a3b5\n    a3b6\n    a3b7\n    \n`a` will cycle through each iteration of the last expansion, moving itself\nbackwards through the string; behaving like an normal counter.\n\n#### Creating JSON arrays with `ja`\n\nAs you can see from the previous examples, `a` returns the array as a\nlist of strings. This is so you can stream excessively long arrays, for\nexample every IPv4 address: `a: [0..254].[0..254].[0..254].[0..254]`\n(this kind of array expansion would hang bash).\n\nHowever if you needed a JSON string then you can use all the same syntax\nas `a` but forgo the streaming capability:\n\n    » ja: [Monday..Sunday]\n    [\n        \"Monday\",\n        \"Tuesday\",\n        \"Wednesday\",\n        \"Thursday\",\n        \"Friday\",\n        \"Saturday\",\n        \"Sunday\"\n    ]\n    \nThis is particularly useful if you are adding formating that might break\nunder `a`'s formatting (which uses the `str` data type).\n\n### See Also\n\n* [`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [`ta`](../commands/ta.md):\n  A sophisticated yet simple way to build an array of a user defined data-type\n* [mtac](../commands/mtac.md):\n  "

}
