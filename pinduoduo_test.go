package pinduoduo

import (
	"fmt"
	"github.com/dtapps/go-pinduoduo/type"
	"testing"
)

func TestSearch(t *testing.T) {
	duo := PinDuoDuo{
		ClientId:     "fd5177be0c354b5f87f8314798eb95dag",
		ClientSecret: "dcb7aa6fc9776e78b208dcbdd281d82cb8d8a0bdc",
	}
	param := Parameter{
		"keyword": "小米",
	}
	_, err := duo.Send(_type.GoodsSearch, param)
	if err != nil {
		t.Errorf("错误：%v", err)
		return
	}
}

func TestDetail(t *testing.T) {
	duo := PinDuoDuo{
		ClientId:     "fd5177be0c354b5f87f8314798eb95dag",
		ClientSecret: "dcb7aa6fc9776e78b208dcbdd281d82cb8d8a0bdc",
	}
	param := Parameter{
		"goods_sign":        "Y9z2lthmqx5NM-yBwfbZdh_ZcH08P8Pp_JXEw8DRMx",
		"pid":               "11270182_153701524",
		"custom_parameters": "",
	}
	send, err := duo.Send(_type.GoodsDetail, param)
	fmt.Printf("send：%v\n", send)
	if err != nil {
		t.Errorf("错误：%v", err)
		return
	}
}
