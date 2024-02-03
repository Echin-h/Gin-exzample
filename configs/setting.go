package configs

import "github.com/spf13/viper"

// 配置一个基础
func NewSetting() (*Setting, error) {
	vp := viper.New()
	// 配置基础信息
	vp.SetConfigName("config")
	vp.AddConfigPath("./configs")
	vp.AddConfigPath(".") //如果在当前目录找不到的话，可以在根目录上找
	vp.SetConfigType("yaml")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}

func SetUpSettings() error {
	setting, err := NewSetting()
	if err != nil {
		return err
	}
	//...可以有多个Section
	err1 := setting.ReadSection("Database", &DbSettings)
	if err1 != nil {
		return err1
	}

	err2 := setting.ReadSection("jwt", &JwtSettings)
	if err2 != nil {
		return err2
	}

	return nil
}

// 将yaml数据绑定到结构体上
func (s *Setting) ReadSection(name string, v interface{}) error {
	err := s.vp.UnmarshalKey(name, v)
	if err != nil {
		return err
	}

	return nil
}
