package avutil

import (
	"fmt"
	"testing"
)

func TestStrError(t *testing.T) {
	t.Log("error string:", StrError(-22))

	var code = errorCode{fmt.Errorf("123"), EAGAIN}
	code.err = Error(code.code)
	t.Log(code)
}
