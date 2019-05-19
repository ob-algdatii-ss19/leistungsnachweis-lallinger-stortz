pipeline {
    agent none
    stages {
        stage('Build') {
            agent {
                docker { image 'obraun/vss-protoactor-jenkins' }
            }
            steps {
                sh 'go build main.go'
            }
        }
        stage('Test') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }
            steps {
                sh 'cd sortingalgorithms'
                sh 'echo go test -v'
            }
        }
        stage('Lint') {
            agent {
                docker { image 'obraun/vss-jenkins' }
            }   
            steps {
                sh 'echo no lint'
                //sh 'golangci-lint run --enable-all'
            }
        }
        stage('Build Docker Image') {
            agent {
                label 'master'
            }
            steps {
                sh "docker-build-and-push -b ${BRANCH_NAME} -s sortvisualization -f sortvisualization.dockerfile"
                sh "docker-build-and-push -b ${BRANCH_NAME}"
            }
        }
    }
    post {
        changed {
            script {
                if (currentBuild.currentResult == 'FAILURE') { // Other values: SUCCESS, UNSTABLE
                    // Send an email only if the build status has changed from green/unstable to red
                    emailext subject: '$DEFAULT_SUBJECT',
                        body: '$DEFAULT_CONTENT',
                        recipientProviders: [
                            [$class: 'DevelopersRecipientProvider']
                        ], 
                        replyTo: '$DEFAULT_REPLYTO'
                }
            }
        }
    }
}
