package error

import (
	origin_errors "errors"
	"testing"

	"github.com/pkg/errors"
)

func TestErr(t *testing.T) {

	//t.Errorf("nothing")
	// t.Log(genErr().Error())
	// fmt.Errorf("haha:%v", genErr())
	// errors.Errorf()

	err := biz2222()
	if err != nil {
		t.Logf("%v", err)
		t.Logf("+v %+v", err)
		t.Logf("cause err: %+v, %v", errors.Unwrap(err), errors.Unwrap(err) == RecordNotFound)
		t.Logf("cause err equal?: %v", errors.Unwrap(err) == RecordNotFound)
		t.Logf("cause err equal?Is: %v", origin_errors.Is(err, RecordNotFound))

		t.Logf("cause err equal????: %v", errors.Cause(err) == RecordNotFound)

		// 加包装 fmt.Errorf("message %w",err)
	}
}
