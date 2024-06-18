package main

import "fmt"

// неподходящий интерфейс
type bitcoinWallet struct{}

func (bw *bitcoinWallet) sendBitcoin(amount float64, address string) {
	fmt.Printf("Sending %.2f BTC to address %s\n", amount, address)
}

// используемый интерфейс
type cryptoWallet interface {
	send(amount float64, address string)
}

// адаптер для неподходящего интерфейса
type bitcoinAdapter struct {
	bitcoinWallet *bitcoinWallet
}

func (ba *bitcoinAdapter) send(amount float64, address string) {
	ba.bitcoinWallet.sendBitcoin(amount, address)
}

func taskN21() {
	btcWallet := &bitcoinWallet{}
	btcAdapter := &bitcoinAdapter{btcWallet}

	btcAdapter.send(0.123123, "1JXVJ435dhE4dk34")
}
