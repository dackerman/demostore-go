// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore_test

import (
	"context"
	"os"
	"testing"

	"github.com/dackerman/demostore-go/v2"
	"github.com/dackerman/demostore-go/v2/internal/testutil"
	"github.com/dackerman/demostore-go/v2/option"
)

func TestUsage(t *testing.T) {
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
	product, err := client.Products.New(context.TODO(), dackermanstore.ProductNewParams{
		Description: "description",
		ImageURL:    "image_url",
		Name:        "name",
		Price:       0,
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", product.ProductID)
}
