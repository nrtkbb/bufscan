package bufscan

import (
	"bufio"
	"io"
	"log"
	"bytes"
)

func BufScan(r *bufio.Reader, f func(line string) error) error {
	var err error
	for i := 1; ; i++ {
		buf, isPrefix, err := r.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}

		if isPrefix {
			copyBuf := make([]byte, len(buf))
			copy(copyBuf, buf)
			bb := bytes.NewBuffer(copyBuf)
			for {
				b, cont, err := r.ReadLine()
				if err != nil {
					if err != io.EOF {
						log.Fatal(err)
					}
					break
				}

				if _, err := bb.Write(b); err != nil {
					log.Fatal(err)
				}

				if !cont {
					break
				}
			}
			err = f(bb.String())
		} else {
			bb := bytes.NewBuffer(buf)
			err = f(bb.String())
		}

	}
	return err
}
