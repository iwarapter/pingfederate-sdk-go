package pingfederate

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/iwarapter/pingfederate-sdk-go/services/version"

	"github.com/ory/dockertest"
)

var pfUrl *url.URL

func TestMain(m *testing.M) {

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	devOpsUser, devOpsUserExists := os.LookupEnv("PING_IDENTITY_DEVOPS_USER")
	devOpsKey, devOpsKeyExists := os.LookupEnv("PING_IDENTITY_DEVOPS_KEY")

	var options *dockertest.RunOptions
	dir, _ := os.Getwd()

	if devOpsUserExists && devOpsKeyExists {
		options = &dockertest.RunOptions{
			Hostname:   "pingfederate",
			Repository: "pingidentity/pingfederate",
			Mounts:     []string{dir + "/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip"},
			Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES", fmt.Sprintf("PING_IDENTITY_DEVOPS_USER=%s", devOpsUser), fmt.Sprintf("PING_IDENTITY_DEVOPS_KEY=%s", devOpsKey)},
			Tag:        "10.0.2-edge",
		}
	} else {
		options = &dockertest.RunOptions{
			Hostname:   "pingfederate",
			Repository: "pingidentity/pingfederate",
			Env:        []string{"PING_IDENTITY_ACCEPT_EULA=YES"},
			Mounts: []string{
				dir + "/pingfederate.lic:/opt/in/instance/server/default/conf/pingfederate.lic",
				dir + "/pingfederate-data.zip:/opt/in/instance/server/default/data/drop-in-deployer/data.zip",
			},
			Tag: "10.0.2-edge",
		}
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(options)
	resource.Expire(90)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	pool.MaxWait = time.Minute * 2

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		pfUrl, _ = url.Parse(fmt.Sprintf("https://localhost:%s", resource.GetPort("9999/tcp")))
		client := version.New("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)

		log.Println("Attempting to connect to PingFederate admin API....")
		_, _, err = client.GetVersion()
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	os.Setenv("PINGFEDERATE_BASEURL", fmt.Sprintf("https://localhost:%s", resource.GetPort("9999/tcp")))
	log.Println("Connected to PingFederate admin API....")
	code := m.Run()
	log.Println("Tests complete shutting down container")

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)

}

// func config(t *testing.T, filePath, path, body string) *PfClient {
// 	server := setupTestServer(t, filePath, path, body)
// 	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
// 	url, _ := url.Parse(server.URL)
// 	return NewClient("Administrator", "2Federate", url, "/pf-admin-api/v1", nil)
// }

//func basicClient() *client.PfClient {
//	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//	return client.NewClient("Administrator", "2Federate", pfUrl, "/pf-admin-api/v1", nil)
//}

// ok fails the test if an err is not nil.
func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
