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

func TestProductNew(t *testing.T) {
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
	_, err := client.Products.New(context.TODO(), dackermanstore.ProductNewParams{
		Description: dackermanstore.F("description"),
		ImageURL:    dackermanstore.F("image_url"),
		Name:        dackermanstore.F("name"),
		Price:       dackermanstore.F(int64(0)),
	})
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductGet(t *testing.T) {
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
	_, err := client.Products.Get(context.TODO(), "product_id")
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductUpdate(t *testing.T) {
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
	_, err := client.Products.Update(
		context.TODO(),
		"product_id",
		dackermanstore.ProductUpdateParams{
			Description: dackermanstore.F("description"),
			ImageURL:    dackermanstore.F("image_url"),
			Name:        dackermanstore.F("name"),
			Price:       dackermanstore.F(int64(0)),
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

func TestProductList(t *testing.T) {
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
	_, err := client.Products.List(context.TODO())
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProductDelete(t *testing.T) {
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
	_, err := client.Products.Delete(context.TODO(), "product_id")
	if err != nil {
		var apierr *dackermanstore.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
