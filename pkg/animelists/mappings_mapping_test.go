package animelists

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestMapping_ParseRange(t *testing.T) {
	m := &Mapping{
		Start:  "1",
		End:    "20",
		Offset: "0",
	}
	anidbToTvdb, ok := m.parseRange()
	if !ok {
		t.Error("failed to parse range")
	}
	if len(anidbToTvdb) != 20 {
		t.Error("failed to parse range")
	}

	spew.Dump(anidbToTvdb)
}

func TestMapping_ParseValue(t *testing.T) {
	m := &Mapping{
		Value: "1-2;2-3;3-4+5;4-0;5-0;",
	}
	anidbToTvdb, ok := m.parseValue()
	if !ok {
		t.Error("failed to parse value")
	}
	if len(anidbToTvdb) != 5 {
		t.Error("failed to parse value")
	}

	if len(anidbToTvdb[3]) != 2 {
		t.Error("failed to parse value")
	}

	spew.Dump(anidbToTvdb)
}
