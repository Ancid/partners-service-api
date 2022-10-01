package api

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartnersList(t *testing.T) {
	f := make(url.Values)
	httptest.NewRequest(http.MethodGet, "/partners/list", strings.NewReader(f.Encode()))
	res := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestPartnersSearchWithAllParams(t *testing.T) {
	f := make(url.Values)
	httptest.NewRequest(http.MethodGet, "/partners?latitude=49.671072&longitude=8.850669&material=wood", strings.NewReader(f.Encode()))
	res := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestAllPartners(t *testing.T) {
	f := make(url.Values)
	httptest.NewRequest(http.MethodGet, "/partners", strings.NewReader(f.Encode()))
	res := httptest.NewRecorder()

	assert.Equal(t, http.StatusOK, res.Code)
}
