package chap02

import (
	"testing"
	"time"
)

func TestTimeFunction(t *testing.T) {
	var now time.Time = time.Now()
	year := now.Year()
	t.Log(year)
}
