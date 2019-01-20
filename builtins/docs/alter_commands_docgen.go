package docs

func init() {

	Definition["alter"] = "# _murex_ Language Guide\n\n## Command Reference: `alter`\n\n> Change a value within a structured data-type and pass that change along the pipeline without altering the original source input\n\n### Description\n\n`alter` a value within a structured data-type.\n\nThe path separater is defined by the first character in the path. For example\n`/path/to/key`, `,path,to,key`, `|path|to|key` and `#path#to#key` are all valid\nhowever you should remember to quote or escape any special characters (tokens)\nused by the shell (such as pipe, `|`, and hash, `#`).\n\n### Usage\n\n    <stdin> -> alter: /path value -> <stdout>\n\n### Examples\n\n    » config: -> [ shell ] -> [ prompt ] -> alter: /Value moo\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Value\": \"moo\"\n    }\n    \n`alter` also accepts JSON as a parameter for adding structured data:\n\n    config: -> [ shell ] -> [ prompt ] -> alter: /Example { \"Foo\": \"Bar\" }\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Example\": {\n            \"Foo\": \"Bar\"\n        },\n        \"Value\": \"{ out 'murex » ' }\"\n    }\n    \nHowever it is also data type aware so if they key you're updating holds a string\n(for example) then the JSON data a will be stored as a string:\n\n    » config: -> [ shell ] -> [ prompt ] -> alter: /Value { \"Foo\": \"Bar\" }\n    {\n        \"Data-Type\": \"block\",\n        \"Default\": \"{ out 'murex » ' }\",\n        \"Description\": \"Interactive shell prompt.\",\n        \"Value\": \"{ \\\"Foo\\\": \\\"Bar\\\" }\"\n    }\n    \nNumbers will also follow the same transparent convertion treatment:\n\n    » tout: json { \"one\": 1, \"two\": 2 } -> alter: /two \"3\"\n    {\n        \"one\": 1,\n        \"two\": 3\n    }\n    \n> Please note: `alter` is not changing the value held inside `config` but\n> instead took the STDOUT from `config`, altered a value and then passed that\n> new complete structure through it's STDOUT.\n>\n> If you require modifying a structure inside _murex_ config (such as http\n> headers) then you can use `config alter`. Read the config docs for reference.\n\n### Detail\n\n#### Path\n\nThe path parameter can take any character as node separators. The separator is\nassigned via the first character in the path. For example\n\n    config -> alter: .shell.prompt.Value moo\n    config -> alter: >shell>prompt>Value moo\n    \nJust make sure you quote or escape any characters used as shell tokens. eg\n\n    config -> alter: '#shell#prompt#Value' moo\n    config -> alter: ' shell prompt Value' moo\n    \n#### Supported data-types\n\nYou can check what data-types are available via the `runtime` command:\n\n    runtime --marshallers\n    \nMarshallers are enabled at compile time from the `builtins/data-types` directory.\n\n### See Also\n\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [config](../commands/config.md):\n  \n* [format](../commands/format.md):\n  \n* [runtime](../commands/runtime.md):\n  \n* [square-bracket-open](../commands/square-bracket-open.md):\n  "

}
