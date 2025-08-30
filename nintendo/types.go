package nintendo

type UserResponse struct {
	Birthday                  string `json:"birthday"`
	AnalyticsOptedInUpdatedAt int    `json:"analyticsOptedInUpdatedAt"`
	IconUri                   string `json:"iconUri"`
	EmailOptedIn              bool   `json:"emailOptedIn"`
	Timezone                  struct {
		Name             string `json:"name"`
		UtcOffsetSeconds int    `json:"utcOffsetSeconds"`
		UtcOffset        string `json:"utcOffset"`
		Id               string `json:"id"`
	} `json:"timezone"`
	ClientFriendsOptedIn  bool `json:"clientFriendsOptedIn"`
	EmailOptedInUpdatedAt int  `json:"emailOptedInUpdatedAt"`
	AnalyticsPermissions  struct {
		TargetMarketing struct {
			UpdatedAt int  `json:"updatedAt"`
			Permitted bool `json:"permitted"`
		} `json:"targetMarketing"`
		InternalAnalysis struct {
			Permitted bool `json:"permitted"`
			UpdatedAt int  `json:"updatedAt"`
		} `json:"internalAnalysis"`
	} `json:"analyticsPermissions"`
	Id               string      `json:"id"`
	CreatedAt        int         `json:"createdAt"`
	Nickname         string      `json:"nickname"`
	Region           interface{} `json:"region"`
	AnalyticsOptedIn bool        `json:"analyticsOptedIn"`
	IsChild          bool        `json:"isChild"`
	Language         string      `json:"language"`
	EachEmailOptedIn struct {
		Survey struct {
			UpdatedAt int  `json:"updatedAt"`
			OptedIn   bool `json:"optedIn"`
		} `json:"survey"`
		Deals struct {
			OptedIn   bool `json:"optedIn"`
			UpdatedAt int  `json:"updatedAt"`
		} `json:"deals"`
	} `json:"eachEmailOptedIn"`
	EmailVerified                 bool   `json:"emailVerified"`
	PhoneNumberEnabled            bool   `json:"phoneNumberEnabled"`
	UpdatedAt                     int    `json:"updatedAt"`
	ScreenName                    string `json:"screenName"`
	Country                       string `json:"country"`
	Gender                        string `json:"gender"`
	ClientFriendsOptedInUpdatedAt int    `json:"clientFriendsOptedInUpdatedAt"`
}

type FriendListResponse struct {
	Status int `json:"status"`
	Result struct {
		Friends []struct {
			Id               int64  `json:"id"`
			NsaId            string `json:"nsaId"`
			ImageUri         string `json:"imageUri"`
			Image2Uri        string `json:"image2Uri"`
			Name             string `json:"name"`
			IsFriend         bool   `json:"isFriend"`
			IsFavoriteFriend bool   `json:"isFavoriteFriend"`
			IsServiceUser    bool   `json:"isServiceUser"`
			IsNew            bool   `json:"isNew"`
			FriendCreatedAt  int    `json:"friendCreatedAt"`
			Route            struct {
				AppName  string `json:"appName"`
				UserName string `json:"userName"`
				ShopUri  string `json:"shopUri"`
				ImageUri string `json:"imageUri"`
				Channel  string `json:"channel"`
			} `json:"route"`
			Presence struct {
				State     string `json:"state"`
				UpdatedAt int    `json:"updatedAt"`
				LogoutAt  int    `json:"logoutAt"`
				Game      struct {
					Name           string `json:"name,omitempty"`
					ImageUri       string `json:"imageUri,omitempty"`
					ShopUri        string `json:"shopUri,omitempty"`
					TotalPlayTime  int    `json:"totalPlayTime,omitempty"`
					FirstPlayedAt  int    `json:"firstPlayedAt,omitempty"`
					SysDescription string `json:"sysDescription,omitempty"`
				} `json:"game"`
			} `json:"presence"`
		} `json:"friends"`
	} `json:"result"`
	CorrelationId string `json:"correlationId"`
}

type ShowSelfResponse struct {
	Status int `json:"status"`
	Result struct {
		Id                int64  `json:"id"`
		NsaId             string `json:"nsaId"`
		ImageUri          string `json:"imageUri"`
		Image2Uri         string `json:"image2Uri"`
		Name              string `json:"name"`
		SupportId         string `json:"supportId"`
		IsChildRestricted bool   `json:"isChildRestricted"`
		Etag              string `json:"etag"`
		Links             struct {
			NintendoAccount struct {
				Membership struct {
					Active bool `json:"active"`
				} `json:"membership"`
			} `json:"nintendoAccount"`
			FriendCode struct {
				Regenerable   bool   `json:"regenerable"`
				RegenerableAt int    `json:"regenerableAt"`
				Id            string `json:"id"`
			} `json:"friendCode"`
		} `json:"links"`
		Permissions struct {
			PlayLog                string `json:"playLog"`
			Presence               string `json:"presence"`
			FriendRequestReception bool   `json:"friendRequestReception"`
		} `json:"permissions"`
		Presence struct {
			State     string `json:"state"`
			UpdatedAt int    `json:"updatedAt"`
			LogoutAt  int    `json:"logoutAt"`
			Game      struct {
				Name           string `json:"name"`
				ImageUri       string `json:"imageUri"`
				ShopUri        string `json:"shopUri"`
				TotalPlayTime  int    `json:"totalPlayTime"`
				FirstPlayedAt  int    `json:"firstPlayedAt"`
				SysDescription string `json:"sysDescription"`
			} `json:"game"`
			Platform int `json:"platform"`
		} `json:"presence"`
	} `json:"result"`
	CorrelationId string `json:"correlationId"`
}
