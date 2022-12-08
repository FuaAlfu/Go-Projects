package handlers

import(
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct{
	suite.Suite
}

func TestEndToEndSuite(t *testing.T){
	suite.Rue(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler(){
	c := http.Client{}
}