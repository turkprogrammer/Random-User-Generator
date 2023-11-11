package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UserRepository интерфейс для получения случайного пользователя
type UserRepository interface {
	GetRandomUser() (*User, error)
}

// RandomUserRepository реализация UserRepository
type RandomUserRepository struct {
	apiURL string
}

// NewRandomUserRepository создает новый экземпляр RandomUserRepository
func NewRandomUserRepository(apiURL string) *RandomUserRepository {
	return &RandomUserRepository{
		apiURL: apiURL,
	}
}

// GetRandomUser возвращает случайного пользователя из RandomUser API
func (r *RandomUserRepository) GetRandomUser() (*User, error) {
	resp, err := http.Get(r.apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to RandomUser API: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response RandomUserResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %v", err)
	}

	if len(response.Results) == 0 {
		return nil, fmt.Errorf("no user found in the response")
	}

	result := response.Results[0]
	user := &User{
		FirstName: result.Name.First,
		LastName:  result.Name.Last,
		Email:     result.Email,
		Gender:    getString(result.Gender),
		Location: Location{
			Street: Street{
				Number: result.Location.Street.Number,
				Name:   result.Location.Street.Name,
			},
			City:     result.Location.City,
			State:    result.Location.State,
			Country:  result.Location.Country,
			Postcode: result.Location.Postcode,
			Coordinates: Coordinates{
				Latitude:  result.Location.Coordinates.Latitude,
				Longitude: result.Location.Coordinates.Longitude,
			},
			Timezone: Timezone{
				Offset:      result.Location.Timezone.Offset,
				Description: result.Location.Timezone.Description,
			},
		},
		Login: Login{
			UUID:     result.Login.UUID,
			Username: result.Login.Username,
			Password: result.Login.Password,
			Salt:     result.Login.Salt,
			MD5:      result.Login.MD5,
			SHA1:     result.Login.SHA1,
			SHA256:   result.Login.SHA256,
		},
		DOB: DOB{
			Date: result.DOB.Date,
			Age:  result.DOB.Age,
		},
		Registered: Registered{
			Date: result.Registered.Date,
			Age:  result.Registered.Age,
		},
		Phone: result.Phone,
		Cell:  result.Cell,
		ID: ID{
			Name:  result.ID.Name,
			Value: result.ID.Value,
		},
		Nat: result.Nat,
		Picture: Picture{
			Large:     result.Picture.Large,
			Medium:    result.Picture.Medium,
			Thumbnail: result.Picture.Thumbnail,
		},
	}

	return user, nil
}

// getString возвращает строковое значение или пустую строку, если значение равно nil
func getString(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}
