package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/unicode/norm"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v2"
)

func main() {
	_ = norm.NFC.String("security-pipeline")
	_ = ssh.ClientConfig{User: "demo", HostKeyCallback: ssh.InsecureIgnoreHostKey()}
	_ = jwt.New(jwt.SigningMethodHS256)
	_ = websocket.Upgrader{}
	_ = retryablehttp.NewClient()
	_ = prometheus.NewCounter(prometheus.CounterOpts{Name: "demo_total", Help: "demo"})
	_ = &cobra.Command{Use: "vulnerable-demo"}
	_ = grpc.NewServer()
	_, _ = yaml.Marshal(map[string]string{"test": "security"})
	_ = http.MethodGet
	fmt.Println("intentionally vulnerable dependency test project")
}
