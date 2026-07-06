package practice

import "fmt"

type paymentMethod interface {
	Pay(amount float64) string
}

type cardPayment struct {
	LastFour string
}

func (card cardPayment) Pay(amount float64) string {
	return fmt.Sprintf(
		"Paid $%.2f with card ending %s",
		amount,
		card.LastFour,
	)
}

type cashPayment struct{}

func (cashPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f with cash", amount)
}

func checkout(method paymentMethod, amount float64) string {
	return "Checkout: " + method.Pay(amount)
}

func InterfacesEx15() {
	fmt.Println("1. Satisfy an interface implicitly")
	fmt.Println("A type satisfies an interface by providing all its methods.")
	fmt.Println("---")
	fmt.Println("type paymentMethod interface {")
	fmt.Println("    Pay(amount float64) string")
	fmt.Println("}")
	fmt.Println("func (card cardPayment) Pay(amount float64) string {")
	fmt.Printf("%s\n", `    return fmt.Sprintf("Paid $%.2f with card ending %s", amount, card.LastFour)`)
	fmt.Println("}")
	fmt.Println(`var current paymentMethod = cardPayment{LastFour: "4242"}`)
	var current paymentMethod = cardPayment{LastFour: "4242"}
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> %s\n", current.Pay(49.99))

	fmt.Println("\n2. Accept different concrete types")
	fmt.Println("An interface parameter depends on behavior instead of one concrete type.")
	fmt.Println("---")
	fmt.Println("func checkout(method paymentMethod, amount float64) string {")
	fmt.Println(`    return "Checkout: " + method.Pay(amount)`)
	fmt.Println("}")
	fmt.Println(`card := cardPayment{LastFour: "4242"}`)
	fmt.Println("cash := cashPayment{}")
	fmt.Println("checkout(card, 49.99)")
	fmt.Println("checkout(cash, 20)")
	card := cardPayment{LastFour: "4242"}
	cash := cashPayment{}
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> %s\n", checkout(card, 49.99))
	fmt.Printf(">> %s\n", checkout(cash, 20))

	fmt.Println("\n3. Use safe type assertions")
	fmt.Println("The comma-ok form reports whether the concrete type matches.")
	fmt.Println("---")
	fmt.Println(`current = cardPayment{LastFour: "1234"}`)
	fmt.Println("concreteCard, cardOK := current.(cardPayment)")
	fmt.Println("_, cashOK := current.(cashPayment)")
	current = cardPayment{LastFour: "1234"}
	concreteCard, cardOK := current.(cardPayment)
	_, cashOK := current.(cashPayment)
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> card ending = %q, ok = %t\n", concreteCard.LastFour, cardOK)
	fmt.Printf(">> cash assertion ok = %t\n", cashOK)

	fmt.Println("\n4. Compare nil interface values")
	fmt.Println("An interface holding a typed nil pointer is not a nil interface.")
	fmt.Println("---")
	fmt.Println("var empty paymentMethod")
	fmt.Println("var missingCard *cardPayment")
	fmt.Println("var holdingNil paymentMethod = missingCard")
	var empty paymentMethod
	var missingCard *cardPayment
	var holdingNil paymentMethod = missingCard
	fmt.Println("---")
	fmt.Println("Output:")
	fmt.Printf(">> empty == nil is %t\n", empty == nil)
	fmt.Printf(">> holdingNil == nil is %t\n", holdingNil == nil)
}
