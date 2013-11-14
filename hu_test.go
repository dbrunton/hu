package hu

import "testing"
import "fmt"

func TestSanity(t *testing.T) {

        if false {
                t.Errorf("Unsorted manifests don't match.")
        }


}

var S1 = &signature{"S1", nil, nil}
var D1 = &definition{S1, nil}
var t1 = &term{"t1", nil}
var D2 = &definition{S1, []*term{t1, t1, t1}}
var T2 = &term{"T2", []*definition{D2}}

func TestDefinitionString(t *testing.T) {
	if fmt.Sprintf("%s", D2) != "(S1): T1 t1 t1." {
		t.Errorf("Expected '(S1): T1 t1 t1.' Got ", fmt.Sprintf("%s", D2))
	}
}
