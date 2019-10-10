package main

import (
	"testing"
)

func Testdeposit(t *testing.T){

	x:=bank{acc_no:50101, bal:40000, name:"employee", acctype:"salary"}
    
    if x.deposit(5000)!=45000{
        t.Error("Incorrect balance")
    }
}

func Testwithdraw(t *testing.T){
	
	x:=bank{acc_no:50102, bal:60000, name:"employee", acctype:"salary"}
    
    if x.withdraw(2000)!=58000{
        t.Error("Incorrect balance")
    }
}
