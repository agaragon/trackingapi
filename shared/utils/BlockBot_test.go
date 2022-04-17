package utils

import (
	"testing"
	. "trackingApp/shared/customError"
)

func TestBlockBot(t *testing.T){

	ua := "Bot"
	got := BlockBot(ua)
	want := &AccessDenied{"Bot detectado"}

    if got.Error() != want.Error() {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestDontBlockMozilla(t *testing.T){
	
		ua := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"
		got := BlockBot(ua)
	
		if got != nil {
			t.Errorf("got %q, wanted nil", got)
		}

}