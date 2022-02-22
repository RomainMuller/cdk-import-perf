const cdk = require('aws-cdk-lib');

const app = new cdk.App();
new cdk.Stack(app, 'TestStack');
app.synth();
