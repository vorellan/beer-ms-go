package modeltests

import (
	"testing"
	"github.com/vorellan/beer-ms-go/beer/controllers"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vorellan/beer-ms-go/beer/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindBeersWhenCallDBThenReturnCorrectLen(t *testing.T) {

	var server = controllers.Server{}

	beer := models.Beer{}
	beers, err := beer.FindBeers(server.DB)
	if err != nil {
		t.Errorf("error to get beeers: %v\n", err)
		return
	}
	assert.Equal(t, len(*beers), 2)
}



