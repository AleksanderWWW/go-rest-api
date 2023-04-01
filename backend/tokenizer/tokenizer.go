package tokenizer

import (
	"fmt"

	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

type Model interface {
	EncodeSingle(input string, addSpecialTokensOpt ...bool) (*tokenizer.Encoding, error)
}

func GetModelSwitch(model_string string, addPrefixSpace bool, trimOffsets bool) (Model, error) {
	switch model_string {
	case "bert":
		return pretrained.BertBaseUncased(), nil
	case "gpt2":
		return pretrained.GPT2(trimOffsets, trimOffsets), nil
	case "roberta":
		return pretrained.RobertaBase(trimOffsets, trimOffsets), nil
	default:
		return nil, fmt.Errorf("unknown model name %s", model_string)
	}
}
