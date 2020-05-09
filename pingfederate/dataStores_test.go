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
	svc.LogDebug = true

	create := CreateJdbcDataStoreInput{}
	create.Body = JdbcDataStore{
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
	result, response, err := svc.DataStores.CreateJdbcDataStore(&create)
	equals(t, err, nil)
	equals(t, response.StatusCode, 201)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
	equals(t, *result.Name, "myDb")
	ID := *result.Id

	get := GetDataStoreInput{Id: ID}

	resultDS, response, err := svc.DataStores.GetDataStore(&get)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *resultDS.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *resultDS.Type, "JDBC")
	equals(t, *resultDS.Name, "myDb")

	update := UpdateJdbcDataStoreInput{Id: ID}
	update.Body = JdbcDataStore{
		Type:                      String("JDBC"),
		MaskAttributeValues:       Bool(false),
		ConnectionUrl:             String("jdbc:h2:mem:test_mem"),
		DriverClass:               String("org.h2.Driver"),
		UserName:                  String("user"),
		Password:                  String("pass"),
		ValidateConnectionSql:     String(""),
		AllowMultiValueAttributes: Bool(true),
	}
	result, response, err = svc.DataStores.UpdateJdbcDataStore(&update)
	equals(t, err, nil)
	equals(t, response.StatusCode, 200)
	equals(t, *result.ConnectionUrl, "jdbc:h2:mem:test_mem")
	equals(t, *result.Type, "JDBC")
}
