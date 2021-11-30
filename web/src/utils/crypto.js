import Crypto from 'crypto-js'
// 加密
export default {
  // 设置key和偏移量
  key: Crypto.enc.Latin1.parse('key'),
  iv: Crypto.enc.Latin1.parse('iv'),

  // 加密
  encrypto (password) {
    const res = Crypto.AES.encrypt(password, this.key, {
      iv: this.iv,
      mode: Crypto.mode.CBC,
      padding: Crypto.pad.ZeroPadding
    }).toString();
    return res;
  },

  // 解密
  decrypto (password) {
    const res = Crypto.AES.decrypt(password, this.key, {
      iv: this.iv,
      mode: Crypto.mode.CBC,
      padding: Crypto.pad.Pkcs7
    })
    return Crypto.enc.Utf8.stringify(res)
  }
}
