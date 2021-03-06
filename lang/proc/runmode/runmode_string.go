// Code generated by "stringer -type=RunMode"; DO NOT EDIT.

package runmode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Normal-0]
	_ = x[Shell-1]
	_ = x[Try-2]
	_ = x[TryPipe-3]
	_ = x[Evil-4]
}

const _RunMode_name = "NormalShellTryTryPipeEvil"

var _RunMode_index = [...]uint8{0, 6, 11, 14, 21, 25}

func (i RunMode) String() string {
	if i < 0 || i >= RunMode(len(_RunMode_index)-1) {
		return "RunMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RunMode_name[_RunMode_index[i]:_RunMode_index[i+1]]
}
