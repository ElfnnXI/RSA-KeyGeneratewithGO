package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	// Ukuran kunci (4096-bit untuk keamanan tinggi)
	bitSize := 4096

	// Membuat private key RSA
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		fmt.Println("Gagal membuat private key:", err)
		return
	}

	// Simpan private key ke file
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		fmt.Println("Gagal membuat file private key:", err)
		return
	}
	defer privateKeyFile.Close()

	// Encode private key ke dalam format PEM
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)
	privateKeyFile.Write(privateKeyPEM)
	fmt.Println("Private key disimpan di 'private_key.pem'.")

	// Ekstrak public key dari private key
	publicKey := &privateKey.PublicKey

	// Simpan public key ke file
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		fmt.Println("Gagal membuat file public key:", err)
		return
	}
	defer publicKeyFile.Close()

	// Marshal dan encode public key ke PEM (PKIX format)
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Gagal encode public key:", err)
		return
	}

	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)
	publicKeyFile.Write(publicKeyPEM)
	fmt.Println("Public key disimpan di 'public_key.pem'.")
}
