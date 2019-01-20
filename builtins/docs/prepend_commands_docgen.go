package docs

func init() {

	Definition["prepend"] = "# _murex_ Language Guide\n\n## Command Reference: `prepend` \n\n> Add data to the start of an array\n\n### Description\n\n`prepend` a data to the start of an array.\n\n### Usage\n\n    <stdin> -> prepend: value -> <stdout>\n\n### Examples\n\n    » a: [January..December] -> prepend: 'New Year'\n    New Year\n    January\n    February\n    March\n    April\n    May\n    June\n    July\n    August\n    September\n    October\n    November\n    December\n\n### Detail\n\nIt's worth noting that `prepend` and `append` are not data type aware. So \nany integers in data type aware structures will be converted into strings:\n\n    » tout: json [1,2,3] -> prepend: new \n    [\n        \"new\",\n        \"1\",\n        \"2\",\n        \"3\"\n    ]\n\n### See Also\n\n* [`a`](../commands/a.md):\n  A sophisticated yet simply way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`len` ](../commands/len.md):\n  Outputs the length of an array\n* [mtac](../commands/mtac.md):\n  \n* [square-bracket-open](../commands/square-bracket-open.md):\n  \n* [update](../commands/update.md):\n  "

}
