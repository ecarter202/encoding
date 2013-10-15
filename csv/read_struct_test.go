package csv

import (
	"reflect"
	"strings"
	"testing"
)

func TestReadStructAll(t *testing.T) {
	for n, tt := range tests {

		var s = []testStruct{}

		w := NewReader(strings.NewReader(tt.CSV))
		err := w.ReadStructAll(&s)
		if err != nil {
			return
		}

		if !reflect.DeepEqual(s, tt.Struct) {
			t.Errorf("#%d: out:%v, want %v", n, s, tt.Struct)
		}
	}
}