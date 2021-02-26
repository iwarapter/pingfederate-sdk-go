package pingfederate_test

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	pf "github.com/iwarapter/pingfederate-sdk-go/pingfederate"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/config"
	"github.com/iwarapter/pingfederate-sdk-go/pingfederate/models"
	"github.com/iwarapter/pingfederate-sdk-go/services/serverSettings"
)

func TestMain(m *testing.M) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cfg := config.NewConfig().WithUsername("Administrator").WithPassword("2FederateM0re").WithEndpoint("https://localhost:9999/pf-admin-api/v1")
	svc := serverSettings.New(cfg)

	_, _, err := svc.UpdateServerSettings(&serverSettings.UpdateServerSettingsInput{
		Body: models.ServerSettings{
			FederationInfo: &models.FederationInfo{
				BaseUrl:       pf.String("https://localhost:9031"),
				Saml2EntityId: pf.String("testing"),
			},
			RolesAndProtocols: &models.RolesAndProtocols{
				IdpRole: &models.IdpRole{
					Enable: pf.Bool(true),
					Saml20Profile: &models.SAML20Profile{
						Enable: pf.Bool(true),
					},
				},
				OauthRole: &models.OAuthRole{
					EnableOauth: pf.Bool(true),
				},
				SpRole: &models.SpRole{
					Enable: pf.Bool(true),
					Saml20Profile: &models.SpSAML20Profile{
						Enable: pf.Bool(true),
					},
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("unable to setup PF %s", err)
	}
	os.Exit(m.Run())
}

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
