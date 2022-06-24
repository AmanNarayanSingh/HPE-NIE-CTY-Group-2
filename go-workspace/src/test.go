package main

import (
	"fmt"
	// "src/customtext"
	"src/iloclient"
)

func main() {
	fmt.Println("Hello")
	// fmt.Println("Who said tea", customtext.WhoSaidTea)
	fmt.Println(iloclient.NewILOClient().Health())
	uid, err := iloclient.NewILOClient().GetRandomUUID()
	// var c []iloclient.UUIDStruct
	if err != nil {
		panic(err)
	}

	fmt.Print(uid)

}
