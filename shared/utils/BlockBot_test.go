package utils

import (
	"testing"
	. "trackingApp/shared/customError"
	// "github.com/google/go-cmp/cmp"

)

func TestBlockBot(t *testing.T){

	ua := "Bot"
	got := BlockBot(ua)
	want := &AccessDenied{"Bot detectado"}

    if got.Error() != want.Error() {
        t.Errorf("got %q, wanted %q", got, want)
    }
}