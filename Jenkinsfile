// todo

String image    = env.JOB_NAME.split('/')[1].toLowerCase()
String registry = "harik8/$image"
String tag      = "latest"

pipeline {
    agent {
    kubernetes {
      	cloud 'kubernetes'
      	label 'todo'
      	defaultContainer 'jnlp'
      	yamlFile 'pod.yaml'
      }
    }

    stages {
        stage("publish") {
            steps {
                container(name: 'kaniko', shell: '/busybox/sh') {
                     withEnv(['PATH+EXTRA=/busybox:/kaniko']) {
                      sh """#!/busybox/sh
                        /kaniko/executor --context=$workspace --destination $registry:$tag
                      """
                     }
                }
            }
        }
    }
}
