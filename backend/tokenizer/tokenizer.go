package tokenizer

import (
	"github.com/sugarme/tokenizer"
	"github.com/sugarme/tokenizer/pretrained"
)

type model interface {
	EncodeSingle(input string, addSpecialTokensOpt ...bool) (*tokenizer.Encoding, error)
}

func GetModelSwitch(model_string string, addPrefixSpace bool, trimOffsets bool) model {
	switch model_string {
	case "bert":
		return pretrained.BertBaseUncased()
	case "gpt2":
		return pretrained.GPT2(trimOffsets, trimOffsets)
	case "roberta":
		return pretrained.RobertaBase(trimOffsets, trimOffsets)
	default:
		return nil
	}
}
