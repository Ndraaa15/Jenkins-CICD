pipeline {
  agent any
  stages {
    stage('Clone Repository') {
      steps {
          git branch : 'main', url:  'https://github.com/Ndraaa15/Jenkins-CICD.git'
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
        agent {
          docker { 
            image 'golang:1.22' 
            args '-r GOCACHE=/tmp/go-cache'
          }
        }


        script {
          dockerImage.inside {
            sh 'go mod download'
            sh 'go test -v .'
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
