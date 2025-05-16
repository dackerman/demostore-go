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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Products.List(context.TODO(), dackermanstore.ProductListParams{
		OrgID: dackermanstore.F("org_id"),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, product := range page.Data {
		t.Logf("%+v\n", product.ProductID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, product := range page.Data {
			t.Logf("%+v\n", product.ProductID)
		}
	}
}
