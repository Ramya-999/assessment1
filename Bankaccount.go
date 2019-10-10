package main
import "fmt"

type bank struct{
    acc_no int
    bal float32
    name,acctype string
}

func  create() {
	b := bank{}

	fmt.Println("enter the name: ")
	fmt.Scan(&b.name)
	fmt.Println("enter the account type: ")
	fmt.Scan(&b.acctype)
	fmt.Println("enter the account no: ")
	fmt.Scan(&b.acc_no)
	fmt.Println("enter the balance: ")
	fmt.Scan(&b.bal)
	
        fmt.Println()

        fmt.Println("Account no: ", b.acc_no)
	fmt.Println("Name: ", b.name)
	fmt.Println("Account type: ", b.acctype)
	fmt.Println("Balance: ", b.bal)
	for{
	fmt.Println("To deposit money enter 1 -  To withdraw money enter 2 - To check the balance enter 3 - To exit enter 4")
	var c int
	fmt.Scan(&c)
	if c == 4{
		fmt.Println("Exited")
		break
	}
	switch c {
	case 1:
		fmt.Println("enter money to be deposited")
                var x int
                fmt.Scan(&x)
		fmt.Println("deposited: ", b.deposit(x))
	case 2:
		fmt.Println("enter money to be withdrawn")
	        var y int
	        fmt.Scan(&y)
		fmt.Println("remaining balance: ", b.withdraw(y))
	case 3:
		fmt.Println("Available Balance: ", b.check())	
	default:
		fmt.Println("enter a valid number")

	}
    }
	
}

func (b *bank) deposit(x int) float32{	
	b.bal += float32(x)
  	return b.bal
}
func (b *bank) withdraw(y int) float32 {
	
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
		create()	
	}
	
