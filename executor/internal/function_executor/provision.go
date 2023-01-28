package function_executor

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/orted-org/isdn/util"
)

func (fe *FunctionExecutor) Provision() error {

	log.Println("provisioning resources for request id:", fe.params.RequestID)

	// create dir
	fe.workingDirectory = path.Join(BASE_DIR, fe.params.RequestID)
	log.Println("creating directory at path", fe.workingDirectory)
	err := os.MkdirAll(fe.workingDirectory, os.ModePerm)
	if err != nil {
		return err
	}

	extension := fe.langHandler.GetExtension(fe.params.Language)
	templateExtension := extension
	if fe.params.IsZip {
		extension = "zip"
	}

	if fe.params.IsZip {
		// unzip the zip file from staging to main area

		tempZipPath := path.Join(BASE_DIR, "temp_zips", fmt.Sprintf("%s.%s", fe.params.RequestID, extension))
		log.Println("creating new zip file at", tempZipPath)

		b, err := io.ReadAll(fe.params.Code)
		if err != nil {
			return util.MergeErrors(errors.New("could not read zip file"), err)
		}

		// putting the file in temp area
		err = fe.fileManger.Put(tempZipPath, b)
		if err != nil {
			return util.MergeErrors(errors.New("could not save zip file"), err)
		}

		// moving the file to main area
		err = unzip(tempZipPath, fe.workingDirectory)
		if err != nil {
			return util.MergeErrors(errors.New("could not unzip file"), err)
		}
		log.Println("unzipped file at", fe.workingDirectory)
	} else {
		// just move code from staging to main area
		fe.codeFilePath = path.Join(fe.workingDirectory, fmt.Sprintf("code.%s", fe.langHandler.GetExtension(fe.params.Language)))
		log.Printf("creating FILE=%s in DIR=%s\n", fe.codeFilePath, fe.workingDirectory)

		// saving file in system
		b, err := io.ReadAll(fe.params.Code)
		if err != nil {
			return errors.New("could not read zip file")
		}
		err = fe.fileManger.Put(fe.codeFilePath, b)
		if err != nil {
			log.Printf("could not save file")
			return err
		}
	}

	// copy template to working directory
	template, err := os.ReadFile(path.Join(TEMPLATES_DIR, templateExtension, "template."+templateExtension))
	if err != nil {
		return err
	}

	entryFile, err := os.Create(path.Join(fe.workingDirectory, "code."+templateExtension))
	if err != nil {
		return err
	}
	defer entryFile.Close()
	entryFile.Write(template)

	// copy input to working directory
	inputFile, err := os.Create(path.Join(fe.workingDirectory, "input.in"))
	if err != nil {
		return err
	}
	defer inputFile.Close()
	inputFile.Write([]byte(fe.params.Input))

	// create output file
	outputFile, err := os.Create(path.Join(fe.workingDirectory, "output.out"))
	if err != nil {
		return err
	}
	defer outputFile.Close()

	return nil
}

func (fe *FunctionExecutor) Clean() error {
	log.Println("removing resources for request id:", fe.params.RequestID)
	err := os.RemoveAll(fe.workingDirectory)
	if err != nil {
		return err
	}

	if fe.params.IsZip {
		err := os.RemoveAll(path.Join(BASE_DIR, "temp_zips", fe.params.RequestID+".zip"))
		if err != nil {
			return err
		}
	}
	return nil
}

func unzip(zipPath string, dst string) error {

	archive, err := zip.OpenReader(zipPath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			return errors.New("invalid file path")
		}
		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	return nil
}
