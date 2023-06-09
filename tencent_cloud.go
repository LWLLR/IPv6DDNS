package main

import (
	"context"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	"golang.org/x/net/proxy"
	"net/http"
	"net/url"
)

type TencentCloudUpdater struct {
	client *dnspod.Client
}

func NewTencentCloudUpdater() (*TencentCloudUpdater, error) {
	c, err := dnspod.NewClient(
		common.NewCredential(Conf.GetTencentConfig().GetSecretId(), Conf.GetTencentConfig().GetSecretKey()),
		"",
		profile.NewClientProfile())
	if err != nil {
		return nil, err
	}

	if Conf.GetSocks5Config() != "" {
		proxyURL, err := url.Parse(Conf.GetSocks5Config())
		if err != nil {
			return nil, err
		}
		dialer, err := proxy.FromURL(proxyURL, proxy.Direct)
		if err != nil {
			return nil, err
		}
		dialerCtx, ok := dialer.(proxy.ContextDialer)
		if !ok {
			return nil, fmt.Errorf("implement ContextDialer fail")
		}
		tr := &http.Transport{
			DialContext: dialerCtx.DialContext,
		}
		c.WithHttpTransport(tr)
		fmt.Printf("use proxy:%s\n", Conf.GetSocks5Config())
	}
	return &TencentCloudUpdater{
		client: c,
	}, nil
}

func (t *TencentCloudUpdater) Update(ctx context.Context, ipv6 string) error {
	subDomain := Conf.GetTencentConfig().GetSubDomain()
	if subDomain == "" {
		return fmt.Errorf("sub_domain is empty")
	}
	descRsp, err := t.DescribeRecordList(ctx)
	if err != nil {
		return err
	}
	if descRsp == nil || descRsp.Response == nil {
		return nil
	}
	record := t.GetRecordBySubdomain(subDomain, descRsp.Response.RecordList)
	if record == nil || record.Value == nil {
		return nil
	}
	if *record.Value != ipv6 {
		fmt.Printf("ip不一致，旧ip:%s;新ip:%s\n", *record.Value, ipv6)
		if err = t.ModifyDynamicDNS(ctx, record, ipv6); err != nil {
			return err
		}
	}
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

func (t *TencentCloudUpdater) GetRecordBySubdomain(subDomain string, list []*dnspod.RecordListItem) *dnspod.RecordListItem {
	for _, item := range list {
		if item == nil || item.Name == nil {
			continue
		}
		if *item.Name == subDomain {
			return item
		}
	}
	return nil
}

func (t *TencentCloudUpdater) ModifyDynamicDNS(ctx context.Context, record *dnspod.RecordListItem, ipv6 string) error {
	req := dnspod.NewModifyDynamicDNSRequest()
	req.Domain = common.StringPtr(Conf.GetTencentConfig().GetDomain())
	req.SubDomain = record.Name
	req.RecordId = record.RecordId
	req.RecordLine = record.Line
	req.Value = common.StringPtr(ipv6)
	rsp, err := t.client.ModifyDynamicDNSWithContext(ctx, req)
	if err != nil {
		fmt.Println("ModifyDynamicDNS requestID:", rsp.Response.RequestId)
		return err
	}
	return nil
}
