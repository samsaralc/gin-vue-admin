package response

type TryonResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
	Success   bool   `json:"success"`
	Data      struct {
		ResultKey string `json:"result_key"`
		TaskList  []struct {
			Status     string   `json:"status"`
			ModelImage string   `json:"model_image"`
			TopsImage  string   `json:"tops_image"`
			PantsImage string   `json:"pants_image"`
			MaskImage  string   `json:"mask_image,omitempty"`
			Results    []string `json:"results,omitempty"` // 仅 SUCCESS 时返回
		} `json:"task_list"`
	} `json:"data"`
}

type TryOnResponse struct {
	ResultImage *string `json:"ResultImage"`
}
