// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: proto.proto

package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/ptypes/empty"
)

type webGreeterHandler struct {
	r chi.Router
	h GreeterHandler
}

func (h *webGreeterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webGreeterHandler) Say(w http.ResponseWriter, r *http.Request) {

	req := &SayRequest{}

	resp := &SayResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.Say(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webGreeterHandler) SayAnything(w http.ResponseWriter, r *http.Request) {
	req := &empty.Empty{}

	resp := &SayResponse{}

	if err := h.h.SayAnything(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func RegisterGreeterWeb(r chi.Router, i GreeterHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webGreeterHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/say", handler.Say)
	r.MethodFunc("POST", "/api/anything", handler.SayAnything)
}

// SayRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SayRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SayRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SayRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SayRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SayRequest)(nil)

// SayRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SayRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SayRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SayRequest) UnmarshalJSON(b []byte) error {
	return SayRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SayRequest)(nil)

// SayResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SayResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SayResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SayResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SayResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SayResponse)(nil)

// SayResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SayResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SayResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SayResponse) UnmarshalJSON(b []byte) error {
	return SayResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SayResponse)(nil)
