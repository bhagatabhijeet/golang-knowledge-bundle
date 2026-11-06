// Structs and methods: value vs pointer receivers.
//
// Run from assets/code:
//
//	go run ./ch03/structs
package main

import "fmt"

// Account is a plain struct: a named bundle of fields.
type Account struct {
	Owner   string
	Balance int // cents, to avoid float rounding
}

// String has a value receiver: it only reads, so a copy is fine.
func (a Account) String() string {
	return fmt.Sprintf("%s: $%d.%02d", a.Owner, a.Balance/100, a.Balance%100)
}

// Deposit has a pointer receiver: it must change the original.
func (a *Account) Deposit(cents int) {
	a.Balance += cents
}

// DepositOnCopy shows the classic bug: a value receiver mutates a copy.
func (a Account) DepositOnCopy(cents int) {
	a.Balance += cents // lost when the method returns
}

// NewAccount is the idiomatic "constructor": a plain function.
func NewAccount(owner string) *Account {
	return &Account{Owner: owner}
}

func main() {
	acct := NewAccount("Ada")
	acct.Deposit(1050)
	fmt.Println(acct) // fmt calls the String method

	acct.DepositOnCopy(99999)
	fmt.Println("after DepositOnCopy:", acct)

	// Addressable values get an automatic &: v.Deposit() means (&v).Deposit().
	v := Account{Owner: "Grace"}
	v.Deposit(250)
	fmt.Println(v)

	// Zero values and comparison.
	var zero Account
	fmt.Printf("zero value: Owner=%q Balance=%d\n", zero.Owner, zero.Balance)
	fmt.Println("equal?", Account{"Linus", 1} == Account{"Linus", 1})

	// Anonymous struct: handy for one-off shapes such as test tables.
	point := struct{ X, Y int }{3, 4}
	fmt.Println("point:", point.X, point.Y)
}
