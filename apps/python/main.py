import aws_cdk

app = aws_cdk.App()
aws_cdk.Stack(app, "TestStack")
app.synth()
