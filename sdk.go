package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var (
	api = &CoreApi{
		opt: &Option{},
	}
)

type CoreApi struct {
	opt *Option

	*userCli
	*bucketCli
	*marketCli
	*ipfsCli
}

type HttpResponse struct {
	Status  int    `json:"status"`  //http状态码
	Message string `json:"message"` //错误消息
	Value   string `json:"value"`   //json格式的响应体
}

func Init(opts ...Options) (*CoreApi, error) {
	for _, fn := range opts {
		fn(api.opt)
	}
	if len(api.opt.AppID) == 0 {
		return nil, errors.New("app-id invalid")
	}
	if len(api.opt.Name) == 0 || len(api.opt.Pwd) == 0 {
		return nil, errors.New("user invalid")
	}
	if len(api.opt.Addr) == 0 || !strings.HasPrefix(api.opt.Addr, "http") {
		return nil, errors.New("addr invalid")
	}

	api.userCli = newUserCli()
	api.bucketCli = newBucketCli()
	api.marketCli = newMarketCli()
	api.ipfsCli = newIpfsCli()
	if err := api.Login(); err != nil {
		return nil, err
	}
	return api, nil
}

func Do(path string, rq proto.Message, rs proto.Message) error {
	var (
		err  error
		req  *http.Request
		rsp  *http.Response
		data []byte
	)
	if req, err = http.NewRequest("POST", api.opt.Addr+path, strings.NewReader(ToJson(rq))); err != nil {
		return err
	} else {
		req.Header.Add("ContentType", "application/json")
		req.Header.Add("Authorization", "Idp "+api.Token())
	}
	if rsp, err = http.DefaultClient.Do(req); err != nil {
		return err
	} else if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code: %v", rsp.Status)
	}

	defer rsp.Body.Close()
	if data, err = ioutil.ReadAll(rsp.Body); err != nil {
		return err
	}

	m := &HttpResponse{}
	if err = json.Unmarshal(data, m); err != nil {
		return err
	} else if m.Status == http.StatusOK {
		if err = FromJson([]byte(m.Value), rs); err != nil {
			return err
		}
	} else {
		return errors.New(m.Message)
	}

	return nil
}

func ToJson(val interface{}) string {
	data, _ := json.MarshalIndent(val, "", "   ")
	return string(data)
}

func FromJson(data []byte, val proto.Message) error {
	m := &jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	return m.Unmarshal(bytes.NewBuffer(data), val)
}
