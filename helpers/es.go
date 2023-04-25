package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grpc/database_connection"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/esapi"
)

/* các API GET thì nên lấy ở Elasticsearch */

/*
	các API POST thì nên lấy ở các DB, đồng thời đồng bộ cùng với Elasticsearch
    để không bị sai sót data
*/

func InsertDocumentES(index string, document interface{}, ID string) error {
	ESconn := database_connection.ConnectES()
	doc, err := json.Marshal(document)
	if err != nil {
		return err
	}

	var req = esapi.IndexRequest{}
	if ID != "" {
		req = esapi.IndexRequest{
			Index:      index,
			DocumentID: ID,
			Body:       strings.NewReader(string(doc)),
			Refresh:    "true",
		}
	} else {
		req = esapi.IndexRequest{
			Index:   index,
			Body:    strings.NewReader(string(doc)),
			Refresh: "true",
		}
	}

	res, err1 := req.Do(context.Background(), ESconn)
	if err1 != nil {
		return err1
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println(res.Status())
		log.Println(res.String())
	}
	return nil
}

func UpdateES(index string, query map[string]interface{}, idDoc string) error {
	conn := database_connection.ConnectES()
	bodyRQ, _ := json.Marshal(query)

	req := esapi.UpdateRequest{
		Index:      index,
		DocumentID: idDoc,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, bodyRQ))),
	}
	res, err := req.Do(context.Background(), conn)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}

func QueryES(index string, query map[string]interface{}) (map[string]interface{}, error) {
	conn := database_connection.ConnectES()
	bodyRQ, _ := json.Marshal(query)

	res, err := conn.Search(
		conn.Search.WithContext(context.Background()),
		conn.Search.WithIndex(index),
		conn.Search.WithBody(strings.NewReader(string(bodyRQ))),
		conn.Search.WithTrackTotalHits(true),
		conn.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		responseES := make(map[string]interface{})
		if err := json.NewDecoder(res.Body).Decode(&responseES); err == nil {
			return responseES, nil
		}
	}
	return nil, nil
}

func DeleteES(index string, query map[string]interface{}) error {
	conn := database_connection.ConnectES()
	bodyRQ, _ := json.Marshal(query)
	res, err := conn.DeleteByQuery([]string{index}, bytes.NewReader(bodyRQ))
	if err != nil {
		return err
	} else {
		defer res.Body.Close()
		return nil
	}
}
