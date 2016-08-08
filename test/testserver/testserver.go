package testserver

import (
	"crypto/tls"
	"errors"
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/log"
	"github.com/citwild/wfe/test/testdb"
	"github.com/uber-go/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type TestServer struct {
	serveCmd *exec.Cmd
	tempDir  string
	testDB   *testdb.TestDB
}

func New() *TestServer {
	s := &TestServer{}

	dir, err := ioutil.TempDir("", "testserver")
	if err != nil {
		log.Fatal("Failed to create temp dir.", zap.Error(err))
	}
	s.tempDir = dir

	keyFile := filepath.Join(s.tempDir, "testserver.key")
	err = ioutil.WriteFile(keyFile, []byte(localhostKey), 0600)
	if err != nil {
		log.Fatal("Failed to write private key.", zap.Error(err))
	}

	certFile := filepath.Join(s.tempDir, "testserver.crt")
	err = ioutil.WriteFile(certFile, []byte(localhostCert), 0600)
	if err != nil {
		log.Fatal("Failed to write cert.", zap.Error(err))
	}

	cmdPath := os.Getenv("GOPATH")
	if cmdPath != "" {
		cmdPath = filepath.Join(cmdPath, "bin")
	}
	cmdPath = filepath.Join(cmdPath, "wfe")

	s.testDB = testdb.New()

	cmd := exec.Command(cmdPath)
	cmd.Args = append(cmd.Args, "serve")
	cmd.Args = append(cmd.Args, "--tls-key="+keyFile)
	cmd.Args = append(cmd.Args, "--tls-cert="+certFile)
	s.serveCmd = cmd

	return s
}

func (s *TestServer) Start() error {
	err := s.testDB.Start()
	if err != nil {
		s.testDB.Close()
		return err
	}

	err = s.serveCmd.Start()
	if err != nil {
		return err
	}

	// disable security check for self-signed certificate
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	const timeout = 10 * time.Second
	start := time.Now()
	for true {
		_, err := client.Get("https://localhost:8443")
		if err == nil {
			break
		}
		if time.Since(start) > timeout {
			s.Close()
			return errors.New("Timed out waiting for server to start")
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (s *TestServer) NewClient() (*api.Client, error) {
	cred := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.Dial("localhost:8443", grpc.WithTransportCredentials(cred))
	if err != nil {
		return nil, err
	}
	return api.NewClient(conn), nil
}

func (s *TestServer) Close() {
	err := s.serveCmd.Process.Kill()
	if err != nil {
		log.Fatal("Failed to kill serve process.", zap.Error(err))
	}

	err = os.RemoveAll(s.tempDir)
	if err != nil {
		log.Fatal("Failed to delete temp dir.", zap.Error(err))
	}

	s.testDB.Close()
}

const localhostKey = `
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDstcSQUxfuq71pBmWGZyHJO9HhlVmxUS0mpzmgTKlYrTC7asp4
HlfP8apJHjHmoxsrQYW8N527BQZdD+85zI9gu5uU/8nbm3o0Y2jhUlsTodoUpXah
sVYw5e+qJx3t9Sfdho/1pnNnzaqGJnU2iU9Cg7h9K9TDJ7MyVfk6E9UAkwIDAQAB
AoGBAKIGLTzI07nXrjfEXBvqXfL7UkdBznoT4X7yufwiXUxIh/HRQDOGOX8poPCZ
jXicLo2mNN9HhlysLNXogUwj3AL+h+Vu53Tg5T2tMSP7frYGnACSJVHK3JXFbwvr
WyJdKnAJ0RwbCjenXCZ7XorskZ8VrcLFvswNpRVwQ4crMUZhAkEA9oPIaFmTpXu9
UwZk8hyZvRuGmzicqK15Ssv5q5gegQip3g/YNXOYQ2DRBQn2jtMijQMyeMHSEG9Y
fsY92ZtvkQJBAPXRaAiglxPbWpFQJ5qdWSVxHOnywlZ3/Tv1DTiPJH0MZB/VeLtT
/jZqqMofUdWU7yVyj1lcPZxFnpuqlF8cY+MCQEvdaKq8jpTKDQzlFeHw7Vtmgjl4
5bV/lalwNskZSqH5UZW2mJpylbR+sjTzyP5RefudtUW2mHhYXAq/5b73eBECQBST
uO+v8bWZ0RUE7qvErCe0NkVnKtluJeaU9sxbPySwmtkHV5nmyArqFsMEqCcG1pX9
5v4F+KpSMZq6Rr1HdR0CQQDyt3KQb9CkUw0nSA/wcXOgYKHsGnXqYMn9P6i1E7i7
vedE+GlspO1E9LjV4Y5lk0mNwzUENbfWEeB6BWjZN6sq
-----END RSA PRIVATE KEY-----`

const localhostCert = `
-----BEGIN CERTIFICATE-----
MIICEjCCAXugAwIBAgIRAK62Q0LiITS38o7E0yolCoIwDQYJKoZIhvcNAQELBQAw
EjEQMA4GA1UEChMHQWNtZSBDbzAgFw03MDAxMDEwMDAwMDBaGA8yMDg0MDEyOTE2
MDAwMFowEjEQMA4GA1UEChMHQWNtZSBDbzCBnzANBgkqhkiG9w0BAQEFAAOBjQAw
gYkCgYEA7LXEkFMX7qu9aQZlhmchyTvR4ZVZsVEtJqc5oEypWK0wu2rKeB5Xz/Gq
SR4x5qMbK0GFvDeduwUGXQ/vOcyPYLublP/J25t6NGNo4VJbE6HaFKV2obFWMOXv
qicd7fUn3YaP9aZzZ82qhiZ1NolPQoO4fSvUwyezMlX5OhPVAJMCAwEAAaNmMGQw
DgYDVR0PAQH/BAQDAgKkMBMGA1UdJQQMMAoGCCsGAQUFBwMBMA8GA1UdEwEB/wQF
MAMBAf8wLAYDVR0RBCUwI4IJbG9jYWxob3N0hwR/AAABhxAAAAAAAAAAAAAAAAAA
AAABMA0GCSqGSIb3DQEBCwUAA4GBAEkAPtaUVRNMUovPAIe8wrAE/hTBfqNlbXk8
DJHP8luLVs5GzG6OvouLC9dvnVdCb380QmqHnXagTzIYObf850AYACq2GXCIP0Qu
ictwafmQSBI281Xxk2MGeQm6smiTej8fomSqUUIqdEuy1+qhnSAIOxBZWhJbISrf
lIyC/HF0
-----END CERTIFICATE-----`
