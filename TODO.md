# Dev

## Error triggered with tpm2

Using aes gcm to libtpm2_pkcs11 triggers :

```sh
github.com/ThalesIgnite/crypto11.genericAead.Seal({0xc000013668?, 0x0?, 0xc0000d15b0?, 0xc0001f0ba0?}, {0xc00002c9a8, 0x0, 0x14}, {0xc00021c
0b0, 0x10, 0x10}, ...)           
/home/kaio/Projects/thales/k8s-kms-plugin/vendor/github.com/ThalesIgnite/crypto11/aead.go:159 +0x1e8
github.com/ThalesIgnite/gose.(*AesGcmCryptor).Seal(0xc00007b260, {0xa1608e, 0x7}, {0xc00021c0b0, 0x10, 0x10}, {0xc000015f78, 0x4, 0x8}, {0xc
000028780, ...})
/home/kaio/Projects/thales/k8s-kms-plugin/vendor/github.com/ThalesIgnite/gose/aes_gcm_cryptor.go:100 +0x331
github.com/ThalesIgnite/gose.(*JweDirectEncryptionEncryptorImpl).Encrypt(0xc000013680, {0xc000015f78, 0x4, 0x8}, {0x0, 0x0, 0x0})
/home/kaio/Projects/thales/k8s-kms-plugin/vendor/github.com/ThalesIgnite/gose/jwe_direct_encryptor.go:79 +0x30c
github.com/thalescpl-io/k8s-kms-plugin/pkg/providers.(*P11).Encrypt(0xc0000cb080, {0x45684b?, 0x10?}, 0xc0001bb5e0)
/home/kaio/Projects/thales/k8s-kms-plugin/pkg/providers/p11.go:427 +0x326
```

So :

```
(k8s-kms-plugin/p11.go) P11.Encrypt()
   -> (gose/jwe_direct_encryptor.go) gose.JweDirectEncryptionEncryptorImpl.Encrypt() 
       -> (gose/aes_gcm_cryptor.go) gose.AesGcmCryptor.SeaLl()
           -> (crypto11/aead.go) crypto11.genericAead.Seal()
```


```
27: C_EncryptInit              
2024-01-19 16:36:29.921       
[in] hSession = 0x100000000000006                    
[in] pMechanism->type = CKM_AES_CBC_PAD              
[in] pMechanism->pParameter[ulParameterLen] 00007f1420000eb0 / 16
    00000000  00 6F 99 8A 51 9F B5 B3 4D E6 85 45 06 DB 52 89  .o..Q...M..E..R.
[in] hKey = 0x1
Returned:  0 CKR_OK

28: C_EncryptUpdate
2024-01-19 16:36:29.926
[in] hSession = 0x100000000000006
[in] pPart[ulPartLen] 000000c000015f30 / 4
    00000000  70 69 6E 67                                      ping             
[out] pEncryptedPart[*pulEncryptedPartLen] NULL [size : 0x0 (0)]
Returned:  0 CKR_OK

29: C_EncryptUpdate                                                   
2024-01-19 16:36:29.926  
[in] hSession = 0x100000000000006  
[in] pPart[ulPartLen] 000000c000015f30 / 4                            
    00000000  70 69 6E 67                                      ping
[out] pEncryptedPart[*pulEncryptedPartLen] 00007f1420000e50 / 0      
    00000000
Returned:  0 CKR_OK 
```

Détecter si la chaine est <16 bytes.
Si oui :
- conserver la longueur initial dans un champ du header du jwe
- rajouter des '0' pour compléter jusqu'à 16 bytes (suffixe)
- chiffre la chaine de 16 bytes
Sinon chiffrer tel quel avec le padding.

au moment du déchiffrement, détecter dans le header si la chaine plaintext était <16bytes.
Si oui : 
- déchiffrer puis enlever les suffixes jusqu'à obtenir la chaine de taille indiquée dans le header
- retourner cette chaine de taille <16 bytes
Sinon retourner tel quel.