package utils

import (
	"testing"
)

func TestIsIpv4(t *testing.T){

	host := "255.255.255.255"
	got := IsIpv4(host)

    if !got {
        t.Errorf("got false, wanted true")
    }
}

func TestIsNotIpv4(t *testing.T){

	host1 := "12346"
	host2 := "256.0.255.34"
	host3 := "-1.0.0.0"
	host4 := "0,0,0,0"
	got := IsIpv4(host1)
    if got {
        t.Errorf("got true, wanted false")
    }
	got = IsIpv4(host2)
	if got {
		t.Errorf("got true, wanted false")
    }
	got = IsIpv4(host3)
	if got {
		t.Errorf("got true, wanted false")
    }
	got = IsIpv4(host4)
	if got {
        t.Errorf("got true, wanted false")
    }
}