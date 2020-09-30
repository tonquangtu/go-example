package reader

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (outReader *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = outReader.r.Read(b)
	if err != nil {
		return n, err
	}
	for i := 0; i < n; i++ {
		b[i] += 13
	}
	return n, err
}

func readAbc()  {

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
