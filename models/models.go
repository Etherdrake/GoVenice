package main

type Model struct {
	ID        string
	Type      string
	Created   int64
	OwnedBy   string
	ModelSpec ModelSpec
}

type ModelSpec struct {
	AvailableContextTokens  int
	SupportsFunctionCalling bool
	Traits                  []string
	ModelSource             string
}

var (
	model_llama_3_3_70b = Model{
		ID:      "llama-3.3-70b",
		Type:    "text",
		Created: 1733768349,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  65536,
			SupportsFunctionCalling: true,
			Traits:                  []string{"function_calling_default", "default"},
			ModelSource:             "https://huggingface.co/meta-llama/Llama-3.3-70B-Instruct",
		},
	}

	model_llama_3_2_3b = Model{
		ID:      "llama-3.2-3b",
		Type:    "text",
		Created: 1727966436,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  131072,
			SupportsFunctionCalling: true,
			Traits:                  []string{"fastest"},
			ModelSource:             "https://huggingface.co/meta-llama/Llama-3.2-3B",
		},
	}

	model_dolphin_2_9_2_qwen2_72b = Model{
		ID:      "dolphin-2.9.2-qwen2-72b",
		Type:    "text",
		Created: 1726869022,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  32768,
			SupportsFunctionCalling: false,
			Traits:                  []string{"most_uncensored"},
			ModelSource:             "https://huggingface.co/cognitivecomputations/dolphin-2.9.2-qwen2-72b",
		},
	}

	model_llama_3_1_405b = Model{
		ID:      "llama-3.1-405b",
		Type:    "text",
		Created: 1730396371,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  63920,
			SupportsFunctionCalling: false,
			Traits:                  []string{"most_intelligent"},
			ModelSource:             "https://huggingface.co/meta-llama/Meta-Llama-3.1-405B-Instruct",
		},
	}

	model_qwen_2_5_coder_32b = Model{
		ID:      "qwen-2.5-coder-32b",
		Type:    "text",
		Created: 1731628653,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  32768,
			SupportsFunctionCalling: false,
			Traits:                  []string{"default_code"},
			ModelSource:             "https://huggingface.co/Qwen/Qwen2.5-Coder-32B-Instruct-GGUF",
		},
	}

	model_deepseek_r1_llama_70b = Model{
		ID:      "deepseek-r1-llama-70b",
		Type:    "text",
		Created: 1737766445,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  65536,
			SupportsFunctionCalling: false,
			Traits:                  []string{},
			ModelSource:             "https://huggingface.co/deepseek-ai/DeepSeek-R1-Distill-Llama-70B",
		},
	}

	model_deepseek_r1_671b = Model{
		ID:      "deepseek-r1-671b",
		Type:    "text",
		Created: 1738690625,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  131072,
			SupportsFunctionCalling: false,
			Traits:                  []string{},
			ModelSource:             "https://huggingface.co/deepseek-ai/DeepSeek-R1",
		},
	}

	model_qwen_2_5_vl = Model{
		ID:      "qwen-2.5-vl",
		Type:    "text",
		Created: 1739074852,
		OwnedBy: "venice.ai",
		ModelSpec: ModelSpec{
			AvailableContextTokens:  32768,
			SupportsFunctionCalling: false,
			Traits:                  []string{"default_vision"},
			ModelSource:             "https://huggingface.co/Qwen/Qwen2.5-VL-72B-Instruct",
		},
	}
)

func GetModels() map[string]*Model {
	return map[string]*Model{
		"llama-3.3-70b":           &model_llama_3_3_70b,
		"llama-3.2-3b":            &model_llama_3_2_3b,
		"dolphin-2.9.2-qwen2-72b": &model_dolphin_2_9_2_qwen2_72b,
		"llama-3.1-405b":          &model_llama_3_1_405b,
		"qwen-2.5-coder-32b":      &model_qwen_2_5_coder_32b,
		"deepseek-r1-llama-70b":   &model_deepseek_r1_llama_70b,
		"deepseek-r1-671b":        &model_deepseek_r1_671b,
		"qwen-2.5-vl":             &model_qwen_2_5_vl,
	}
}
