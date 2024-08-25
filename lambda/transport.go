package lambda

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var SourceClient *lambda.Client
var TargetClient *lambda.Client

const downloadPath = "./downloads"

func DownLoadSourceLambdaCode(client *lambda.Client, functionName string) error {
	log := Logger
	// Get the source lambda function code
	getFunctionOutput, err := client.GetFunction(context.TODO(), &lambda.GetFunctionInput{
		FunctionName: &functionName,
	})
	if err != nil {
		panic("get source lambda function code error, " + err.Error())
	}
	// Download the ZIP file from the provided location
	location := getFunctionOutput.Code.Location
	resp, err := http.Get(*location)
	if err != nil {
		log.Error("Error downloading the ZIP file:", "error", err)
		return err
	}
	defer resp.Body.Close()

	sourceFileName := filepath.Join(downloadPath, functionName+".zip")
	file, err := os.Create(sourceFileName)
	if err != nil {
		log.Error("Error creating file:", "error", err)
		return err
	}
	defer file.Close()

	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		log.Error("Error saving the ZIP file:", "error", err)
		return err
	}
	log.Info("Downloaded the ZIP file from the source lambda function", "filename", sourceFileName)
	return nil
}

func UploadTargetLambdaCode(client *lambda.Client, functionName string) error {
	log := Logger
	// Read the ZIP file
	fileToUpload, err := os.ReadFile(filepath.Join(downloadPath, functionName+".zip"))
	if err != nil {
		log.Error("Error opening the ZIP file:", "error", err)
		return err
	}

	updateFunctionCodeOutput, err := client.UpdateFunctionCode(context.TODO(), &lambda.UpdateFunctionCodeInput{
		FunctionName: &functionName,
		ZipFile:      fileToUpload,
	})
	if err != nil {
		log.Error("Error updating Lambda function code:", "error", err)
		return err
	}
	log.Info("Uploaded the ZIP file to the target lambda function", "functionName", functionName, "revisionId", *updateFunctionCodeOutput.RevisionId)
	return nil
}
