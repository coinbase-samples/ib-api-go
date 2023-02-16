/**
 * Copyright 2022-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerPassHealth(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	handler := aw.MakeHttpHandler()(nextHandler)

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	sc := w.Result().StatusCode
	if sc != 200 {
		t.Error("should return a 200 status code for health endpoint")
	}
}

func TestHandlerFailUnauthed(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	handler := aw.MakeHttpHandler()(nextHandler)

	req := httptest.NewRequest("GET", "/profile", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	sc := w.Result().StatusCode
	if sc == 200 {
		t.Error("should NOT return a 200 status code for any endpoint")
	}
}

func TestHandlerFailMissingBearer(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	handler := aw.MakeHttpHandler()(nextHandler)

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Add("Authorization", "")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	sc := w.Result().StatusCode
	if sc == 200 {
		t.Error("should NOT return a 200 status code for any endpoint")
	}
}

func TestHandlerFailInvalidBearer(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	handler := aw.MakeHttpHandler()(nextHandler)

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Add("Authorization", "bearer badToken")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	sc := w.Result().StatusCode
	if sc == 200 {
		t.Error("should NOT return a 200 status code for invalid bearer")
	}
}

func TestHandlerSucceed(t *testing.T) {
	cip := MockCognito{}
	aw := Middleware{Cip: &cip}

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	handler := aw.MakeHttpHandler()(nextHandler)

	req := httptest.NewRequest("GET", "/profile", nil)
	req.Header.Add("Authorization", "bearer goodToken")
	req = req.WithContext(context.Background())
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	sc := w.Result().StatusCode
	err := w.Result().Body
	if sc != 200 {
		t.Error("should return a 200 status code for any endpoint with valid bearer")
	}
	if err == nil {
		t.Error("expected a return body")
	}
}
