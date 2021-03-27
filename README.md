# Lambda text detection

Scanning images and detecting text in it, based on [Rekognition text detection](https://docs.aws.amazon.com/rekognition/latest/dg/text-detection.html) to detect
text from images supplied.

## Detecting text in images flow

![](https://user-images.githubusercontent.com/20343969/112707937-b6013f00-8e8d-11eb-97d9-df2bcaf4c78f.png)

## Interesting Info

- [Lambda Getting started](https://docs.aws.amazon.com/lambda/latest/dg/getting-started.html)
- [Setting the Lambda IAM Role](https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html)
- [Lambda in Golang](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [AWS sdk for Golang](https://aws.github.io/aws-sdk-go-v2/docs/getting-started/)
- [Deploying your lambda in Golang](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
- [Amazon Rekognition](https://aws.amazon.com/rekognition/)



## To deploy your lambda from local:
- Make sure the lambda already exists

- Update the name of your lambda in ```./scripts/deploy.sh```

- run: ```sh ./scripts/deploy.sh``` 


