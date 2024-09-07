package resourcecontroller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	ResourceService "github.com/MohitVachhani/go-learn/cmd/service/resource"

	ResourceInterface "github.com/MohitVachhani/go-learn/pkg/structs/resource"
	"go.mongodb.org/mongo-driver/bson"
)

func GetResource(w http.ResponseWriter, r *http.Request) {

	// body parameters
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	var getResourceByIdInput ResourceInterface.GetResourceInputInterface
	json.Unmarshal(body, &getResourceByIdInput)

	resourcePayload := ResourceService.GetResourceById(getResourceByIdInput)

	json.NewEncoder(w).Encode(bson.M{"success": true, "resource": resourcePayload})

}
