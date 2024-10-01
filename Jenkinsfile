pipeline {
    agent any

    environment {
        // Define any environment variables here
        GIT_REPO_URL = 'git@github.com:valmiranogueira/test.git'
        BRANCH_NAME = 'main'
    }

    stages {
        stage('Build') {
            steps {
                echo 'Building Docker image...'
                script {
                    // Build the Docker image
                    def image = docker.build("test:latest")
                }
            }
        }

        stage('Deploy') {
            steps {
                echo 'Running Docker container...'
                script {
                    // Run the Docker container
                    docker.image("test:latest").run('-p 8000:8000')
                }
            }
        }
    }

    post {
        success {
            echo 'Pipeline succeeded!'
        }
        failure {
            echo 'Pipeline failed!'
            // Optionally notify or take action on failure
        }
        always {
            echo 'Cleaning up...'
            // Perform any cleanup actions, if necessary
        }
    }
}
