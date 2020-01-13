package docs

func init() {

	Definition["or"] = "# _murex_ Shell Docs\n\n## Command Reference: `or`\n\n> Returns `true` or `false` depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.\n\n## Description\n\nReturns a boolean results (`true` or `false`) depending on whether any of the\ncode-blocks included as parameters are successful or not.\n\n## Usage\n\n    or: { code-block } { code-block } -> <stdout>\n    \n    !or: { code-block } { code-block } -> <stdout>\n    \n`or` supports as many or as few code-blocks as you wish.\n\n## Examples\n\n    if { or { = 1+1==2 } { = 2+2==5 } { = 3+3==6 } } then {\n        out: At least one of those equations are correct\n    }\n\n## Detail\n\n`or` does not set the exit number on failure so it is safe to use inside a `try`\nor `trypipe` block.\n\nIf `or` is prefixed by a bang (`!or`) then it returns `true` when one or more\ncode-blocks are unsuccessful (ie the opposite of `or`).\n\n### Code-Block Testing\n\n* `or` only executes code-blocks up until one of the code-blocks is successful\n  then it exits the function and returns `true`.\n\n* `!or` only executes code-blocks while the code-blocks are successful. Once one\n  is unsuccessful `!or` exits and returns `true` (ie it `not`s every code-block).\n\n## Synonyms\n\n* `or`\n* `!or`\n\n\n## See Also\n\n* [commands/`!` (not)](../commands/not.md):\n  Reads the STDIN and exit number from previous process and not's it's condition\n* [commands/`and`](../commands/and.md):\n  Returns `true` or `false` depending on whether multiple conditions are met\n* [commands/`catch`](../commands/catch.md):\n  Handles the exception code raised by `try` or `trypipe` \n* [commands/`false`](../commands/false.md):\n  Returns a `false` value\n* [commands/`if`](../commands/if.md):\n  Conditional statement to execute different blocks of code depending on the result of the condition\n* [commands/`true`](../commands/true.md):\n  Returns a `true` value\n* [commands/`try`](../commands/try.md):\n  Handles errors inside a block of code\n* [commands/`trypipe`](../commands/trypipe.md):\n  Checks state of each function in a pipeline and exits block on error"

}
