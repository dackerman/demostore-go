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

func TestProductVariantNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Variants.New(context.TODO(), dackermanstore.ProductVariantNewParams{
		PathProductID: dackermanstore.F("product_id"),
		AddlPrice:     dackermanstore.F(0.000000),
		ImageURL:      dackermanstore.F("image_url"),
		Name:          dackermanstore.F("name"),
		BodyProductID: dackermanstore.F("product_id"),
		ID:            dackermanstore.F(int64(0)),
	})
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
	)
	_, err := client.Products.Variants.Get(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductVariantUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Products.Variants.Update(
		context.TODO(),
		int64(0),
		dackermanstore.ProductVariantUpdateParams{
			PathProductID: dackermanstore.F(int64(0)),
			ID: dackermanstore.F[dackermanstore.ProductVariantUpdateParamsIDUnion](dackermanstore.ProductVariantUpdateParamsIDIntSetInput{
				Set: dackermanstore.F(int64(0)),
			}),
			AddlPrice: dackermanstore.F[dackermanstore.ProductVariantUpdateParamsAddlPriceUnion](dackermanstore.ProductVariantUpdateParamsAddlPriceFloatSetInput{
				Set: dackermanstore.F(0.000000),
			}),
			ImageURL:      dackermanstore.F("image_url"),
			Name:          dackermanstore.F("name"),
			BodyProductID: dackermanstore.F("product_id"),
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
	)
	_, err := client.Products.Variants.List(context.TODO(), int64(0))
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
	)
	_, err := client.Products.Variants.Delete(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
