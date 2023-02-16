package tokenizer

import "testing"

func TestGetModelSwitch(t *testing.T) {
	model_name := "some_unknown_model"

	model := GetModelSwitch(model_name, true, true)

	if model != nil {
		t.Errorf("FAILED: Expected nil, got %d", model)
	} else {
		t.Log("PASSED")
	}

}
