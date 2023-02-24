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

	// check if correct model implements Model interface
	model_name = "bert"

	model = GetModelSwitch(model_name, true, true)

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
		model = GetModelSwitch(modelName, true, true)
		if model == nil {
			t.Errorf("FAILED: Model named %s returned nil", modelName)
		}
	}
	t.Log("PASSED")

}
