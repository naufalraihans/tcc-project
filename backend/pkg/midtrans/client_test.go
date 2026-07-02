package midtrans

import (
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestVerifySignature(t *testing.T) {
	c := New("serverkey-123", false)
	order, status, gross := "TCC-AB12CD-1751234567", "200", "500000.00"

	sum := sha512.Sum512([]byte(order + status + gross + "serverkey-123"))
	valid := hex.EncodeToString(sum[:])

	if !c.VerifySignature(order, status, gross, valid) {
		t.Fatal("valid signature rejected")
	}
	if c.VerifySignature(order, status, gross, "0000") {
		t.Fatal("forged signature accepted")
	}
	if c.VerifySignature(order, "201", gross, valid) {
		t.Fatal("signature accepted with tampered status_code")
	}
}
