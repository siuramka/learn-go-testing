package pointers_errors

import "testing"

func TestWallet_Withdraw(t *testing.T) {
	t.Run("withdraws from wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		AssertNoError(t, err)

		want := Bitcoin(10)
		got := wallet.Balance()
		AssertBalance(t, want, got)
	})

	t.Run("withdraw returns nil error when not overdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("withdraw returns overdraw error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		if err.Error() != overdrawErrorMessage {
			t.Errorf("expected '%s' error, got %s", overdrawErrorMessage, err)
		}
	})
}

func TestWallet_Deposit(t *testing.T) {
	t.Run("deposits bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		got := wallet.Balance()
		AssertBalance(t, want, got)
	})
}

func AssertBalance(t testing.TB, want Bitcoin, got Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
