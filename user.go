package sdk

import (
	"hlm-ipfs/sdk/proto/basic"
	"hlm-ipfs/sdk/proto/idp"
	"sync"
	"time"
)

func newUserCli() *userCli {
	return &userCli{}
}

type userCli struct {
	rw    sync.RWMutex
	id    int64
	token string
}

func (s *userCli) loop() {
	for range time.After(time.Minute * 20) {
		if _, err := s.UserInfo(); err != nil && api.opt.Err != nil {
			api.opt.Err(err)
		}
	}
}

func (s *userCli) Token() string {
	s.rw.RLock()
	defer s.rw.RUnlock()

	return s.token
}

func (s *userCli) Login() error {
	var (
		err error
		rq  = &idp.LoginRequest{
			AppID:  api.opt.AppID,
			Device: &idp.DeviceInfo{IMEI: ""},
			UserPassword: &idp.LoginByUserNameRequest{
				UserName: api.opt.Name,
				Password: api.opt.Pwd,
			},
		}
		rs = &idp.LoginResponse{}
	)

	if err = Do("/idp/idp/login", rq, rs); err != nil {
		return err
	} else {
		s.rw.Lock()
		s.id = rs.UserID
		s.token = rs.AccessToken
		s.rw.Unlock()

		go s.loop()
	}
	return nil
}

func (s *userCli) Logout() error {
	return Do("/idp/idp/logout", &basic.Empty{}, &basic.Empty{})
}

func (s *userCli) UserInfo() (*idp.UserInfo, error) {
	rs := &idp.UserInfo{}
	if err := Do("/idp/idp/user", &basic.Int64{Value: s.id}, rs); err != nil {
		return nil, err
	}
	return rs, nil
}
