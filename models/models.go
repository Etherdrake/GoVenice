package models

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
	Llama3_3_70b = Model{
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

	Llama3_2_3b = Model{
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

	Dolphin2_9_2_Qwen2_72b = Model{
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

	Llama3_1_405b = Model{
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

	Qwen2_5_Coder_32b = Model{
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

	DeepseekR1_Llama_70b = Model{
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

	DeepseekR1_671b = Model{
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

	Qwen2_5_VL = Model{
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

type ModelRegistry struct {
	llama_3_3_70b           *Model
	llama_3_2_3b            *Model
	dolphin_2_9_2_qwen2_72b *Model
	llama_3_1_405b          *Model
	qwen_2_5_coder_32b      *Model
	deepseek_r1_llama_70b   *Model
	deepseek_r1_671b        *Model
	qwen_2_5_vl             *Model
}

var Models = ModelRegistry{
	llama_3_3_70b:           &Llama3_3_70b,
	llama_3_2_3b:            &Llama3_2_3b,
	dolphin_2_9_2_qwen2_72b: &Dolphin2_9_2_Qwen2_72b,
	llama_3_1_405b:          &Llama3_1_405b,
	qwen_2_5_coder_32b:      &Qwen2_5_Coder_32b,
	deepseek_r1_llama_70b:   &DeepseekR1_Llama_70b,
	deepseek_r1_671b:        &DeepseekR1_671b,
	qwen_2_5_vl:             &Qwen2_5_VL,
}

func GetModels() *ModelRegistry {
	return &Models
}
