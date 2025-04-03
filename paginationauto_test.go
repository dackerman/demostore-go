// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore_test

import (
	"context"
	"os"
	"testing"

	"github.com/dackerman/demostore-go"
	"github.com/dackerman/demostore-go/internal/testutil"
	"github.com/dackerman/demostore-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Products.ListAutoPaging(context.TODO(), dackermanstore.ProductListParams{
		OrgID: dackermanstore.F("org_id"),
	})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		product := iter.Current()
		t.Logf("%+v\n", product.ProductID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
