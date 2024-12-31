package http

import (
	"context"
	"fmt"
	"testing"
)

func TestHttpGet(t *testing.T) {
	fmt.Println("TestHttpGet")
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	client := NewHTTPClient()

	resp, err := client.PostWithHeader(ctx,
		"https://api.shipos.cn/api/getLabels.php",
		[]byte(`{"orderid": "odr-0878cf33-0fad-4834-a99e-253569cdf606"}`),
		map[string]string{"appkey": "shipos_959", "appsecret": "194aa9287722361e0e9e22a950c2b98a"})
	if err != nil {
		t.Fatalf("failed to get response: %v", err)
	}

	fmt.Println(string(resp))

}
