package basics

import (
	"fmt"
)

// через интерфейсы реализован полиморфизм
// в go типизация неявная

type Payer interface {
	Pay(int) error
}

type Ringer interface {
	Ring(string) error
}

type NFCPhone interface {
	Payer  // = Pay(int) error
	Ringer // = Ring(string) error
}

type Stringer interface {
	String() string
}

//--------------------------------------------

type Phone struct {
	Money   int
	AppleID string
}

func (p *Phone) Pay(amount int) error {
	if p.Money < amount {
		return fmt.Errorf("недостаточно средств на счёте")
	}
	p.Money -= amount
	return nil
}

func (p *Phone) Ring(number string) error {
	if number == "" {
		return fmt.Errorf("введите номер")
	}
	return nil
}

//--------------------------------------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("нe хвататет денег")
	}
	w.Cash -= amount
	return nil
}

func (w *Wallet) String() string {
	return "кошелёк с деньгами"
}

//--------------------------------------------

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("недостаточно средств на карте")
	}
	c.Balance -= amount
	return nil
}

//--------------------------------------------

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("недостаточно средств на счёте")
	}
	a.Money -= amount
	return nil
}

//--------------------------------------------

// func Buy(p Payer) {
// 	err := p.Pay(10)
// 	if err != nil {
// 		fmt.Printf("Ошибка при оплате через %T\n", p)
// 		return
// 	}
// 	fmt.Printf("спасибо за покупку через %T\n", p)
// }

func Buy(p Payer) {
	switch p := p.(type) {
	case *Wallet:
		fmt.Println("оплата наличными")
		fmt.Printf("спасибо за покупку через %T\n", p)
	case *Card:
		fmt.Println("вставляйте карту")
		fmt.Printf("спасибо за покупку через %T\n", p)
	default:
		fmt.Println("что-то новое")
		fmt.Printf("Ошибка при оплате через %T\n", p)
	}
}

//--------------------------------------------

func BuyEmpty(in interface{}) {
	var p Payer
	var ok bool
	if p, ok = in.(Payer); !ok {
		fmt.Printf("%T не является платёжным средством\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("ошибка при оплате %T: %v\n", p, err)
		return
	}

	fmt.Printf("спасибо за покупку через %T\n", in)
}

//--------------------------------------------

func PayWithPhone(phone NFCPhone) {
	err := phone.Pay(1)
	if err != nil {
		fmt.Printf("ошибка при оплате %v\n", err)
		return
	}
	fmt.Printf("турникет открыт через %T\n", phone)
}

//--------------------------------------------

func InterfacesPrint() {
	MyWallet := Wallet{Cash: 100}
	Buy(&MyWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "enot"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)

	wal := &Wallet{Cash: 100}
	fmt.Printf("%#v\n", wal) // &basics.Wallet{Cash:100}
	fmt.Printf("%s\n", wal)  // кошелёк с деньгами

	// работа с пустым интерфейсом
	pw := &Wallet{Cash: 100}
	BuyEmpty(pw)

	pc := &Card{Balance: 9}
	BuyEmpty(pc)

	pa := &ApplePay{Money: 100}
	BuyEmpty(pa)

	// работа со встраиваемыми интерфейсами
	myPhone := &Phone{Money: 9}
	PayWithPhone(myPhone)
}
