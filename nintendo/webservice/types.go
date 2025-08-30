package webservice

type GetWebServiceTokenRequest struct {
	Parameter struct {
		F                 string `json:"f"`
		Id                int64  `json:"id"`
		RegistrationToken string `json:"registrationToken"`
		RequestId         string `json:"requestId"`
		Timestamp         int64  `json:"timestamp"`
	} `json:"parameter"`
}

type GetWebServiceTokenResponse struct {
	Status int `json:"status"`
	Result struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:"expiresIn"`
	} `json:"result"`
	CorrelationId string `json:"correlationId"`
}

type ListWebServicesResponse struct {
	Status int `json:"status"`
	Result []struct {
		Id               int64  `json:"id"`
		Uri              string `json:"uri"`
		CustomAttributes []struct {
			AttrValue string `json:"attrValue"`
			AttrKey   string `json:"attrKey"`
		} `json:"customAttributes"`
		WhiteList []string `json:"whiteList"`
		Name      string   `json:"name"`
		ImageUri  string   `json:"imageUri"`
	} `json:"result"`
	CorrelationId string `json:"correlationId"`
}
