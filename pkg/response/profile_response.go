package response

import "github.com/whitenight1201/go-devconnector/pkg/entity"

type ProfileResponse struct {
	ID             int64      `json:"id"`
	Company        string     `json:"company"`
	Website        string     `json:"website"`
	Location       string     `json:"location"`
	Status         string     `json:"status"`
	Skills         string     `json:"skills"`
	GithubUsername string     `json:"githubusername"`
	Bio            string     `json:"bio"`
	Social         SocialResp `json:"social"`
	Education      []string   `json:"education"`
	Experience     []string   `json:"experience"`
	UserID         string     `json:"user_id"`
}

type SocialResp struct {
	Twitter   string `json:"twitter"`
	Facebook  string `json:"facebook"`
	Linkedin  string `json:"linkedin"`
	Youtube   string `json:"youtube"`
	Instagram string `json:"instagram"`
}

func NewProfileResponse(profile entity.Profile) ProfileResponse {
	socialRes := SocialResp{
		Twitter:   profile.Twitter,
		Facebook:  profile.Facebook,
		Linkedin:  profile.Linkedin,
		Youtube:   profile.Youtube,
		Instagram: profile.Instagram,
	}

	educationresp := []string{}
	experienceresp := []string{}
	return ProfileResponse{
		ID:             profile.ID,
		Company:        profile.Company,
		Website:        profile.Website,
		Location:       profile.Location,
		Status:         profile.Status,
		Skills:         profile.Skills,
		GithubUsername: profile.GithubUsername,
		Bio:            profile.Bio,
		Social:         socialRes,
		Education:      educationresp,
		Experience:     experienceresp,
		UserID:         profile.UserID,
	}
}

func NewProfileResponseArray(profile []entity.Profile) []ProfileResponse {
	profileRes := []ProfileResponse{}
	for _, value := range profile {
		socialRes := SocialResp{
			Twitter:   value.Twitter,
			Facebook:  value.Facebook,
			Linkedin:  value.Linkedin,
			Youtube:   value.Youtube,
			Instagram: value.Instagram,
		}
		educationresp := []string{}
		experienceresp := []string{}

		profile := ProfileResponse{
			ID:             value.ID,
			Company:        value.Company,
			Website:        value.Website,
			Location:       value.Location,
			Status:         value.Status,
			Skills:         value.Skills,
			GithubUsername: value.GithubUsername,
			Bio:            value.Bio,
			Social:         socialRes,
			Education:      educationresp,
			Experience:     experienceresp,
			UserID:         value.UserID,
		}

		profileRes = append(profileRes, profile)
	}
	return profileRes
}
