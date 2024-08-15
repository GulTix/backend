package oauth

func (s *ServiceImpl) GetGoogleLoginURL() string {
	return s.oauth.GetRedirectURL()
}
