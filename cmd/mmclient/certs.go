//go:build !windows

package main

import (
	"crypto/x509"
)

func SystemRoots() (*x509.CertPool, error) {
	return x509.SystemCertPool()
}
