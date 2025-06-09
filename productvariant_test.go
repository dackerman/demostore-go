// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dackermanstore_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/dackerman/demostore-go/v2"
	"github.com/dackerman/demostore-go/v2/internal/testutil"
	"github.com/dackerman/demostore-go/v2/option"
)

func TestProductVariantNew(t *testing.T) {
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
	_, err := client.Products.Variants.New(
		context.TODO(),
		"product_id",
		dackermanstore.ProductVariantNewParams{
			ImageURL: "image_url",
			Name:     "name",
			Price:    0,
		},
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductVariantGet(t *testing.T) {
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
	_, err := client.Products.Variants.Get(
		context.TODO(),
		"product_id",
		"variant_id",
		dackermanstore.ProductVariantGetParams{},
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductVariantUpdate(t *testing.T) {
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
	_, err := client.Products.Variants.Update(
		context.TODO(),
		"product_id",
		"variant_id",
		dackermanstore.ProductVariantUpdateParams{
			ImageURL: "image_url",
			Name:     "name",
			Price:    0,
		},
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductVariantList(t *testing.T) {
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
	_, err := client.Products.Variants.List(
		context.TODO(),
		"product_id",
		dackermanstore.ProductVariantListParams{},
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductVariantDelete(t *testing.T) {
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
	_, err := client.Products.Variants.Delete(
		context.TODO(),
		"product_id",
		"variant_id",
		dackermanstore.ProductVariantDeleteParams{},
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
