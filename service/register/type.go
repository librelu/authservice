package register

const (
	defaultPasswordLen = 8
)

//Request request description
type Request struct {
	Username string
	Email    string
	Password string
}
