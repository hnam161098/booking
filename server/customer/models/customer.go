package models

const IndexCustomer string = "customer"

type Customer struct {
	ID         string   `bson:"_id, omitempty" json:"_id"`
	Name       string   `bson:"name" json:"name"`
	IdPersonal string   `bson:"id_personal" json:"id_personal"`
	Age        string   `bson:"age" json:"age"`
	Address    string   `bson:"address" json:"address"`
	Tags       []string `bson:"tags" json:"tags"`
	CreatedAt  string   `bson:"created_at" json:"created_at"`
	UpdatedAt  string   `bson:"updated_at" json:"updated_at"`
}

type CustomerRequest struct {
	IdPersonal string `bson:"id_personal" json:"id_personal"`
}

type CustomerESResult struct {
	Hits map[string]interface{} `json:"hits"`
}

type CustomerESModel struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Age        string   `json:"age"`
	Address    string   `json:"address"`
	IdPersonal string   `json:"id_personal"`
	Tags       []string `json:"tags"`
	CreatedAt  string   `json:"created_at"`
	UpdatedAt  string   `json:"updated_at"`
}

type AddTags struct {
	ID   string   `json:"_id"`
	Tags []string `json:"tags"`
}

type DeleteTags struct {
	ID   string   `json:"_id"`
	Tags []string `json:"tags"`
}
