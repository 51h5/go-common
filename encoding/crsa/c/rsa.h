#ifndef _RSA_H_
#define _RSA_H_

#define XRSA_KEY_LEN (1024 / 8)
#define XRSA2_KEY_LEN (2048 / 8)

#ifdef __cplusplus
extern "C" {
#endif

void *LoadPrivateKey(char *file);
void *LoadPublicKey(char *file);
void DisposeKey(void *_p_rsa);

const char *RsaSign(void *_p_rsa, char *cstr);
int RsaVerify(void *_p_rsa, char *cstr, char *sign);

const char *Rsa2Sign(void *_p_rsa, char *cstr);
int Rsa2Verify(void *_p_rsa, char *cstr, char *sign);

#ifdef __cplusplus
}
#endif

#endif