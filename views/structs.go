package views

type ActionRequest struct {
	Type    string                 `json:"type"`
	Game    string                 `json:"game"`
	Details map[string]interface{} `json:"details"`
}
