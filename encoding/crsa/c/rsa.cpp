#include "base64.h"
#include "rsa.h"

#include <string>
#include <string.h>
#include <openssl/pem.h>
#include <openssl/rsa.h>

// base64编码
const char *base64Encode(const char *in_str, int in_len) {
    if (in_str == NULL) {
        return NULL;
    }

    BIO *b64 = NULL;
	BIO *bm = NULL;
	BUF_MEM *bptr = NULL;

	b64 = BIO_new(BIO_f_base64());
	BIO_set_flags(b64, BIO_FLAGS_BASE64_NO_NL);

	bm = BIO_new(BIO_s_mem());
	bm = BIO_push(b64, bm);

	BIO_write(bm, in_str, in_len);
	BIO_flush(bm);

	BIO_get_mem_ptr(bm, &bptr);
	char *buff = (char *)malloc(bptr->length + 1);
	memcpy(buff, bptr->data, bptr->length);
	buff[bptr->length] = '\0';

	BIO_free_all(bm);

	return buff;
}

// base64解码
int base64Decode(char *in_str, int in_len, char *out_str) {
    if (in_str == NULL || out_str == NULL) {
      return -1;
    }

	BIO * b64 = NULL;
	BIO * bm = NULL;
	int size = 0;

	b64 = BIO_new(BIO_f_base64());
	BIO_set_flags(b64, BIO_FLAGS_BASE64_NO_NL);

	bm = BIO_new_mem_buf(in_str, in_len);
	bm = BIO_push(b64, bm);

    size = BIO_read(bm, out_str, in_len);
	BIO_free_all(bm);

	return size;
}

// RSA 签名
const char *RsaSign(void *_p_rsa, char *cstr) {
  if (cstr == NULL || strlen(cstr) == 0) {
    return NULL;
  }

  RSA *p_rsa = (RSA *)_p_rsa;
  if (p_rsa == NULL) {
    return NULL;
  }

  unsigned char hash[SHA_DIGEST_LENGTH] = {0};
  SHA1((unsigned char *)cstr, strlen(cstr), hash);

  unsigned char sign[XRSA_KEY_LEN] = {0};
  unsigned int sign_len = sizeof(sign);

  int r = RSA_sign(NID_sha1, hash, SHA_DIGEST_LENGTH, sign, &sign_len, p_rsa);
  if (0 != r && sizeof(sign) == sign_len) {
    return base64Encode((const char *)sign, sign_len);
  }

  return NULL;
}

// RSA 验签
int RsaVerify(void *_p_rsa, char *cstr, char *sign) {
  if (cstr == NULL || sign == NULL) {
    return -1;
  }

  RSA *p_rsa = (RSA *)_p_rsa;
  if (p_rsa == NULL) {
    return -2;
  }

  char sign_cstr[XRSA_KEY_LEN];

  // 解码签名
  int ss = base64Decode(sign, strlen(sign), sign_cstr);
  if (ss <= 0) {
    return -3;
  }

  // 生成SHA1摘要
  unsigned char md[SHA_DIGEST_LENGTH] = {0};
  SHA1((unsigned char *)cstr, strlen(cstr), md);

  // 签名验签
  int r = RSA_verify(NID_sha1, md, SHA_DIGEST_LENGTH, (unsigned char *)sign_cstr, XRSA_KEY_LEN, p_rsa);

  return r > 0 ? 1 : -4;
}

// RSA2 签名
const char *Rsa2Sign(void *_p_rsa, char *cstr) {
  if (cstr == NULL || strlen(cstr) == 0) {
    return NULL;
  }

  RSA *p_rsa = (RSA *)_p_rsa;
  if (p_rsa == NULL) {
    return NULL;
  }

  unsigned char hash[SHA256_DIGEST_LENGTH] = {0};
  SHA256((unsigned char *)cstr, strlen(cstr), hash);

  unsigned char sign[XRSA2_KEY_LEN] = {0};
  unsigned int sign_len = sizeof(sign);

  int r = RSA_sign(NID_sha256, hash, SHA256_DIGEST_LENGTH, sign, &sign_len, p_rsa);
  if (0 != r && sizeof(sign) == sign_len) {
    return base64Encode((const char *)sign, sign_len);
  }

  return NULL;
}

// RSA2 验签
int Rsa2Verify(void *_p_rsa, char *cstr, char *sign) {
  if (cstr == NULL || sign == NULL) {
    return -1;
  }

  RSA *p_rsa = (RSA *)_p_rsa;
  if (p_rsa == NULL) {
    return -2;
  }

  unsigned char sign_cstr[XRSA2_KEY_LEN] = {0};

  // 解码签名
  int r1 = base64Decode(sign, strlen(sign), (char *)sign_cstr);
  if (r1 <= 0) {
    return -3;
  }

  // 生成SHA1摘要
  unsigned char md[SHA256_DIGEST_LENGTH] = {0};
  SHA256((unsigned char *)cstr, strlen(cstr), md);

  // 签名验签
  int r = RSA_verify(NID_sha256, md, SHA256_DIGEST_LENGTH, sign_cstr, XRSA2_KEY_LEN, p_rsa);

  return r > 0 ? 1 : -4;
}

void *LoadPrivateKey(char *file) {
  if (file == NULL || strlen(file) == 0) {
    return NULL;
  }

  FILE *f = fopen(file, "rb");
  if (f == NULL) {
    return NULL;
  }

  RSA *p_rsa = RSA_new();
  p_rsa = PEM_read_RSAPrivateKey(f, &p_rsa, NULL, NULL);
  fclose(f);

  return p_rsa == NULL ? NULL : (void *)p_rsa;
}

void *LoadPublicKey(char *file) {
  if (file == NULL || strlen(file) == 0) {
    return NULL;
  }

  FILE *f = NULL;
  f = fopen(file, "rb");
  if (f == NULL) {
    return NULL;
  }

  RSA *p_rsa = RSA_new();
  p_rsa = PEM_read_RSA_PUBKEY(f, &p_rsa, 0, 0);
  fclose(f);

  return p_rsa == NULL ? NULL : (void *)p_rsa;
}

void DisposeKey(void *_p_rsa) {
  if (_p_rsa != NULL) {
    RSA_free((RSA *)_p_rsa);
  }
}



//int main() {
//  // RSA 测试
//  char str[] = "app_id=xx&auth_token=xx&charset=UTF-8&code=xx&format=JSON&grant_type=xxxx&method=alipay.system.oauth.token&refresh_token=xxxx&sign_type=RSA&timestamp=2020-01-19 01:17:39&version=1.0";
//  void *_p_rsa = LoadPrivateKey((char*)"../rsa_private.pem");
//  RSA *p_rsa = (RSA *)_p_rsa;
//
//  const char *sign_str = RsaSign(p_rsa, str);
//  int r1 = RsaVerify(p_rsa, str, (char *)sign_str);
//  DisposeKey(p_rsa);
//
//  printf("RSA签名: %s\n", sign_str);
//  printf("RSA验签: %d\n", r1);
//
//
//
//  // RSA2 测试
//  char str2[] = "app_id=xx&auth_token=xx&charset=UTF-8&code=xx&format=JSON&grant_type=xxxx&method=alipay.system.oauth.token&refresh_token=xxxx&sign_type=RSA2&timestamp=2020-01-19 01:17:39&version=1.0";
//  void *_p_rsa2 = LoadPrivateKey((char*)"../rsa2_private.pem");
//  RSA *p_rsa2 = (RSA *)_p_rsa2;
//
//  const char *sign_str2 = Rsa2Sign(p_rsa2, str2);
//  int r2 = Rsa2Verify(p_rsa2, str2, (char*)sign_str2);
//  DisposeKey(p_rsa2);
//
//  printf("RSA2签名: %s\n", sign_str2);
//  printf("RSA2验签: %d\n", r2);
//}