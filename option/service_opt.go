package option

type ServiceOpt struct {
	DogsVsCatsServiceOpt DogsVsCatsServiceOpt `mapstructure:"dogs_vs_cats_service_opt"`
}

type DogsVsCatsServiceOpt struct {
	Endpoint           string `mapstructure:"endpoint"`
	Bucket             string `mapstructure:"bucket"`
	AccessKeyID        string `mapstructure:"access_key_id"`
	AccessKeySecret    string `mapstructure:"access_key_secret"`
	RegionID           string `mapstructure:"region_id"`
	RoleArn            string `mapstructure:"role_arn"`
	RoleSessionName    string `mapstructure:"role_session_name"`
	UploadPathPrefix   string `mapstructure:"upload_path_prefix"`
	ClassifyPathPrefix string `mapstructure:"classify_path_prefix"`
	ModelPath          string `mapstructure:"model_path"`
}
