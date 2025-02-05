package config

type (
	Config struct {
		DBConnectionString        string
		Port                      string
		ClientId                  string
		ClientSecret              string
		RedirectUrl               string
		JwtSecret                 string
		JwtExpiredTime            string
		JwtIssuer                 string
		MidtransServerKey         string
		Env                       string
		ApiKey                    string
		BaseUrl                   string
		DefaultBlasterAccessToken string
		DefaultRefreshToken       string
		DefaultTokenType          string
		DefaultExpiredTime        string
	}
)
