pipeline {
  agent any
  stages {
    stage('Clone Repository') {
      steps {
        script{
          sh 'echo Clone Repository'
        }

        git branch: 'main', url: 'https://github.com/Ndraaa15/Jenkins-CICD.git'
      }
    }

    stage('Build Docker Image') {
      steps {
        script {
          sh 'echo Build Image'
          dockerImage = docker.build("hello-world-app")
        }
      }
    }

    stage('Run Unit Tests') {
      steps {
        script {
          docker.image('golang:1.22').inside('-v $WORKSPACE:/go/src/app -w /go/src/app') {
            sh 'echo Running Test'
            sh 'go mod download'
            sh 'go test -v ./...'
          }
        }
      }
    }

    stage('Deploy to Production') {
      steps {
        script {
          sh 'echo Deploying to Production'
          dockerImage.run('-p 8000:8000')
        }
      }
    }
  }
}
