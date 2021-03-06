package exec

import (
	"fmt"
	"testing"
	"github.com/bottos-project/bottos/contract"
	"github.com/bottos-project/bottos/common/types"
	"github.com/bottos-project/bottos/contract/msgpack"
)


func TestCallSubTrx(t *testing.T) {
	type transferparam struct {
		To			string
		Amount		uint32
	}

	param := transferparam {
		To     : "Clinton",
		Amount : 1233,
	}

	bf , err :=  msgpack.Marshal(param)
	fmt.Println(" TestCallSubTrx bf = ",bf," , err = ",err)

	trx := &types.Transaction{
		Version:1,
		CursorNum:1,
		CursorLabel:1,
		Lifetime:1,
		Sender:"bottos",
		Contract: "usermng",
		Method:  "r",
		Param: bf,
		SigAlg:1,
		Signature:[]byte{},
	}

	ctx := &contract.Context{ Trx:trx}


	res , err := GetInstance().Start(ctx, 1, false)
	if err != nil {
		fmt.Println("*ERROR* fail to execute start !!! ",err.Error())
		return
	}

	//check sub trx
	var tf transferparam
	for _ , sub_trx := range res {
		msgpack.Unmarshal(sub_trx.Param , &tf)
		fmt.Println("TestCallSubTrx sub_trx = ",sub_trx.Param ," , tf = ",tf)
	}
}