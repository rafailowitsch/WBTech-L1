package task

import "fmt"

// неподходящий интерфейс
type bitcoinWallet struct{}

// метод sendBitcoin, который используется в неподходящем интерфейсе
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

// метод send адаптера, который вызывает метод sendBitcoin из неподходящего интерфейса
func (ba *bitcoinAdapter) send(amount float64, address string) {
	ba.bitcoinWallet.sendBitcoin(amount, address)
}

// функция taskN21 демонстрирует использование паттерна адаптер
func TaskN21() {
	var wallet cryptoWallet

	btcWallet := &bitcoinWallet{}
	btcAdapter := &bitcoinAdapter{btcWallet}

	wallet = btcAdapter

	wallet.send(0.123123, "1JXVJ435dhE4dk34")
}
