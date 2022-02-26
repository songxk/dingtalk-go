package dingtalk

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type client struct {
	AccessToken string
	Secret      string
}

type response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

const httpTimeout = time.Duration(500) * time.Microsecond

func NewClient(accessToken string, secret string) *client {
	return &client{
		AccessToken: accessToken,
		Secret:      secret,
	}
}

//依赖robot文件中的robotMap
func NewRobotClient(robotId int) *client {
	if robot, ok := robotMap[robotId]; !ok {
		return nil
	} else {
		return &client{
			AccessToken: robot.AccessToken,
			Secret:      robot.Secret,
		}
	}
}

func (c *client) Send(message Message) (*response, error) {
	buf, err := message.ToByte()
	if err != nil {
		return nil, err
	}
	reqUrl, err := c.getUrl()

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")
	client := new(http.Client)
	client.Timeout = httpTimeout
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &response{}
	err = json.Unmarshal(res, ret)
	if err != nil {
		return nil, err
	}
	if ret.ErrCode > 0 {
		return nil, errors.New(fmt.Sprintf("%s(%d)", ret.ErrMsg, ret.ErrCode))
	}
	return ret, nil
}

func (c *client) getUrl() (string, error) {
	dingTalkUrl := url.URL{
		Scheme: "https",
		Host:   "oapi.dingtalk.com",
		Path:   "robot/send",
	}
	values := url.Values{}
	values.Set("access_token", c.AccessToken)
	if c.Secret == "" {
		return dingTalkUrl.String(), nil
	}
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	base := fmt.Sprintf("%s\n%s", timestamp, c.Secret)
	h := hmac.New(sha256.New, []byte(c.Secret))
	h.Write([]byte(base))
	base64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))
	sign := url.QueryEscape(base64Str)
	values.Set("timestamp", timestamp)
	values.Set("sign", sign)
	dingTalkUrl.RawQuery = values.Encode()
	return dingTalkUrl.String(), nil
}
