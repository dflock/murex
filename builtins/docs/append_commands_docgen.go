package docs

func init() {

	Definition["append"] = "# _murex_ Language Guide\n\n## Command Reference: `append`\n\n> Add data to the end of an array\n\n### Description\n\n`append` data to the end of an array.\n\n### Usage\n\n    <stdin> -> append: value -> <stdout>\n\n### Examples\n\n    » a: [Monday..Sunday] -> append: Funday\n    Monday\n    Tuesday\n    Wednesday\n    Thursday\n    Friday\n    Saturday\n    Sunday\n    Funday\n\n### Detail\n\nIt's worth noting that `prepend` and `append` are not data type aware. So \nany integers in data type aware structures will be converted into strings:\n\n    » tout: json [1,2,3] -> append: new \n    [\n        \"1\",\n        \"2\",\n        \"3\",\n        \"new\"\n    ]\n\n### See Also\n\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a`](../commands/a.md):\n  A sophisticated yet simply way to build an array or list\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [mtac](../commands/mtac.md):\n  \n* [range](../commands/range.md):\n  \n* [square-bracket-open](../commands/square-bracket-open.md):\n  \n* [update](../commands/update.md):\n  "

}
