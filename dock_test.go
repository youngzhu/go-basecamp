package basecamp

import (
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	atype := reflect.TypeOf(TypeKanbanBoard)

	want := "dockType"
	if atype.Name() != want {
		t.Errorf("want: %s, but got: %s", want, atype)
	}
}
