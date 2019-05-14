package message

import (
	"mime"

	"github.com/emersion/go-message/textproto"
)

func parseHeaderWithParams(s string) (f string, params map[string]string, err error) {
	f, params, err = mime.ParseMediaType(s)
	if err != nil {
		return s, nil, err
	}
	for k, v := range params {
		params[k], _ = decodeHeader(v)
	}
	return
}

func formatHeaderWithParams(f string, params map[string]string) string {
	encParams := make(map[string]string)
	for k, v := range params {
		encParams[k] = encodeHeader(v)
	}
	return mime.FormatMediaType(f, encParams)
}

// A Header represents the key-value pairs in a message header.
type Header struct {
	textproto.Header
}

// ContentType parses the Content-Type header field.
//
// If no Content-Type is specified, it returns "text/plain".
func (h *Header) ContentType() (t string, params map[string]string, err error) {
	v := h.Get("Content-Type")
	if v == "" {
		return "text/plain", nil, nil
	}
	return parseHeaderWithParams(v)
}

// SetContentType formats the Content-Type header field.
func (h *Header) SetContentType(t string, params map[string]string) {
	h.Set("Content-Type", formatHeaderWithParams(t, params))
}

// ContentDisposition parses the Content-Disposition header field, as defined in
// RFC 2183.
func (h *Header) ContentDisposition() (disp string, params map[string]string, err error) {
	return parseHeaderWithParams(h.Get("Content-Disposition"))
}

// SetContentDisposition formats the Content-Disposition header field, as
// defined in RFC 2183.
func (h *Header) SetContentDisposition(disp string, params map[string]string) {
	h.Set("Content-Disposition", formatHeaderWithParams(disp, params))
}

// Text parses a plaintext header field. The field charset is automatically
// decoded to UTF-8.
func (h *Header) Text(k string) (string, error) {
	return decodeHeader(h.Get(k))
}

// SetText sets a plaintext header field.
func (h *Header) SetText(k, v string) {
	h.Set(k, encodeHeader(v))
}
