package yaml_test

import (
	"testing"

	. "ireul.com/check"
)

func Test(t *testing.T) { TestingT(t) }

type S struct{}

var _ = Suite(&S{})
