pipeline {
  agent any
  stages {
    stage('Clone Repository') {
      steps {
        git 'https://github.com/Ndraaa15/Jenkins-CICD.git'
        }
    }
    
    stage('Build Docker Image') {\
      steps {
        script {
          dockerImage = docker.build("hello-world-app")
        }
      }
    }
    
    stage('Run Unit Tests') {
      steps {
        script {
          dockerImage.inside {
            sh 'go mod download'
            sh 'go test main_test.go'
          }
        }
      }
    }
    
    stage('Deploy to Production') {
      steps {
        script {
          dockerImage.run('-p 8000:8000')
        }
      }
    }
  }
}
