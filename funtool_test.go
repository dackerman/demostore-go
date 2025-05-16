// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dackerman/demostore-go"
	"github.com/dackerman/demostore-go/internal/testutil"
	"github.com/dackerman/demostore-go/option"
)

func TestFuntoolSetDarkmode(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := dackermanstore.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("123e4567-e89b-12d3-a456-426614174000"),
		option.WithOrgID("my_org"),
	)
	_, err := client.Funtools.SetDarkmode(context.TODO(), dackermanstore.FuntoolSetDarkmodeParams{
		Darkmode: true,
	})
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
