package tokenizer

import "testing"

func TestGetModelSwitch(t *testing.T) {
	modelName := "some_unknown_model"

	model, _ := GetModelSwitch(modelName, true, true)

	if model != nil {
		t.Errorf("FAILED: Expected nil, got %d", model)
	} else {
		t.Log("PASSED")
	}

	// check if correct model implements Model interface
	modelName = "bert"

	model, err := GetModelSwitch(modelName, true, true)

	if err != nil {
		t.Error(err)
	}

	var i interface{} = model
	_, ok := i.(Model)

	if !ok {
		t.Errorf("FAILED: %d does not implement Model interface", model)
	} else {
		t.Log("PASSED")
	}

	// check if all correct models return non-nil values
	modelNames := []string{
		"bert", "gpt2", "roberta",
	}

	for _, modelName := range modelNames {
		model, _ = GetModelSwitch(modelName, true, true)
		if model == nil {
			t.Errorf("FAILED: Model named %s returned nil", modelName)
		}
	}
	t.Log("PASSED")

}
