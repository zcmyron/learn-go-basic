package submodels

import (
	"com.jtthink/models"
	"github.com/pquerna/ffjson/ffjson"
)

type SportsNews struct {
	Tags []string  //array
	models.NewsModel
}
func (sn SportsNews) ToJSON() string {
	result,err:=ffjson.Marshal(sn)
	if err!=nil {
		return err.Error()
	} else {
		return string(result)
	}
}
