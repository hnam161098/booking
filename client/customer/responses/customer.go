package responses

type CustomerModelResponse struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	IdPersonal string   `json:"id_personal"`
	Age        string   `json:"age"`
	Address    string   `json:"address"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}
