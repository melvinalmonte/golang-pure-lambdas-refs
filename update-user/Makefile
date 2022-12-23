build:
	GOARCH=amd64 GOOS=linux go build main.go
package:
	zip my-go-lambda.zip main
deploy:
	aws lambda update-function-code --function-name go-tester-update --zip-file fileb://my-go-lambda.zip --no-cli-pager
clean:
	rm main
	rm my-go-lambda.zip
all:
	make build
	make package
	make deploy
	make clean
