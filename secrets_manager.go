package main
import (
   "github.com/aws/aws-sdk-go/aws/session"
   "github.com/aws/aws-sdk-go/service/secretsmanager"
)
func getSecretValue(name string) (string, error){
   sess := session.Must(session.NewSession())
   sm := secretsmanager.New(sess)
   output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &name})
   if err != nil {
      return "", err
   }

   return *output.SecretString, nil
}
