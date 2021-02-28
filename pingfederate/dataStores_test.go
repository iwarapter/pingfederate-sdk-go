package pingfederate_test

import (
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"

	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"
)

func TestDataStoreSimpleGetRequest(t *testing.T) {
	svc := dataStores.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1").WithDebug(true))

	result1, resp1, err1 := svc.GetDataStores()
	ok(t, err1)
	equals(t, resp1.StatusCode, 200)
	equals(t, *(*result1.Items)[0].(*models.JdbcDataStore).ConnectionUrl, "jdbc:hsqldb:${pf.server.data.dir}${/}hypersonic${/}ProvisionerDefaultDB;hsqldb.lock_file=false")
}

func TestDataStoreSimpleRequests(t *testing.T) {
	svc := dataStores.New(config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1"))
	//svc.LogDebug = true

	create := dataStores.CreateJdbcDataStoreInput{}
	create.Body = models.JdbcDataStore{
		Type:                      pingfederate.String("JDBC"),
		MaskAttributeValues:       pingfederate.Bool(false),
		ConnectionUrl:             pingfederate.String("jdbc:h2:mem:test_mem"),
		DriverClass:               pingfederate.String("org.h2.Driver"),
		UserName:                  pingfederate.String("user"),
		Name:                      pingfederate.String("myDb"),
		Password:                  pingfederate.String("pass"),
		ValidateConnectionSql:     pingfederate.String(""),
		AllowMultiValueAttributes: pingfederate.Bool(true),
	}
	result, response, err := svc.CreateJdbcDataStore(&create)
	ok(t, err)
	equals(t, 201, response.StatusCode)
	equals(t, "jdbc:h2:mem:test_mem", *result.ConnectionUrl)
	equals(t, "JDBC", *result.Type)
	equals(t, "myDb", *result.Name)
	ID := *result.Id

	get := dataStores.GetJdbcDataStoreInput{Id: ID}

	resultDS, response, err := svc.GetJdbcDataStore(&get)
	ok(t, err)
	equals(t, 200, response.StatusCode)
	equals(t, "jdbc:h2:mem:test_mem", *resultDS.ConnectionUrl)
	equals(t, "JDBC", *resultDS.Type)
	equals(t, "myDb", *resultDS.Name)

	update := dataStores.UpdateJdbcDataStoreInput{Id: ID}
	update.Body = models.JdbcDataStore{
		Type:                      pingfederate.String("JDBC"),
		MaskAttributeValues:       pingfederate.Bool(false),
		ConnectionUrl:             pingfederate.String("jdbc:h2:mem:test_mem"),
		DriverClass:               pingfederate.String("org.h2.Driver"),
		UserName:                  pingfederate.String("user"),
		Password:                  pingfederate.String("pass"),
		ValidateConnectionSql:     pingfederate.String(""),
		AllowMultiValueAttributes: pingfederate.Bool(true),
	}
	result, response, err = svc.UpdateJdbcDataStore(&update)
	ok(t, err)
	equals(t, 200, response.StatusCode)
	equals(t, "jdbc:h2:mem:test_mem", *result.ConnectionUrl)
	equals(t, "JDBC", *result.Type)

	genericResult, _, err := svc.GetDataStore(&dataStores.GetDataStoreInput{Id: ID})
	ok(t, err)
	equals(t, pingfederate.String("jdbc:h2:mem:test_mem"), genericResult.JdbcDataStore.ConnectionUrl)

	deleteInput := dataStores.DeleteDataStoreInput{Id: ID}
	_, response, err = svc.DeleteDataStore(&deleteInput)
	ok(t, err)
	equals(t, 204, response.StatusCode)
}
