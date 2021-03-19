package test

import (
	"fmt"
	"gin/library/token"
	"testing"
)

var testTokenEncode = token.New()
var testTokenString = ""
var testTokenDecode = token.New()
var secret = "token"
var data = map[string]interface{}{
	"appKey":  "testAppKey",
	"channel": 2,
	"data1":   1,
	"data2":   "data2",
}

func TestToken_Encode(t *testing.T) {
	testTokenEncode.SetData(data)
	testTokenEncode.SetSecret(secret)
	tokenString, err := testTokenEncode.Encode()
	if err != nil {
		t.Fatal("token生成失败", err)
	}
	testTokenString = tokenString
	fmt.Println("token生成成功\n", tokenString)
}

func TestToken_Decode(t *testing.T) {
	testTokenDecode.SetSecret(secret)
	testTokenDecode.SetToken(testTokenString)

	testTokenData, err := testTokenDecode.Decode()
	if err != nil {
		t.Fatal("token解析失败", err)
	}
	fmt.Println("token解析成功\n", testTokenData)
}

func TestToken_DecodeSegment(t *testing.T) {
	testTokenData, err := testTokenDecode.DecodeSegment()
	if err != nil {
		t.Fatal("token直接解析失败", err)
	}
	fmt.Println("token直接解析成功\n", testTokenData)
}
