package main

type Message struct {
	Role             string  `json:"role"`
	ReasoningContent *string `json:"reasoning_content"` // Note: using pointer for nil case
	Content          string  `json:"content"`
}

type Choice struct {
	Index     int     `json:"index"`
	Message   Message `json:"message"`
	ToolCalls []any   `json:"tool_calls"`
	Logprobs  *any    `json:"logprobs"`
}

type ChatCompletion struct {
	ID             string   `json:"id"`
	Object         string   `json:"object"`
	Created        int64    `json:"created"`
	Model          string   `json:"model"`
	Choices        []Choice `json:"choices"`
	Usage          *Usage   `json:"usage"`
	PromptLogprobs *any     `json:"prompt_logprobs"`
}

type Usage struct {
	PromptTokens        int  `json:"prompt_tokens"`
	TotalTokens         int  `json:"total_tokens"`
	CompletionTokens    int  `json:"completion_tokens"`
	PromptTokensDetails *any `json:"prompt_tokens_details"`
}
