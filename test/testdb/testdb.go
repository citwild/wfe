package testdb

import (
	"bytes"
	"errors"
	"github.com/citwild/wfe/log"
	"github.com/uber-go/zap"
	"net"
	"os/exec"
	"strings"
	"time"
)

const mongoImage = "mvertes/alpine-mongo"

type TestDB struct {
	containerId string
}

func New() *TestDB {
	d := &TestDB{}

	_, err := exec.LookPath("docker")
	if err != nil {
		log.Fatal("Docker not available in path.")
	}

	out, err := exec.Command("docker", "images", "--no-trunc").Output()
	if err != nil {
		log.Fatal("Error running docker to check for image.")
	}

	haveImage := bytes.Contains(out, []byte(mongoImage))
	if !haveImage {
		log.Info("Pulling image...", zap.String("image", mongoImage))
		out, err = exec.Command("docker", "pull", mongoImage).Output()
		if err != nil {
			log.Fatal("Error pulling", zap.Error(err))
		}
	}

	return d
}

func (d *TestDB) Start() error {
	out, err := exec.Command("docker", "run", "-d", "-p", "27017:27017", mongoImage).Output()
	if err != nil {
		return err
	}

	d.containerId = strings.TrimSpace(string(out))
	if d.containerId == "" {
		errors.New("Unexpected empty output from `docker run`")
	}

	const timeout = 10 * time.Second
	start := time.Now()
	for true {
		c, err := net.Dial("tcp", ":27017")
		if err == nil {
			c.Close()
			break
		}
		if time.Since(start) > timeout {
			d.Close()
			return errors.New("Timed out waiting for db to start")
		}
		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (d *TestDB) Close() {
	exec.Command("docker", "kill", d.containerId).Run()
}
