package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func search(ctx context.Context, root string, pattern string) ([]string, error) {
	g, ctx := errgroup.WithContext(ctx)
	paths := make(chan string, 100)

	g.Go(func() error {
		defer close(paths)

		return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.Mode().IsRegular() {
				return nil
			}

			if !info.IsDir() && !strings.HasSuffix(info.Name(), ".go") {
				return nil
			}

			select {
			case paths <- path:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	})

	c := make(chan string, 100)
	for path := range paths {
		p := path
		g.Go(func() error {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			if !bytes.Contains(data, []byte(pattern)) {
				return nil
			}

			select {
			case c <- p:
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}

	go func() {
		g.Wait()
		close(c)
	}()

	var m []string
	for r := range c {
		m = append(m, r)
	}
	return m, g.Wait()
}

func main() {
	duration := flag.Duration("timeout", 5*time.Millisecond, "timeout in millisecond")
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Printf("       filegrep [flags] path pattern \n")
		fmt.Println("Flags:")
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(-1)
	}

	path := flag.Arg(0)
	pattern := flag.Arg(1)

	ctx, _ := context.WithTimeout(context.Background(), *duration)
	m, err := search(ctx, path, pattern)

	if err != nil {
		log.Fatal(err)
	}

	for _, name := range m {
		fmt.Println(name)
	}

	fmt.Println(len(m), "hits")
}
