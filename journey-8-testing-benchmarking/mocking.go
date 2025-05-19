package main

import "fmt"

type Notifier interface {
    Send(msg string)
}

type EmailNotifier struct{}

func (e EmailNotifier) Send(msg string) {
    fmt.Println("Mengirim email:", msg)
}

func KirimPeringatan(n Notifier) {
    n.Send("Peringatan: Sistem dalam beban tinggi!")
}
