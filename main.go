package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gopaytech/internal-access-helper/config"
	"github.com/gopaytech/internal-access-helper/settings"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ArgoCDSecret struct {
	CACert string `json:"ca.crt"`
	Token  string `json:"token"`
}

func main() {
	settings, err := settings.NewSettings()
	if err != nil {
		log.Fatal(err)
	}

	k8s, err := config.LoadKubernetes()
	if err != nil {
		log.Fatal(err)
	}

	client := k8s.Client()

	secret, err := client.CoreV1().Secrets(settings.ArgoCDNamespace).Get(context.Background(), settings.ArgoCDManagerSecretName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get namespace %s secret %s: %s\n", settings.ArgoCDNamespace, settings.ArgoCDManagerSecretName, err.Error())
	}

	argoCDSecret := ArgoCDSecret{
		CACert: base64.StdEncoding.EncodeToString(secret.Data["ca.crt"]),
		Token:  base64.StdEncoding.EncodeToString(secret.Data["token"]),
	}

	argoCDSecretByte, err := json.Marshal(argoCDSecret)
	if err != nil {
		log.Printf("failed to marshal argoCDSecret")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(argoCDSecretByte)
	})

	log.Fatal(http.ListenAndServe(":"+settings.HTTPPort, nil))
}
