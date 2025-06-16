package config

type Config struct {
	ServerPort                string `env:"SERVER_PORT"`
	IDTemplateDefault         string `env:"ID_TEMPLATE_DEFAULT"`
	ClientIDOfficer           string `env:"CLIENT_ID_OFFICER"`
	ClientSecretOfficer       string `env:"CLIENT_SECRET_OFFICER"`
	ClientIDOutlet            string `env:"CLIENT_ID_OUTLET"`
	ClientSecretOutlet        string `env:"CLIENT_SECRET_OUTLET"`
	VertexBaseURL             string `env:"VERTEX_BASE_URL"`
	IdentityBaseURL           string `env:"IDENTITY_BASE_URL"`
	IdentityClientID          string `env:"IDENTITY_CLIENT_ID"`
	IdentityClientSecret      string `env:"IDENTITY_CLIENT_SECRET"`
	IdentityRedirectURI       string `env:"IDENTITY_REDIRECT_URI"`
	CookieDomain              string `env:"COOKIE_DOMAIN"`
	DefaultWhatsappFeeAmount  string `env:"DEFAULT_WHATSAPP_FEE_AMOUNT"`
	IDCountryOnline           string `env:"ID_COUNTRY_ONLINE"`
	IDProvinceOnline          string `env:"ID_PROVINCE_ONLINE"`
	IDDistrictOnline          string `env:"ID_DISTRICT_ONLINE"`
	IDRegionOnline            string `env:"ID_REGION_ONLINE"`
	AuthHMACExpiry            string `env:"AUTH_HMAC_EXPIRY"`
	MailgunDomain             string `env:"MAILGUN_DOMAIN"`
	MailgunAPIKey             string `env:"MAILGUN_API_KEY"`
	NewRelicLicense           string `env:"NEWRELIC_LICENSE"`
	AllowOrigins              string `env:"ALLOW_ORIGINS"`
	AllowHeaders              string `env:"ALLOW_HEADERS"`
	RefreshThresholdSecond    string `env:"REFRESH_THRESHOLD_SECOND"`
	CookieSecure              string `env:"COOKIE_SECURE"`
	InternalAccessKey         string `env:"INTERNAL_ACCESS_KEY"`
	RateLimitThreshold        string `env:"RATE_LIMIT_THRESHOLD"`
	RateLimitBucketLeakSecond string `env:"RATE_LIMIT_BUCKET_LEAK_SECOND"`
	MailgunBaseURL            string `env:"MAILGUN_BASE_URL"`
	GoCoreBaseURL             string `env:"GOCORE_BASE_URL"`
	CloudflareBaseURL         string `env:"CLOUDFLARE_BASE_URL"`
	CloudflareZone            string `env:"CLOUDFLARE_ZONE"`
	CloudflareToken           string `env:"CLOUDFLARE_TOKEN"`
	SlackBaseURL              string `env:"SLACK_BASE_URL"`
	IdentityExpiredInSecond   string `env:"IDENTITY_EXPIRED_IN_SECOND"`
	MailerAwsRegion           string `env:"MAILER_AWS_REGION"`
	MailerAwsAccessKey        string `env:"MAILER_AWS_ACCESS_KEY"`
	MailerAwsSecretKey        string `env:"MAILER_AWS_SECRET_KEY"`
	DefaultQuotaSendEmail     string `env:"DEFAULT_QUOTA_SEND_EMAIL"`
	S3PrivateBucket        	  string `env:"S3_PRIVATE_BUCKET"`
	S3PrivateAccessKey        string `env:"S3_PRIVATE_ACCESS_KEY"`
	S3PrivateSecretKey        string `env:"S3_PRIVATE_SECRET_KEY"`
	S3PrivateRegion           string `env:"S3_PRIVATE_REGION"`
	S3PublicBucket            string `env:"S3_PUBLIC_BUCKET"`
	S3PublicAccessKey         string `env:"S3_PUBLIC_ACCESS_KEY"`
	S3PublicSecretKey         string `env:"S3_PUBLIC_SECRET_KEY"`
	S3PublicRegion            string `env:"S3_PUBLIC_REGION"`
	KafkaBrokers              string `env:"KAFKA_BROKERS"`
	KafkaTopicEventCdc        string `env:"KAFKA_TOPIC_EVENT_CDC"`
	PrefixPathUploadS3        string `env:"PREFIX_PATH_UPLOAD_S3"`
	OrionURL                  string `env:"ORION_URL"`
	DebugHTTP                 string `env:"DEBUG_HTTP"`
	SpotlightAssetLocation    string `env:"SPOTLIGHT_ASSET_LOCATION"`
	AssetBuildVersion         string `env:"ASSET_BUILD_VERSION"`
	SkipPurgeCF               string `env:"SKIP_PURGE_CF"`
}

func NewConfig() *Config {
	c := &Config{}
	LoadEnv()
	MarshalEnv(c)
	if c.DefaultQuotaSendEmail == "" {
		c.DefaultQuotaSendEmail = "1"
	}
	if c.RateLimitThreshold == "" {
		c.RateLimitThreshold = "1000"
	}
	if c.RateLimitBucketLeakSecond == "" {
		c.RateLimitBucketLeakSecond = "1"
	}
	return c
}
