package builtins

// This is where you can add or remove built in functions. Imports here require underscoring.
// Each builtin package should include a README.md with details about its use. However the code should also be readable
// so take a look through the .go files if you're still undecided about whether to include a builtin or not.
//
// My recommendation is to keep everything in this file included, but disable any of the Go files prefixed `optional`.
import (
	// Docs
	_ "github.com/lmorg/murex/builtins/docs" // docs for builtin murex functions

	// CLI functions:
	_ "github.com/lmorg/murex/builtins/core/arraytools"        // tools for working with arrays and maps
	_ "github.com/lmorg/murex/builtins/core/datatools"         // utilities for manipulating structured data
	_ "github.com/lmorg/murex/builtins/core/httpclient"        // builtins for http
	_ "github.com/lmorg/murex/builtins/core/io"                // OS IO builtins
	_ "github.com/lmorg/murex/builtins/core/management"        // murex management builtins
	_ "github.com/lmorg/murex/builtins/core/mkarray"           // rapidly makes arrays
	_ "github.com/lmorg/murex/builtins/core/open"              // `open` and default handlers
	_ "github.com/lmorg/murex/builtins/core/random"            // random data builtin
	_ "github.com/lmorg/murex/builtins/core/ranges"            // working with ranges within arrays (`@[..]`)
	_ "github.com/lmorg/murex/builtins/core/shellautocomplete" // shell autocompleter tools
	_ "github.com/lmorg/murex/builtins/core/shellconfig"       // shell config tools
	_ "github.com/lmorg/murex/builtins/core/shellpipe"         // shell tools for named pipes
	_ "github.com/lmorg/murex/builtins/core/shellruntime"      // shell runtime inspection
	_ "github.com/lmorg/murex/builtins/core/shelltest"         // testing framework for murex shell scripts
	_ "github.com/lmorg/murex/builtins/core/structs"           // control structures
	_ "github.com/lmorg/murex/builtins/core/textmanip"         // text manipulation builtins
	_ "github.com/lmorg/murex/builtins/core/typemgmt"          // type handling and management builtins

	// Events:
	_ "github.com/lmorg/murex/builtins/events/onFileSystemChange" // file system watcher event type
	_ "github.com/lmorg/murex/builtins/events/onKeyPress"         // readline key-press event type
	_ "github.com/lmorg/murex/builtins/events/onSecondsElapsed"   // timer-based event type

	// Pipes:
	_ "github.com/lmorg/murex/builtins/pipes/file"    // writing to a file (required for history)
	_ "github.com/lmorg/murex/builtins/pipes/null"    // null interface (required!)
	_ "github.com/lmorg/murex/builtins/pipes/streams" // standard interfaces for pipes (required!)
	_ "github.com/lmorg/murex/builtins/pipes/term"    // writing to the terminal / TTY (required!)

	// Data types:
	_ "github.com/lmorg/murex/builtins/types/binary"  // basic data type for handing binary data
	_ "github.com/lmorg/murex/builtins/types/generic" // generic (string) data type
	_ "github.com/lmorg/murex/builtins/types/json"    // JSON data type
	_ "github.com/lmorg/murex/builtins/types/numeric" // formatting numeric data types (int, float, number)
	_ "github.com/lmorg/murex/builtins/types/string"  // string data type (soon to be depreciated)
)
