package update

import (
	"bytes"
	"context"
	"io/ioutil"
	"path"

	"github.com/codeclysm/extract/v3"
)

func ExtractTarGz(gzipfile string, destiny string) (files []string, err error) {
	data, err := ioutil.ReadFile(gzipfile)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(data)
	files = make([]string, 0, 10)
	var shift = func(p string) string {
		files = append(files, path.Join(destiny, p))
		return p
	}

	err = extract.Archive(context.Background(), buffer, destiny, shift)

	return

	// gzipStream, err := os.Open(gzipfile)
	// if err != nil {
	// 	return nil, err
	// }

	// if stat, err := os.Stat(destiny); err != nil || !stat.IsDir() {
	// 	return nil, fmt.Errorf("Destiny %s is not a directory", destiny)
	// }

	// uncompressedStream, err := gzip.NewReader(gzipStream)
	// if err != nil {
	// 	return nil, err
	// }

	// tarReader := tar.NewReader(uncompressedStream)
	// files = make([]string, 0, 10)
	// var header *tar.Header
	// for header, err = tarReader.Next(); err != nil; header, err = tarReader.Next() {
	// 	destinyFile := path.Join(destiny, header.Name)
	// 	switch header.Typeflag {
	// 	case tar.TypeDir:
	// 		if err := os.Mkdir(destinyFile, 0755); err != nil {
	// 			return nil, fmt.Errorf("ExtractTarGz: Mkdir() failed: %w", err)
	// 		}
	// 	case tar.TypeReg:
	// 		outFile, err := os.Create(destinyFile)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("ExtractTarGz: Create() failed: %w", err)
	// 		}

	// 		if _, err := io.Copy(outFile, tarReader); err != nil {
	// 			// outFile.Close error omitted as Copy error is more interesting at this point
	// 			outFile.Close()
	// 			return nil, fmt.Errorf("ExtractTarGz: Copy() failed: %w", err)
	// 		}
	// 		if err := outFile.Close(); err != nil {
	// 			return nil, fmt.Errorf("ExtractTarGz: Close() failed: %w", err)
	// 		}
	// 		files = append(files, destinyFile)
	// 	default:
	// 		return nil, fmt.Errorf("ExtractTarGz: uknown type: %b in %s", header.Typeflag, header.Name)
	// 	}
	// }
	// if err != io.EOF {
	// 	return nil, fmt.Errorf("ExtractTarGz: Next() failed: %w", err)
	// }
	// return files, err
}
