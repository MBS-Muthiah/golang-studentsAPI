package response
import (
	"net/http"
	"encoding/json"
	
)

type Response struct{
	status string `json:"status"`
	message string `json:"message"`
}

const(
	StatusOK ="ok"
	StatusError ="error"
)
func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) Response{
	return Response{
		status: StatusError,
		message: err.Error()
	}
}