package sdk

import (
	"hlm-ipfs/sdk/proto/basic"
	"hlm-ipfs/sdk/proto/market"
)

func newBucketCli() *bucketCli {
	return &bucketCli{}
}

type bucketCli struct {
}

func (s *bucketCli) BucketNodes(parentID int64) (*market.BucketNodesResponse, error) {
	rs := &market.BucketNodesResponse{}
	if err := Do("/market/bucket/nodes", &basic.Int64{Value: parentID}, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *bucketCli) BucketList(parentID int64) (*market.BucketListResponse, error) {
	rs := &market.BucketListResponse{}
	if err := Do("/market/bucket/list", &basic.Int64{Value: parentID}, rs); err != nil {
		return nil, err
	}
	return rs, nil
}

func (s *bucketCli) CreateBucket(in *market.BucketCreateRequest) (int64, error) {
	rs := &basic.Int64{}
	if err := Do("/market/bucket/create", in, rs); err != nil {
		return 0, err
	}
	return rs.Value, nil
}

func (s *bucketCli) UpdateBucket(in *market.BucketUpdateRequest) error {
	return Do("/market/bucket/update", in, &basic.Empty{})
}

func (s *bucketCli) RemoveBucket(id int64) error {
	return Do("/market/bucket/remove", &basic.Int64{Value: id}, &basic.Empty{})
}
