package __interface

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

//tar命令
func TestTar(t *testing.T) {
	tr := tar.NewReader(os.Stdin)
	for {
		hdr, err := tr.Next()
		if err != nil {
			return
		}
		fmt.Println(hdr.Name)
		fileInfo := hdr.FileInfo()
		if fileInfo.IsDir(){
			continue
		}
		io.Copy(ioutil.Discard, tr)
	}
}

