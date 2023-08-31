package dto

type CreateProfileRequest struct {
	Company        string `json:"company"`
	Website        string `json:"website"`
	Location       string `json:"location"`
	Status         string `json:"status"`
	Skills         string `json:"skills"`
	GithubUsername string `json:"githutusername"`
	Bio            string `json:"bio"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Linkedin       string `json:"linkedin"`
	Youtube        string `json:"youtube"`
	Instagram      string `json:"instagram"`
	UserID         string `json:"user_id"`
}

type UpdateProfileRequest struct {
	ID             int64  `json:"id"`
	Company        string `json:"company"`
	Website        string `json:"website"`
	Location       string `json:"location"`
	Status         string `json:"status"`
	Skills         string `json:"skills"`
	GithubUsername string `json:"githutusername"`
	Bio            string `json:"bio"`
	Twitter        string `json:"twitter"`
	Facebook       string `json:"facebook"`
	Linkedin       string `json:"linkedin"`
	Youtube        string `json:"youtube"`
	Instagram      string `json:"instagram"`
	UserID         string `json:"user_id"`
}
