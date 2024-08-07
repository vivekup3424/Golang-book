package bank1

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() {
	return <-balances
}
