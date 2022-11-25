package function_executor

import (
	"fmt"
	"log"
	"os"
	"path"
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

	// forming filename
	fe.codeFilePath = path.Join(fe.workingDirectory, fmt.Sprintf("code.%s", fe.langHandler.GetExtension(fe.params.Language)))
	log.Printf("creating FILE=%s in DIR=%s\n", fe.codeFilePath, fe.workingDirectory)

	// saving file in system
	err = fe.fileManger.Put(fe.codeFilePath, []byte(fe.params.Code))
	if err != nil {
		log.Printf("could not save file")
		return err
	}

	// handling file input
	if fe.params.Input.InputFileName != "" {
		// create file for input
		err = fe.fileManger.Put(path.Join(fe.params.RequestID, fe.params.Input.InputFileName), []byte(fe.params.Input.File))
		if err != nil {
			return err
		}
	}

	return nil
}

func (fe *FunctionExecutor) Clean() error {
	log.Println("removing resources for request id:", fe.params.RequestID)
	err := os.RemoveAll(fe.workingDirectory)
	if err != nil {
		return err
	}
	return nil
}
