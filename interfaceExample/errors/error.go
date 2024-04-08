package errors

type error interface {
	Error() string
}
type errorString struct {
	text string
}
func New(e string) error {
	return &errorString{e}
}

func (e errorString) Error() string {
	return e.text
}

func main() {

}