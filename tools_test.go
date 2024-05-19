package toolkit

import "testing"

func TestTools_RandomString(test *testing.T){
	var testTools Tools
	
	result := testTools.RandomString(10)
	if len(result) != 10{
		test.Error("wrong length random string returned")
	}
}