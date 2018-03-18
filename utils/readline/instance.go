package readline

import (
	"regexp"
)

type Instance struct {
	// PasswordMask is what character to hide password entry behind.
	// Once enabled, set to 0 (zero) to disable the mask again.
	PasswordMask rune

	// SyntaxHighlight is a helper function to provide syntax highlighting.
	// Once enabled, set to nil to disable again.
	SyntaxHighlighter func([]rune) string

	// History is an interface for querying the readline history.
	// This is exposed as an interface to allow you the flexibility to define how
	// you want your history managed (eg file on disk, database, cloud, or even no
	// history at all). By default it uses a dummy interface that only stores
	// historic items in memory.
	History History

	// HistoryAutoWrite defines whether items automatically get written to history.
	// Enabled by default. Set to false to disable.
	HistoryAutoWrite bool // = true

	// TabCompleter is a simple function that offers completion suggestions.
	// It takes the readline line ([]rune) and cursor pos. Returns a prefix string
	// and an array of suggestions.
	TabCompleter func([]rune, int) (string, []string)

	// MaxTabCompletionRows is the maximum number of rows to display in the tab
	// completion grid.
	MaxTabCompleterRows int // = 4

	// SyntaxCompletion is used to autocomplete code syntax (like braces and
	// quotation marks). If you want to complete words or phrases then you might
	// be better off using the TabCompletion function.
	// SyntaxCompletion takes the line ([]rune) and cursor position, and returns
	// the new line and cursor position.
	SyntaxCompleter func([]rune, int) ([]rune, int)

	// HintText is a helper function which displays hint text the prompt.
	// HintText takes the line input from the promt and the cursor position.
	// It returns the hint text to display.
	HintText func([]rune, int) []rune

	// readline operating parameters
	prompt    string //  = ">>> "
	promptLen int    //= 4
	line      []rune
	pos       int

	// history
	lineBuf []rune
	histPos int

	// helper
	hintY int //= 0

	// tab completer
	modeTabGrid   bool
	tcPrefix      string
	tcSuggestions []string
	tcPosX        int
	tcPosY        int
	tcMaxX        int
	tcMaxY        int
	tcUsedY       int
	tcMaxLength   int

	// vim
	modeViMode  viMode //= vimInsert
	viIteration string

	// event
	evtKeyPress map[string]func(string, []rune, int) (bool, bool)
}

var (
	termWidth    int
	rxAnsiEscSeq *regexp.Regexp = regexp.MustCompile("\x1b\\[[0-9]+[a-zA-Z]")
)

func NewInstance() *Instance {
	rl := new(Instance)
	getTermWidth()

	rl.History = new(ExampleHistory)
	rl.HistoryAutoWrite = true
	rl.MaxTabCompleterRows = 4
	rl.prompt = ">>> "
	rl.promptLen = 4
	rl.evtKeyPress = make(map[string]func(string, []rune, int) (bool, bool))

	return rl
}
