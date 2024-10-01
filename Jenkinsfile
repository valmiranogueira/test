pipeline {
    agent any

    environment {
        // Define any environment variables here
        GIT_REPO_URL = 'git@github.com:valmiranogueira/test.git'
        BRANCH_NAME = 'main'
    }

    stages {
        stage('Test') {
            steps {
                echo 'Running tests...'
                // Replace with your test command, e.g., Maven, Gradle, etc.
                sh 'echo testing'
            }
        }

        stage('Deploy') {
            steps {
                sh 'echo hi'
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
