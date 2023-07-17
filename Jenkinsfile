pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh  '''
                        echo "Multiline shell steps works too"
                    '''
            }
        }
        stage('Upload to AWS') {
              steps {
                    dir('/var/lib/jenkins/workspace') {
                    pwd();
                    withAWS(region:'ap-south-1',credentials:'sajan-aws-jenkins-id') {                         
                         sh 'echo "Uploading content with AWS creds"'
                         s3Upload(bucket:"sjn-golang-bucket-01", workingDir:'GoLang', includePathPattern:'**/*');                         
                         }
                    }
              }
         }
        stage('Install awscli') {
            steps {
                sh 'aws --version'
            }
        }
        stage('AWS Configure') {
            steps {
                withCredentials([
                    [
                        $class: 'AmazonWebServicesCredentialsBinding',
                        credentialsId: 'sajan-aws-jenkins-id',
                        accessKeyVariable: 'AWS_ACCESS_KEY_ID',
                        secretKeyVariable: 'AWS_SECRET_ACCESS_KEY'
                    ]
                ])
                {                    
                    sh 'aws s3 ls'
                    sh 'aws cloudformation create-stack --stack-name goStack --template-body file://goTemp.yaml --capabilities CAPABILITY_NAMED_IAM'
                }
            }           
        }          
        
    }
}
