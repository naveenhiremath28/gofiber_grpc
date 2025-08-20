package models

import (
	"encoding/json"
	"time"
	"github.com/google/uuid"
)

type APIResponse struct {
	ID           string      `json:"id"`
	Version      string      `json:"ver"`
	Timestamp    string      `json:"ts"`
	Params       Params      `json:"params"`
	ResponseCode string      `json:"responseCode"`
	Result       interface{} `json:"result"`
}

type Params struct {
	MsgID string `json:"msgid"`
}

type ApiRequest struct {
	ID        string          `json:"id"`
	Version   string          `json:"ver"`
	Timestamp string          `json:"ts"`
	Params    Params          `json:"params"`
	Request   json.RawMessage `json:"request"`
}

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"size:50;not null;unique" json:"username"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	FullName  string    `gorm:"size:100" json:"full_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}

func GetApiResponse(id string, response_code string, result any) APIResponse {
	return APIResponse{
		ID:        id,
		Version:   "v1",
		Timestamp: time.Now().Format(time.RFC3339),
		Params: Params{
			MsgID: uuid.New().String(),
		},
		ResponseCode: response_code,
		Result:       result,
	}
}
