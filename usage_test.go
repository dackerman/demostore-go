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
	)
	product, err := client.Products.New(context.TODO(), dackermanstore.ProductNewParams{
		Description: dackermanstore.F("description"),
		ImageURL:    dackermanstore.F("image_url"),
		Name:        dackermanstore.F("name"),
		Price:       dackermanstore.F(int64(0)),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", product.ProductID)
}
