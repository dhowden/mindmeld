package mindmeld_test

import (
	"testing"

	"github.com/dhowden/mindmeld"
)

func TestNewTokenSource(t *testing.T) {
	ts := mindmeld.NewTokenSource()
	_ = ts.Token()
}
