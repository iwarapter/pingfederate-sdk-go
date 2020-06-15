package pingfederate

import (
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"testing"

	"github.com/iwarapter/pingfederate-sdk-go/services/dataStores"
)

func TestDataStoreSimpleGetRequest(t *testing.T) {
	svc := dataStores.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)

	result1, resp1, err1 := svc.GetDataStores()
	equals(t, err1, nil)
	equals(t, resp1.StatusCode, 200)
	equals(t, *(*result1.Items)[0].ConnectionUrl, "jdbc:hsqldb:${pf.server.data.dir}${/}hypersonic${/}ProvisionerDefaultDB;hsqldb.lock_file=false")
}

func TestDataStoreSimpleRequests(t *testing.T) {
	svc := dataStores.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)
	//svc.LogDebug = true

	create := dataStores.CreateJdbcDataStoreInput{}
	create.Body = models.JdbcDataStore{
		Type:                      String("JDBC"),
		MaskAttributeValues:       Bool(false),
		ConnectionUrl:             String("jdbc:h2:mem:test_mem"),
		DriverClass:               String("org.h2.Driver"),
		UserName:                  String("user"),
		Name:                      String("myDb"),
		Password:                  String("pass"),
		ValidateConnectionSql:     String(""),
		AllowMultiValueAttributes: Bool(true),
	}
	result, response, err := svc.CreateJdbcDataStore(&create)
	equals(t, err, nil)
	equals(t, response.StatusCode, 201)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
	equals(t, *result.Name, "myDb")
	ID := *result.Id

	get := dataStores.GetJdbcDataStoreInput{Id: ID}

	resultDS, response, err := svc.GetJdbcDataStore(&get)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *resultDS.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *resultDS.Type, "JDBC")
	equals(t, *resultDS.Name, "myDb")

	update := dataStores.UpdateJdbcDataStoreInput{Id: ID}
	update.Body = models.JdbcDataStore{
		Type:                      String("JDBC"),
		MaskAttributeValues:       Bool(false),
		ConnectionUrl:             String("jdbc:h2:mem:test_mem"),
		DriverClass:               String("org.h2.Driver"),
		UserName:                  String("user"),
		Password:                  String("pass"),
		ValidateConnectionSql:     String(""),
		AllowMultiValueAttributes: Bool(true),
	}
	result, response, err = svc.UpdateJdbcDataStore(&update)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
}
