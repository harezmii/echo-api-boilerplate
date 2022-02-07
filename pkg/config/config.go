package configGetter

import "github.com/spf13/viper"

const (
	STRING  = "string"
	INTEGER = "int"
	BOOL    = "bool"
)

func Get(configName string, configSelection string) interface{} {
	if loadEnvironment() {
		switch configSelection {
		case STRING:
			{
				return viper.GetString(configName)
			}
		case BOOL:
			{
				return viper.GetBool(configName)
			}
		case INTEGER:
			{
				return viper.GetInt(configName)
			}
		}
	}
	return nil
}

func loadEnvironment() bool {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return false
	}
	return true
}
