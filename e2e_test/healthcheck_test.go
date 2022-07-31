package e2e_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestHappyHealthCheck() {
	c := http.Client{}

	r, _ := c.Get("http://localhost:8080/healthcheck")

	s.Equal(http.StatusOK, r.StatusCode)

	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status": "OK", "messages": []}`, string(b))
}
