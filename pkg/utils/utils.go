/*
 * // Copyright 2020 Thales DIS CPL Inc
 * //
 * // Permission is hereby granted, free of charge, to any person obtaining
 * // a copy of this software and associated documentation files (the
 * // "Software"), to deal in the Software without restriction, including
 * // without limitation the rights to use, copy, modify, merge, publish,
 * // distribute, sublicense, and/or sell copies of the Software, and to
 * // permit persons to whom the Software is furnished to do so, subject to
 * // the following conditions:
 * //
 * // The above copyright notice and this permission notice shall be
 * // included in all copies or substantial portions of the Software.
 * //
 * // THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * // EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * // MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * // NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
 * // LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
 * // OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * // WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package utils

import (
	"errors"
	"github.com/ThalesIgnite/crypto11"
	"github.com/ThalesIgnite/gose/jose"
	"os"
)

var (
	ErrNoSuchKey  = errors.New("no such key")
	ErrNoSuchCert = errors.New("no such cert")

	AuthenticatedEncryptedKeyOperations = []jose.KeyOps{jose.KeyOpsDecrypt, jose.KeyOpsEncrypt}
	AsymmetricDecryptionKeyOperations   = []jose.KeyOps{jose.KeyOpsDecrypt, jose.KeyOpsEncrypt}

	SigningKeyOperations      = []jose.KeyOps{jose.KeyOpsSign}
	VerificationKeyOperations = []jose.KeyOps{jose.KeyOpsVerify}
)

func GetCrypto11Config() (config *crypto11.Config) {

	config = &crypto11.Config{
		Path:            os.Getenv("P11_LIB"),
		TokenLabel:      os.Getenv("P11_LABEL"),
		Pin:             os.Getenv("P11_PIN"),
		UseGCMIVFromHSM: true,
	}
	return
}
