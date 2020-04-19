package register

const (
	defaultPasswordLen = 8
)

//Request register request
type Request struct {
	Username string
	Email    string
	Password string
}

// Response register response
type Response struct {
	Token string
}
