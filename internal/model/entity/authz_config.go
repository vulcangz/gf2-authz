package entity

import (
	"fmt"
	"time"
)

const (
	DriverPostgres = "postgres"
	DriverMysql    = "mysql"
	DriverSqlite   = "sqlite"
)

type AppConfig struct {
	Name    string     `json:"Name"`
	Audit   AppAudit   `json:"audit"`
	Metrics AppMetrics `json:"metrics"`
	Stats   AppStats   `json:"stats"`
	Trace   AppTrace   `json:"trace"`
}

type AppAudit struct {
	CleanDelay        time.Duration `json:"cleanDelay"`
	CleanDaysToKeep   int           `json:"cleanDaysToKeep"`
	FlushDelay        time.Duration `json:"flushDelay"`
	ResourceKindRegex string        `json:"resourceKindRegex"`
}

type AppStats struct {
	CleanDelay        time.Duration `json:"cleanDelay"`
	CleanDaysToKeep   int           `json:"cleanDaysToKeep"`
	FlushDelay        time.Duration `json:"flushDelay"`
	ResourceKindRegex string        `json:"resourceKindRegex"`
}

type AppMetrics struct {
	Enabled bool `json:"enabled"`
}

type AppTrace struct {
	Enabled         bool          `json:"enabled"`
	Exporter        string        `json:"exporter"`
	JaegerEndpoint  string        `json:"jaegerEndpoint"`
	OtlpDialTimeout time.Duration `json:"otlpDialTimeout"`
	OtlpEndpoint    string        `json:"otlpEndpoint"`
	ZipkinURL       string        `json:"zipkinURL"`
	SampleRatio     float64       `json:"sampleRatio"`
}

func (AppConfig) DefaultConfig() *AppConfig {
	return &AppConfig{
		Audit: AppAudit{
			CleanDelay:        1 * time.Hour,
			CleanDaysToKeep:   7,
			FlushDelay:        3 * time.Second,
			ResourceKindRegex: `.*`,
		},
		Metrics: AppMetrics{
			Enabled: false,
		},
		Stats: AppStats{
			CleanDelay:        1 * time.Hour,
			CleanDaysToKeep:   30,
			FlushDelay:        3 * time.Second,
			ResourceKindRegex: `.*`,
		},
		Trace: AppTrace{
			Enabled:         false,
			Exporter:        "jaeger",
			JaegerEndpoint:  "localhost:14250",
			OtlpDialTimeout: 3 * time.Second,
			OtlpEndpoint:    "localhost:30080",
			ZipkinURL:       "http://localhost:9411/api/v2/spans",
			SampleRatio:     1.0,
		},
	}
}

type AuthConfig struct {
	AccessTokenDuration  time.Duration `json:"accessTokenDuration"`
	RefreshTokenDuration time.Duration `json:"refreshTokenDuration"`
	Domain               string        `json:"domain"`
	JWTSignString        []byte        `json:"jwtSignString"`
}

func (AuthConfig) DefaultConfig() *AuthConfig {
	return &AuthConfig{
		AccessTokenDuration:  6 * time.Hour,
		RefreshTokenDuration: 6 * time.Hour,
		Domain:               "http://localhost:8080",
		JWTSignString:        []byte(`4uthz-s3cr3t-valu3-pl3as3-ch4ng3!`),
	}
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSLMode  string `json:"ssl"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"name"`
	Timezone string `json:"timezone"`
}

func (d DatabaseConfig) MysqlDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Dbname,
	)
}

func (d DatabaseConfig) PostgresDSN() string {
	return "host=" + d.Host +
		" user=" + d.User +
		" password=" + d.Password +
		" dbname=" + d.Dbname +
		" port=" + d.Port +
		" sslmode=" + d.SSLMode +
		" TimeZone=" + d.Timezone
}

func (d DatabaseConfig) SqliteDSN() string {
	return fmt.Sprintf("file:%s?cache=shared&mode=rwc&_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)", d.Dbname)
}

func (DatabaseConfig) DefaultConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   "mysql",
		Host:     "localhost",
		User:     "root",
		Password: "password",
		Dbname:   "authz",
		Port:     "3306",
		SSLMode:  "disable",
		Timezone: "UTC",
	}
}

type GRPCServerConfig struct {
	Name             string             `json:"name"`
	Address          string             `json:"address"`
	LogPath          string             `json:"logPath"`
	LogStdout        bool               `json:"logStdout"`
	ErrorLogEnabled  bool               `json:"errorLogEnabled"`
	AccessLogEnabled bool               `json:"accessLogEnabled"`
	ErrorStack       bool               `json:"errorStack"`
	Registry         *GRPCRegistyConfig `json:"registry"`
}

type GRPCRegistyConfig struct {
	Schema               string   `json:"schema"`               // 指定registry 类型协议
	Endpoints            []string `json:"endpoints"`            // 指定 etcd 服务端地址
	DialTimeout          int64    `json:"dialTimeout"`          // etcd 链接超时时间, 单位 second, 默认5s
	DialKeepAliveTime    int64    `json:"dialKeepAliveTime"`    // etcd keepalive 时间, 单位second, 默认5s
	DialKeepAliveTimeout int64    `json:"dialKeepAliveTimeout"` // etcd keepalive 超时时间, 单位 second, 默认5s
	Username             string   `json:"username"`             // etcd 认证 username
	Password             string   `json:"password"`             // etcd 认证 password
}

func (GRPCServerConfig) DefaultConfig() *GRPCServerConfig {
	return &GRPCServerConfig{
		Address: ":8081",
	}
}

type HTTPServerConfig struct {
	Address              string        `json:"address"`
	CORSAllowedDomains   []string      `json:"corsAllowedDomains"`
	CORSAllowedMethods   []string      `json:"corsAllowedMethods"`
	CORSAllowedHeaders   []string      `json:"corsAllowedHeaders"`
	CORSAllowCredentials bool          `json:"corsAllowCredentials"`
	CORSCacheMaxAge      time.Duration `json:"corsCacheMaxAge"`
}

func (HTTPServerConfig) DefaultConfig() *HTTPServerConfig {
	return &HTTPServerConfig{
		Address: ":8080",
	}
}

type OAuthConfig struct {
	Provider            string   `json:"provider"`
	ClientID            string   `json:"clientId"`
	ClientSecret        string   `json:"clientSecret"`
	CookiesDomainName   string   `json:"cookiesDomainName"`
	FrontendRedirectURL string   `json:"frontendRedirectURL"`
	IssuerURL           string   `json:"issuerURL"`
	RedirectURL         string   `json:"redirectURL"`
	Scopes              []string `json:"scopes"`
}

func (OAuthConfig) DefaultConfig() *OAuthConfig {
	return &OAuthConfig{
		CookiesDomainName:   "localhost",
		FrontendRedirectURL: "http://localhost:3000",
		RedirectURL:         "http://localhost:8080/v1/oauth/callback",
		Scopes:              []string{},
	}
}

type EventConfig struct {
	DispatcherEventChannelSize int `json:"dispatcherEventChannelSize"`
}

func (EventConfig) DefaultConfig() *EventConfig {
	return &EventConfig{
		DispatcherEventChannelSize: 10000,
	}
}

type UserConfig struct {
	AdminDefaultPassword string `json:"adminDefaultPassword"`
}

func (UserConfig) DefaultConfig() *UserConfig {
	return &UserConfig{
		AdminDefaultPassword: "changeme",
	}
}
