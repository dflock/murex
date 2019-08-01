# _murex_ Language Guide

## Command Reference: `mtac`

> Reverse the order of an array

### Description

`mtac` takes input from STDIN and reverses the order of it.

It's name is derived from a program called `tac` - a tool that functions
like `cat` but returns the contents in the reverse order. The difference
with the `mtac` builtin is that it is data-type aware. So it doesn't just
function as a replacement for `tac` but it also works on JSON arrays,
s-expressions, and any other data-type supporting arrays compiled into
_murex_. 

### Usage

    <stdin> -> mtac -> <stdout>

### Examples

    » ja: [Monday..Friday] -> mtac
    [
        "Friday",
        "Thursday",
        "Wednesday",
        "Tuesday",
        "Monday"
    ]
    
    # Normal output (without mtac)
    » ja: [Monday..Friday]
    [
        "Monday",
        "Tuesday",
        "Wednesday",
        "Thursday",
        "Friday"
    ]

### Detail

Please bare in mind that while _murex_ is optimised with concurrency and
streaming in mind, it's impossible to reverse an incomplete array. Thus all
all of STDIN must have been read and that file closed before `mtac` can
output.

In practical terms you shouldn't notice any difference except for when
STDIN is a long running process or non-standard stream (eg network pipe).

### See Also

* [`2darray` ](../commands/2darray.md):
  Create a 2D JSON array from multiple input sources
* [`a` (make array)](../commands/a.md):
  A sophisticated yet simple way to build an array or list
* [`append`](../commands/append.md):
  Add data to the end of an array
* [`ja`](../commands/ja.md):
  A sophisticated yet simply way to build a JSON array
* [`jsplit` ](../commands/jsplit.md):
  Splits STDIN into a JSON array based on a regex parameter
* [`len` ](../commands/len.md):
  Outputs the length of an array
* [`map` ](../commands/map.md):
  Creates a map from two data sources
* [`msort` ](../commands/msort.md):
  Sorts an array - data type agnostic
* [`prepend` ](../commands/prepend.md):
  Add data to the start of an array
* [`pretty`](../commands/pretty.md):
  Prettifies JSON to make it human readable
* [`ta`](../commands/ta.md):
  A sophisticated yet simple way to build an array of a user defined data-type
* [prefix](../commands/prefix.md):
  
* [suffix](../commands/suffix.md):
  