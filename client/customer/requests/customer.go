package requests

type CustomerModelRequest struct {
	ID         string   `json:"id" `
	Name       string   `json:"name" binding:"required"`
	IdPersonal string   `json:"id_personal" binding:"required"`
	Age        string   `json:"age"`
	Address    string   `json:"address"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type CustomerIdRequest struct {
	IdPersonal string `bson:"id_personal" json:"id_personal"`
}

type UpdateCustomerModelRequest struct {
	ID         string   `json:"id" binding:"required"`
	Name       string   `json:"name"`
	IdPersonal string   `bson:"id_personal"`
	Age        string   `json:"age"`
	Address    string   `json:"address"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type DeleteCustomerRequest struct {
	ID string `json:"id" binding:"required"`
}

type AddTagsCustomerRequest struct {
	ID   string   `json:"id" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}

type DeleteTagsOfCustomerRequest struct {
	ID   string   `json:"id" binding:"required"`
	Tags []string `json:"tags" binding:"required"`
}
