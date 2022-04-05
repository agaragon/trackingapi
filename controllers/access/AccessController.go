package access

import (
	"fmt"
	"net/http"
	"trackingApp/models"
	"io/ioutil"
	"encoding/json"
	. "trackingApp/repositories/accessRepo"
)

type AccessController struct {}

func (ac AccessController) Post(w http.ResponseWriter, req *http.Request) {
	access := models.Access{}
	body,_ := ioutil.ReadAll(req.Body)
	
	err := json.Unmarshal(body,&access)
	if err != nil {
		fmt.Fprintf(w,"{\"error_message\":\"requisição inválida\"}")
		return
	}
	ar := AccessRepo{}
	ar.Save(access)
	
	// output,_ := json.Marshal(access)
	fmt.Fprintf(w,"Acesso salvo")
}

// func (ac AccessController) Validation(req *http.Request) {
// 	access := models.Access{}
// 	body,err := ioutil.ReadAll(req.Body)
	
// 	err = json.Unmarshal(body,&access)
// 	if err != nil {
// 		fmt.Fprintf(w,"Dados em formato inválido")
// 		log.Fatal("Request Inválida")
// 	}
// }