package main

import (
	"bytes"
	"os"

	"gopkg.in/yaml.v3"
)

var Conf *Config

type Config struct {
	Interval      int            `yaml:"interval"`
	TencentConfig *TencentConfig `yaml:"tencent_config"`
	Socks5Config  string         `yaml:"socks_5_config"`
}

type TencentConfig struct {
	Domain     string `yaml:"domain"`
	RecordType string `yaml:"record_type"`
	SecretId   string `yaml:"secret_id"`
	SecretKey  string `yaml:"secret_key"`
	SubDomain  string `yaml:"sub_domain"`
}

func init() {
	Conf = &Config{}
	f, err := os.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(f))
	if err = decoder.Decode(Conf); err != nil {
		panic(err)
	}
}

func (c *Config) GetInterval() int {
	if c == nil {
		return 0
	}
	return c.Interval
}

func (c *Config) GetTencentConfig() *TencentConfig {
	if c == nil {
		return nil
	}
	return c.TencentConfig
}

func (c *Config) GetSocks5Config() string {
	if c == nil {
		return ""
	}
	return c.Socks5Config
}

func (t *TencentConfig) GetDomain() string {
	if t == nil {
		return ""
	}
	return t.Domain
}

func (t *TencentConfig) GetRecordType() string {
	if t == nil {
		return ""
	}
	return t.RecordType
}

func (t *TencentConfig) GetSecretId() string {
	if t == nil {
		return ""
	}
	return t.SecretId
}

func (t *TencentConfig) GetSecretKey() string {
	if t == nil {
		return ""
	}
	return t.SecretKey
}

func (t *TencentConfig) GetSubDomain() string {
	if t == nil {
		return ""
	}
	return t.SubDomain
}
