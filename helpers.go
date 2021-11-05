package SibylSystem

func Client(token string, opts *ClientOpts) (*SibylSystem, error) {
	if opts == nil {
		opts = &ClientOpts{
			ApiUrl: DefaultApiUrl,
		}
	}
	if opts.ApiUrl == "" {
		opts.ApiUrl = DefaultApiUrl
	}
	if string(opts.ApiUrl[len(opts.ApiUrl)-1]) != "/" {
		opts.ApiUrl = opts.ApiUrl + "/"
	}
	var s = &SibylSystem{}
	s.TimeOut = opts.TimeOut
	s.Token = token
	s.ApiUrl = opts.ApiUrl
	// Ignore error on CheckToken method as it would already be handled in succeeding calls
	t, err := s.CheckToken()
	if !t {
		if err != nil {
			return nil, err
		}
		return nil, ErrInvalidToken
	}
	return s, nil
}
