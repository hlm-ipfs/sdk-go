package sdk

import (
	"hlm-ipfs/sdk/proto/basic"
	"hlm-ipfs/sdk/proto/market"
)

func newMarketCli() *marketCli {
	return &marketCli{}
}

type marketCli struct {
}

func (s *marketCli) GetConfig() (*market.ProfileInfo, error) {
	rs := &market.ProfileInfo{}
	if err := Do("/market/user/Profile", &basic.Empty{}, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *marketCli) SetConfig(in *market.SettingRequest) error {
	return Do("/market/user/Setting", in, &basic.Empty{})
}

func (s *marketCli) MarketCount() (*market.CountResponse, error) {
	rs := &market.CountResponse{}
	if err := Do("/market/market/count", &basic.Empty{}, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *marketCli) StorageOrders(in *market.StorageListRequest) (*market.StorageListResponse, error) {
	rs := &market.StorageListResponse{}
	if err := Do("/market/storage/list", in, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *marketCli) RetrievalOrders(in *market.RetrievalListRequest) (*market.RetrievalListResponse, error) {
	rs := &market.RetrievalListResponse{}
	if err := Do("/market/retrieval/list", in, rs); err != nil {
		return nil, err
	}
	return rs, nil
}
