package main
import "fmt"

type bank struct{
    acc_no int
    bal float32
    name,acctype string
}



func (b *bank) deposit() float32{
	fmt.Println("enter money to be deposited")
    var x int
    fmt.Scan(&x)
	
	b.bal += float32(x)
  	return b.bal
}
func (b *bank) withdraw() float32 {
	fmt.Println("enter money to be withdrawn")
	var y int
	fmt.Scan(&y)
	
	if float32(y) > b.bal {
		fmt.Println("not enough balance")
	}else{
		b.bal -= float32(y)
	}
	return b.bal
}
func (b *bank) check() float32 {
	return b.bal
}

func main(){
	b := bank { acc_no: 50101, bal: 40000, name: "Maya", acctype:"salary"}
	
	fmt.Println("Account no: ", b.acc_no)
	fmt.Println("Name: ", b.name)
	fmt.Println("Account type: ", b.acctype)
	fmt.Println("Balance: ", b.bal)

	
		fmt.Println("deposited: ", b.deposit())
	
		fmt.Println("remaining balance: ", b.withdraw())
		
		fmt.Println("Available Balance: ", b.check())
		
	}
	
