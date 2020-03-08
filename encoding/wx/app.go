package wx

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
)

var apiCode2Session = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

// 微信小程序 code2session 接口返回结构
type AppSession struct {
    OpenId     string `json:"openid"`
    UnionId    string `json:"unionid"`
    SessionKey string `json:"session_key"`
    ErrorCode  int    `json:"errcode"`
    ErrorMsg   string `json:"errmsg"`
}

// const (
//     APP_GENDER_UNKNOW = iota // 未知
//     APP_GENDER_MALE          // 男
//     APP_GENDER_FEMALE        // 女
// )

// 微信小程序 jssdk 的 getUserInfo 接口返回结构
type AppUserInfo struct {
    OpenId    string                `json:"openId"`
    UnionId   string                `json:"unionId"`
    NickName  string                `json:"nickName"`
    Gender    int                   `json:"gender"`
    AvatarUrl string                `json:"avatarUrl"`
    Province  string                `json:"province"`
    City      string                `json:"city"`
    Country   string                `json:"country"`
    Language  string                `json:"language"`
    Watermark *appUserInfoWatermark `json:"watermark"`
}
type appUserInfoWatermark struct {
    AppId     string `json:"appid"`
    Timestamp int    `json:"timestamp"`
}

// 微信小程序
type App struct {
    appId      string
    secret     string
    sessionKey string
    client     *http.Client
}

func NewApp(appid, secret, session string) *App {
    return &App{appId: appid, secret: secret, sessionKey: session}
}

func (w *App) clientFallback() *http.Client {
    if w.client == nil {
        return http.DefaultClient
    }
    return w.client
}

func (w *App) SetClient(c *http.Client) {
    w.client = c
}

// 登录凭证校验 (auth.code2Session)
func (w *App) Code2Session(code string) (*AppSession, error) {
    r, err := w.clientFallback().Get(fmt.Sprintf(apiCode2Session, w.appId, w.secret, code))
    if err != nil {
        if r != nil {
            r.Body.Close()
        }
        return nil, err
    }

    defer r.Body.Close()

    if r.StatusCode != http.StatusOK {
        return nil, errors.New(fmt.Sprintf("status code invalid: %d", r.StatusCode))
    }

    res, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil, err
    }

    var d AppSession
    err = json.Unmarshal(res, &d)
    if err != nil {
        return nil, err
    }

    if d.ErrorCode > 0 {
        return nil, errors.New(d.ErrorMsg)
    }

    // 自动更新实例
    w.sessionKey = d.SessionKey

    return &d, nil
}

// 解密 wx.getUserInfo 返回的 encryptedData 参数
func (w *App) DecryptBizData(encryptedData, iv string) (*AppUserInfo, error) {
    if len(w.sessionKey) != 24 {
        return nil, errors.New("sessionKey invalid")
    }

    aesKey, err := base64.StdEncoding.DecodeString(w.sessionKey)
    if err != nil {
        return nil, err
    }

    if len(iv) != 24 {
        return nil, errors.New("iv invalid")
    }

    aesIv, err := base64.StdEncoding.DecodeString(iv)
    if err != nil {
        return nil, err
    }

    cipherText, err := base64.StdEncoding.DecodeString(encryptedData)
    if err != nil {
        return nil, err
    }

    block, err := aes.NewCipher(aesKey)
    if err != nil {
        return nil, err
    }

    text := make([]byte, len(cipherText))

    mode := cipher.NewCBCDecrypter(block, aesIv)
    mode.CryptBlocks(text, cipherText)
    text = pKCS7UnPadding(text)

    var d AppUserInfo
    err = json.Unmarshal(text, &d)
    if err != nil {
        return nil, err
    }

    if d.Watermark.AppId != w.appId {
        return nil, errors.New("appid not match")
    }

    return &d, nil
}

func pKCS7UnPadding(text []byte) []byte {
    s := len(text)
    if s > 0 {
        return text[:(s - int(text[s-1]))]
    }
    return text
}
