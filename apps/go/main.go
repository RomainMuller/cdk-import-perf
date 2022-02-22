package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	app := awscdk.NewApp(nil)
	awscdk.NewStack(app, jsii.String("TestStack"), nil)
	app.Synth(nil)
}
