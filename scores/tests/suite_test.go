package tests

import (
	"github.com/glebarez/sqlite"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	getbycategories "go-scores/gen/get_by_categories"
	"go-scores/scores/interfaces/controllers"
	"gorm.io/gorm"
	"testing"
)

var client *getbycategories.Client
var DB *gorm.DB

var _ = BeforeSuite(func() {
	{
		var err error
		DB, err = gorm.Open(sqlite.Open("../../database.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		DeferCleanup(func() {
			sqlDB, _ := DB.DB()
			_ = sqlDB.Close()
		})
	}

	endpoint := getbycategories.NewGetAggregatedScoresEndpoint(controllers.NewGetByCategories(DB))
	client = getbycategories.NewClient(endpoint)
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "wallet/tests")
}
