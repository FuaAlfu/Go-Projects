package handlers

import(
	"testing"
	"string"
	"net/http"
	"io/ioutil"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct{
	suite.Suite
}

func TestEndToEndSuite(t *testing.T){
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler(){
	c := http.Client{}
	r,_  := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOk, r.StatusCode)
}

func (s *EndToEndSuite) TestPostNoResult(){
	c := http.Client{}
	r,_  := c.Get("http://localhost:1323/posts/229")
	s.Equal(http.StatusOk, r.StatusCode)
	b,_ := ioutill.ReadAll(r.body)
	s.JSONEq(`{"status": "ok","data":[]}`,string(b))
}