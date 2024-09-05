package main

import "log"

func main() {
	walletFacade := newWalletFacade("abc", 1234)

	err := walletFacade.addMoneyToWallet("abc", 1234, 10)

	if err != nil {
		log.Fatalf("Error : %s\n", err)
	}

	err = walletFacade.deductMoneyFromWallet("abc", 1234, 5)

	if err != nil {
		log.Fatalf("Error : %s\n", err)
	}
}
