# Interfaces in Go

An interface describes behavior as a set of method signatures. A value satisfies an interface when its type provides every required method.

Imagine a checkout that accepts different ways to pay:

```go
type PaymentMethod interface {
	Pay(amount float64) string
}
```

This interface accepts any value that has a `Pay(float64) string` method.

## Interfaces are satisfied implicitly

Go does not require an `implements` declaration:

```go
type CardPayment struct {
	LastFour string
}

func (card CardPayment) Pay(amount float64) string {
	return fmt.Sprintf(
		"Paid $%.2f with card ending %s",
		amount,
		card.LastFour,
	)
}

var method PaymentMethod = CardPayment{LastFour: "4242"}
fmt.Println(method.Pay(49.99))
```

Output:

```text
Paid $49.99 with card ending 4242
```

`CardPayment` satisfies `PaymentMethod` simply because its method set contains the required method. This keeps the concrete type independent from the interface.

An optional compile-time check can document that relationship:

```go
var _ PaymentMethod = CardPayment{}
```

The blank identifier discards the value, but the assignment fails to compile if `CardPayment` stops satisfying `PaymentMethod`.

## Accepting interface values

An interface parameter lets one function work with different concrete types that share behavior:

```go
type CashPayment struct{}

func (CashPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f with cash", amount)
}

func Checkout(method PaymentMethod, amount float64) string {
	return "Checkout: " + method.Pay(amount)
}

fmt.Println(Checkout(CardPayment{LastFour: "4242"}, 49.99))
fmt.Println(Checkout(CashPayment{}, 20))
```

Output:

```text
Checkout: Paid $49.99 with card ending 4242
Checkout: Paid $20.00 with cash
```

The function depends on the small `PaymentMethod` behavior instead of depending directly on card or cash payments.

## Interface values contain a type and a value

An interface value stores:

- the concrete type assigned to it;
- the concrete value of that type.

```go
var method PaymentMethod = CardPayment{LastFour: "4242"}

fmt.Printf("type = %T\n", method)
fmt.Printf("value = %v\n", method)
```

Output:

```text
type = main.CardPayment
value = {4242}
```

Calling `method.Pay(49.99)` uses the concrete type's implementation.

## Type assertions

A type assertion retrieves a particular concrete value from an interface:

```go
card, ok := method.(CardPayment)
if ok {
	fmt.Println(card.LastFour)
}
```

The second value reports whether the interface currently holds that type. This comma-ok form is safe when the concrete type is uncertain.

An assertion without the boolean panics when the type does not match:

```go
card := method.(CardPayment)
```

Use this shorter form only when the program can guarantee the concrete type.

## Type switches

A type switch handles several possible concrete types:

```go
func Describe(method PaymentMethod) string {
	switch concrete := method.(type) {
	case CardPayment:
		return "card ending " + concrete.LastFour
	case CashPayment:
		return "cash"
	default:
		return "unknown payment method"
	}
}
```

The variable declared in the switch has the matching concrete type inside each case.

## Value and pointer method sets

A method with a value receiver belongs to the method sets of both `T` and `*T`:

```go
func (card CardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f", amount)
}

var first PaymentMethod = CardPayment{LastFour: "4242"}
var second PaymentMethod = &CardPayment{LastFour: "1234"}
```

A method with a pointer receiver belongs to the method set of `*T`, not `T`. In that case, only the pointer satisfies the interface.

## Nil interface values

The zero value of an interface is `nil`:

```go
var method PaymentMethod
fmt.Println(method == nil)
```

Output:

```text
true
```

An interface holding a typed nil pointer is different:

```go
var missingCard *CardPayment
var method PaymentMethod = missingCard

fmt.Println(method == nil)
```

Output:

```text
false
```

The interface is not nil because it contains the concrete type `*CardPayment`, even though its concrete pointer value is nil.

## The `any` interface

`any` is an alias for `interface{}`, an interface with no required methods:

```go
var value any = 42
value = "Go"
```

Every type satisfies `any`. It is useful when values genuinely may have unrelated types, but a focused interface usually communicates intent more clearly.

## Summary

Interfaces let code depend on required behavior while concrete types remain free to satisfy that behavior implicitly.

- Define interfaces from small method sets.
- Accept interface parameters when several concrete types share behavior.
- Use comma-ok type assertions when the concrete type is uncertain.
- Remember that an interface holding a typed nil pointer is not itself nil.
