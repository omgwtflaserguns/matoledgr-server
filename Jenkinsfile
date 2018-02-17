node {
    try{
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/gopath/src/github.com/omgwtflaserguns/matomat-server") {

            // Install the desired Go version
            def root = tool name: 'go 1.9.4', type: 'go'

            // Export environment variables pointing to the directory where Go was installed
            withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/gopath"]) {

                stage('Checkout'){
                    echo 'Checking out from ' + env.BRANCH_NAME
                    checkout scm
                }

                stage('Get Dependencies'){
                    echo 'Getting dependencies'

                    echo "GOROOT: $GOROOT"
                    echo "GOPATH: $GOPATH"
                    echo "PATH: $PATH"

                    sh 'go version'
                    sh 'go get -u github.com/golang/dep/...'

                    sh "$GOPATH/bin/dep ensure"

                    sh 'go vet'
                }

                stage('Build'){
                    echo 'Building Executable'
                    sh 'go build'
                }

                stage('Deploy'){
                    if (env.BRANCH_NAME == 'master') {
                        echo 'on master branch, deploy into test...'
                        sh 'sudo ./deployTest.sh'
                    }
                    else if (env.BRANCH_NAME == 'release') {
                        echo 'on release branch, deploy into production...'
                        sh 'sudo /var/lib/jenkins/deployProd.sh'
                    }
                    else {
                        echo 'not on master or release, no deployment for you!'
                    }
                }
            }
        }
    }catch (e) {
        currentBuild.result = "FAILED"
    } finally {
        notifyBuild(currentBuild.result)
    }
}

def notifyBuild(String buildStatus = 'STARTED') {
  // build status of null means successful
  buildStatus =  buildStatus ?: 'SUCCESSFUL'

  // Default values
  def colorCode = '#FF0000'
  def subject = "${buildStatus}: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]'"
  def summary = "${subject} <${env.BUILD_URL}|Job URL> - <${env.BUILD_URL}/console|Console Output>"

  if (buildStatus == 'STARTED') {
    colorCode = '#FFFF00'
  } else if (buildStatus == 'SUCCESSFUL') {
    colorCode = '#00FF00'
  } else {
    colorCode = '#FF0000'
  }

  // Send notifications
  slackSend (color: colorCode, message: summary)
}
