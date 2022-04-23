pipeline {
    agent any
    tools {
        go 'go1.17'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0
    }
    stages {        
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
                sh 'go get -u golang.org/x/lint/golint'
            }
        }
        
        stage('Build') {
            steps {
                echo 'Building docker image'
                sh 'docker build -t mentoring-app:latest'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'go vet .'
                    echo 'Running linting'
                    sh 'golint .'
                }
            }
        }
        
    }
    post {
    }  
}
