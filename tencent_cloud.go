package main

import (
	"context"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
)

type TencentCloudUpdater struct {
	client *dnspod.Client
}

func NewTencentCloudUpdater() (*TencentCloudUpdater, error) {
	client, err := dnspod.NewClient(
		common.NewCredential(Conf.GetTencentConfig().GetSecretId(), Conf.GetTencentConfig().GetSecretKey()),
		"",
		profile.NewClientProfile())
	if err != nil {
		return nil, err
	}
	return &TencentCloudUpdater{
		client: client,
	}, nil
}

func (t *TencentCloudUpdater) Update(ctx context.Context, ipv6 string) error {
	// TODO: implement
	return nil
}

func (t *TencentCloudUpdater) DescribeRecordList(ctx context.Context) (response *dnspod.DescribeRecordListResponse, err error) {
	if t == nil || t.client == nil {
		return nil, fmt.Errorf("client is nil")
	}
	req := dnspod.NewDescribeRecordListRequest()
	req.Domain = common.StringPtr(Conf.GetTencentConfig().GetDomain())
	if Conf.GetTencentConfig().GetRecordType() != "" {
		req.RecordType = common.StringPtr(Conf.GetTencentConfig().GetRecordType())
	}
	return t.client.DescribeRecordListWithContext(ctx, req)
}

func (t *TencentCloudUpdater) ModifyDynamicDNS(ctx context.Context, record *dnspod.RecordListItem, ipv6 string) error {
	req := dnspod.NewModifyDynamicDNSRequest()
	req.Domain = common.StringPtr(Conf.GetTencentConfig().GetDomain())
	req.SubDomain = record.Name
	req.RecordId = record.RecordId
	req.RecordLine = record.Line
	req.Value = common.StringPtr(ipv6)
	return nil
}
