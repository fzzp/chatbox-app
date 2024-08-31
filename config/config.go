package config

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

// LoadConfig 创建一个全局读取配置的对象
//
// filename 文件名，包括扩展名
//
// paths 从那些目录查找配置至少得有一个
//
// obj 将结果映射到target
func LoadConfig(filename string, target interface{}, paths ...string) error {
	vp := viper.New()
	for _, path := range paths {
		if path != "" {
			vp.AddConfigPath(path)
		}
	}

	parts := strings.Split(filename, ".")
	if len(parts) != 2 {
		return errors.New("filename invalid")
	}

	// 设置配置文件名
	vp.SetConfigName(parts[0])

	// 设置配置扩展名
	vp.SetConfigType(parts[1])

	// 读取配置进来
	err := vp.ReadInConfig()
	if err != nil {
		return err
	}

	cfg := &Config{
		vp: vp,
	}

	err = cfg.vp.Unmarshal(target)
	if err != nil {
		return err
	}

	// 监听
	cfg.Watch(target)

	return nil
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
