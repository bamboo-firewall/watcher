package handler

import (
	"fmt"
	"testing"
)

func TestEvent(t *testing.T) {
	var e1 = `/calico/resources/v3/projectcalico.org/hostendpoints/10.110.31.233`
	var e2 = `/calico/resources/v3/projectcalico.org/globalnetworkpolicies/default.airflow-app`
	var e3 = `/calico/resources/v3/projectcalico.org/globalnetworksets/bank-payment-agribank`
	want1 := "hostendpoints"
	key1 := "10.110.31.233"
	want2 := "globalnetworkpolicies"
	key2 := "default.airflow-app"
	want3 := "globalnetworksets"
	key3 := "bank-payment-agribank"
	got1, k1, _ := Event(e1)
	fmt.Printf("%s, %s", got1, k1)
	if got1 != want1 && k1 != key1 {
		t.Errorf("got %q, wanted %q", got1, want1)
	}
	got2, k2, _ := Event(e2)
	if got2 != want2 && k2 != key2 {
		t.Errorf("got %q, wanted %q", got2, want2)
	}
	got3, k3, _ := Event(e3)
	if got3 != want3 && k3 != key3 {
		t.Errorf("got %q, wanted %q", got3, want3)
	}
}
