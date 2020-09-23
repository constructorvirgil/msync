package setting

type ServerSettingS struct {
	IP   string
	Port string
}

type ClientSettingS struct {
	UserName string
	Password string
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	return nil
}
