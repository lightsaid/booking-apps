package configs

import (
	"errors"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 配置对象，使用 viper 包
type Config struct {
	vp *viper.Viper
}

// NewConfig 创建一个读取配置的对象
// filename 文件名，包括扩展名
// paths 多配所在的目录集合
// obj 将结果映射到obj
func NewConfig(filename string, obj interface{}, paths ...string) (*Config, error) {
	vp := viper.New()
	for _, path := range paths {
		if path != "" {
			vp.AddConfigPath(path)
		}
	}

	parts := strings.Split(filename, ".")
	if len(parts) != 2 {
		return nil, errors.New("filename invalid")
	}

	// 设置配置文件名
	vp.SetConfigName(parts[0])
	// 设置配置扩展名
	vp.SetConfigType(parts[1])

	// 读取配置进来
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		vp: vp,
	}

	err = cfg.vp.Unmarshal(obj)
	if err != nil {
		return nil, err
	}

	// 监听
	cfg.Watch(obj)

	return cfg, nil
}

// Watch 监听 配置是否发生变化，更改了就立马重载
func (cfg *Config) Watch(obj interface{}) {
	go func() {
		cfg.vp.WatchConfig()
		cfg.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = cfg.vp.Unmarshal(obj)
		})
	}()
}
