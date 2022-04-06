package access

import (
	"fmt"
	"net/http"
	"trackingApp/models"
	"io/ioutil"
	"encoding/json"
	. "trackingApp/repositories/accessRepo"
	"trackingApp/shared/dal"
	"os"
	. "trackingApp/shared/utils"
	. "trackingApp/shared/customError"
)

type AccessController struct {}

func (ac AccessController) Post(w http.ResponseWriter, req *http.Request) {
	access,err := ac.Validate(req)

	if err != nil {
		output, _ :=json.Marshal(err)
		fmt.Fprintf(w,string(output))
		return
	}

	// var ar AccessRepoBase = &AccessRepo{&dal.DbMongo{os.Getenv("DATABASE_URL")}}
	var ar AccessRepoBase = &AccessRepo{&dal.DbPostgres{os.Getenv("DATABASE_URL")}}
	ar.Save(access)
	output,_ := json.Marshal(&access)
	fmt.Fprintf(w,string(output))
}

func (ac AccessController) Validate(req *http.Request) (models.Access,error) {
	access := models.Access{}
	body,err := ioutil.ReadAll(req.Body)
	if err != nil{
		return access,err
	}
	err = BlockBot(req)
	if err != nil{
		return access,err
	}
	err = json.Unmarshal(body,&access)
	if err != nil{
		return access,err
	}
	if !IsIpv4(access.Ip){
		return access,&InvalidData{"Ip inválido"}
	}
	return access,nil
}


