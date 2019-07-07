pipeline {
    agent any

    stages {
        stage('Build Dockerfile') {
            steps {
                /* `docker build` returns non-zero on test failures,
                *  Command builds the Dockerfile checking the syntax and functional errors of the Golang project
                */
                sh 'docker build ./dldata/dockerfiles/darklist_blacklist'
            }
        }
        stage('Remove Docker Images') {
            steps {
                /* 'docker rmi' returns non-zero on test failures,
                *  Command deletes all docker images to clean the Jenkins server
                */
                sh 'docker system prune -f'
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
                            [$class: 'CulpritsRecipientProvider'],
                            [$class: 'DevelopersRecipientProvider'],
                            [$class: 'RequesterRecipientProvider']
                        ], 
                        replyTo: '$DEFAULT_REPLYTO',
                        to: 'liamjosephkeenan@gmail.com'
                }
            }
        }
    }
}
