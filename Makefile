build:
	GOOS=darwin GOARCH=amd64 go build -o .build/macos/ec2metadata github.com/scottbrown/ec2metadata
	GOOS=linux GOARCH=amd64 go build -o .build/linux/ec2metadata github.com/scottbrown/ec2metadata

upload:
	scp -i ~/Downloads/scott.pem .build/linux/ec2metadata ubuntu@ec2-52-91-185-200.compute-1.amazonaws.com:

test:
	go test ./...
