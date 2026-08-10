package main

import "flag"

type PaymentMethod interface {
	Pay(amount float64) string
}

func main() {
	payment_method := flag.String("method", "", "With which method to pay")
	pay := flag.Float64("pay", 0, "Amount to Pay")

	flag.Parse()

	if *payment_method != "" {
		switch *payment_method {
		case "upi":
		}
	}
}
