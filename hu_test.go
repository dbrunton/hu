package hu

import "testing"
import "fmt"

func TestSanity(t *testing.T) {

        if false {
                t.Errorf("Something wacky is wrong.")
        }


}

var S1 = &signature{"S1", nil, nil}
var S2 = &signature{"S2", signatures{S1, S1}, nil}
var S3 = &signature{"S3", signatures{S1, S2}, signatures{S2, S2}}
var D1 = &definition{S1, nil}
var t1 = &term{"t1", nil}
var D2 = &definition{S1, []*term{t1, t1, t1}}
var t2 = &term{"t2", []*definition{D2}}
var t3 = &term{"t3", []*definition{D2, D1}}

func TestDefinitionString(t *testing.T) {
	if fmt.Sprintf("%s", D2) != "(S1): T1 t1 t1." {
		t.Errorf("Expected '(S1): T1 t1 t1.' Got ", fmt.Sprintf("%s", D2))
	}
}

func TestInputs(t *testing.T) {
	if S3.Inputs() != "S1 S2" {
		t.Errorf("Inputs wrong, expected 'S1 S2', got ", S3.Inputs())
	}
}

func TestOutputs(t *testing.T) {
	if S3.Outputs() != "S2 S2" {
		t.Errorf("Inputs wrong, expected 'S1 S2', got ", S3.Inputs())
	}
}
