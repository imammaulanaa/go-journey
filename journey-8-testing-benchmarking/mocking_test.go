package main

import "testing"

type MockNotifier struct {
    Called bool
    Msg    string
}

func (m *MockNotifier) Send(msg string) {
    m.Called = true
    m.Msg = msg
}

func TestKirimPeringatan(t *testing.T) {
    mock := &MockNotifier{}
    KirimPeringatan(mock)

    if !mock.Called {
        t.Error("Send tidak dipanggil")
    }
    if mock.Msg != "Peringatan: Sistem dalam beban tinggi!" {
        t.Errorf("Pesan tidak sesuai: %s", mock.Msg)
    }
}
