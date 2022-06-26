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
	id, errr := iloclient.NewILOClient().GetID()
	ssn, er := iloclient.NewILOClient().GetSSN()
	// var c []iloclient.UUIDStruct
	if err != nil || errr != nil || er != nil {
		panic(err)
	}
	fmt.Println("ID = ", id)
	fmt.Println("UUID = ", uid)
	fmt.Println("Valid_us_SSN = ", ssn)

}
