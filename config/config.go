package config

import (
	"strings"

	"gitee.com/lyhuilin/log"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	//初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	//初始化日志包
	c.initLog()
	//监控配置文件变化并热加载程序
	c.watchConfig()
	return nil
}
func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) //如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml") //设置配置文件格式为YAML
	viper.AutomaticEnv()        //读取匹配的环境变量
	viper.SetEnvPrefix("APP")   //读取环境变量的前缀为GOSPIDER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { //viper 解析配置文件
		// log.Debugf("配置文件解析错误: %s\n", err)
		return err
	}

	return nil
}

// 日志配置加载
func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	log.InitWithConfig(&passLagerCfg)
}

//监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
}
