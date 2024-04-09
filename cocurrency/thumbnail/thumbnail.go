package thumbnail

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func Image(src image.Image) image.Image {
	xs := src.Bounds().Size().X
	ys := src.Bounds().Size().Y
	width, height := 128, 128
	if aspect := float64(xs) / float64(ys); aspect < 1.0 {
		width = int(128 *aspect)
	} else {
		height = int(128 /aspect)
	}
	xscale := float64(xs) / float64(width)
	yscale := float64(ys) / float64(height)

	dst := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			srcx := int(float64(x) * xscale)
			srcy := int(float64(y) * yscale)
			dst.Set(x, y, src.At(srcx, srcy))
		}
	}
	return dst
}

func ImageStream(w io.Writer, r io.Reader) error {
	src, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	dst := Image(src)
	return jpeg.Encode(w, dst, nil)
}

func ImageFile2(outfile, infile string) (err error) {
	in, err := os.Open(infile)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(outfile)
	if err != nil {
		return err
	}
	if err := ImageStream(out, in); err != nil {
		out.Close()
		return fmt.Errorf("scaling %s to %s: %s", infile, outfile, err)
	}
	return out.Close()
}

func ImageFile(infile string) (string, error) {
	ext := filepath.Ext(infile)
	outFile := strings.Trim(infile, ext) + ".thumb" + ext
	return outFile, ImageFile2(outFile, infile)
}

func makeThumbnails(filesnames []string) {
	for _, f := range filesnames {
		go ImageFile(f)
	}
}

func makeThumbNail(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<- ch
	}
}

// solutions that no goroutine is able to draine the channel
func makeThumbnailNoDraineds(filenames []string) (error){
	ch := make(chan error, len(filenames))
	for _, f := range(filenames) {
		go func(f string) {
			_, err := ImageFile(f)
			ch <- err
		} (f)
	}

	for range(filenames) {
		if e := <- ch; e != nil {
			return e
		}
	}
	return nil
}

func makeThumbNailWaitGroup(filenames []string) int {
	var wg sync.WaitGroup
	ch := make(chan int)
	for _, f := range(filenames) {
		wg.Add(1)
		go func(f string) {
			file, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(file)
			ch <- int(info.Size())
			wg.Done()
		}(f)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var totalSize int
	for s := range ch {
		totalSize += s
	}
	return totalSize
}