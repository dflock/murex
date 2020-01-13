package docs

func init() {

	Definition["if"] = "# _murex_ Shell Docs\n\n## Command Reference: `if`\n\n> Conditional statement to execute different blocks of code depending on the result of the condition\n\n## Description\n\nConditional control flow\n\n`if` can be utilized both as a method as well as a standalone function. As a\nmethod, the conditional state is derived from the calling function (eg if the\nprevious function succeeds then the condition is `true`).\n\n## Usage\n\n### Function `if`:\n\n    if { code-block } then {\n        # true\n    } else {\n        # false\n    }\n    \n### Method `if`:\n\n    command -> if {\n        # true\n    } else {\n        # false\n    }\n    \n### Negative Function `if`:\n\n    !if { code-block } then {\n        # false\n    }\n    \n### Negative Method `if`:\n\n    command -> !if {\n        # false\n    }\n    \n### Please Note:\nthe `then` and `else` statements are optional. So the first usage could\nalso be written as:\n\n    if { code-block } {\n        # true\n    } {\n        # false\n    }\n    \nHowever the practice of omitting those statements isn't recommended beyond\nwriting short one liners in the interactive command prompt.\n\n## Examples\n\nCheck if a file exists:\n\n    if { g somefile.txt } then {\n        out \"File exists\"\n    }\n    \n...or does not exist (both ways are valid):\n\n    !if { g somefile.txt } then {\n        out \"File does not exist\"\n    }\n    \n    if { g somefile.txt } else {\n        out \"File does not exist\"\n    }\n\n## Detail\n\nThe conditional block can contain entire pipelines - even multiple lines of code\nlet alone a single pipeline - as well as solitary commands as demonstrated in\nthe examples above. However the conditional block does not output STDOUT nor\nSTDERR to the rest of the pipeline so you don't have to worry about redirecting\nthe output streams to `null`.\n\nIf you require output from the conditional blocks STDOUT then you will need to\nuse either a _murex_ named pipe to redirect the output, or test or debug flags\n(depending on your use case) if you only need to occasionally inspect the\nconditionals output.\n\n## Synonyms\n\n* `if`\n* `!if`\n\n\n## See Also\n\n* [commands/`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [commands/`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [commands/`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [commands/`debug`](../commands/debug.md):\n  Debugging information\n* [commands/`false`](../commands/false.md):\n  Returns a `false` value\n* [commands/`or`](../commands/or.md):\n  Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n* [commands/`true`](../commands/true.md):\n  Returns a `true` value\n* [commands/`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [commands/`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error\n* [commands/test](../commands/test.md):\n  "

}
