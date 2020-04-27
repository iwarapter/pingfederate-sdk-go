package pingfederate

import (
	"testing"
)

func TestDataStoreSimpleGetRequest(t *testing.T) {
	svc := basicClient()

	result1, resp1, err1 := svc.DataStores.GetDataStores()
	equals(t, err1, nil)
	equals(t, resp1.StatusCode, 200)
	equals(t, *(*result1.Items)[0].ConnectionUrl, "jdbc:hsqldb:${pf.server.data.dir}${/}hypersonic${/}ProvisionerDefaultDB;hsqldb.lock_file=false")
}

func TestDataStoreSimpleRequests(t *testing.T) {
	svc := basicClient()

	create := CreateDataStoreInput{}
	create.Body.JdbcDataStore = &JdbcDataStore{
		Type:                      String("JDBC"),
		MaskAttributeValues:       Bool(false),
		ConnectionUrl:             String("jdbc:h2:mem:test_mem"),
		DriverClass:               String("org.h2.Driver"),
		UserName:                  String("user"),
		Password:                  String("pass"),
		ValidateConnectionSql:     String(""),
		AllowMultiValueAttributes: Bool(true),
	}
	result, response, err := svc.DataStores.CreateDataStore(&create)
	equals(t, err, nil)
	equals(t, response.StatusCode, 201)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
	ID := *result.Id

	get := GetDataStoreInput{Id: ID}

	result, response, err = svc.DataStores.GetDataStore(&get)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")

	update := UpdateDataStoreInput{Id: ID}
	update.Body.JdbcDataStore = &JdbcDataStore{
		Type:                      String("JDBC"),
		MaskAttributeValues:       Bool(false),
		ConnectionUrl:             String("jdbc:h2:mem:test_mem"),
		DriverClass:               String("org.h2.Driver"),
		UserName:                  String("user"),
		Password:                  String("pass"),
		ValidateConnectionSql:     String(""),
		AllowMultiValueAttributes: Bool(true),
	}
	result, response, err = svc.DataStores.UpdateDataStore(&update)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
}
