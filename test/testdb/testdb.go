package testdb

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/citwild/wfe/log"
	"github.com/uber-go/zap"
	"gopkg.in/mgo.v2"
	"net"
	"os/exec"
	"strings"
	"time"
)

const mongoImage = "mvertes/alpine-mongo"

type TestDB struct {
	Address     string
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
	out, err := exec.Command("docker", "run", "-d", "-P", mongoImage).Output()
	if err != nil {
		return fmt.Errorf("Failed on `docker run`: %v", err)
	}

	d.containerId = strings.TrimSpace(string(out))
	if d.containerId == "" {
		return errors.New("Unexpected empty output from `docker run`")
	}

	out, err = exec.Command("docker", "port", d.containerId, "27017/tcp").Output()
	if err != nil {
		d.Close()
		return err
	}

	d.Address = strings.TrimSpace(string(out))
	if d.Address == "" {
		return errors.New("Unexpected empty output from `docker port`")
	}

	const timeout = 10 * time.Second
	start := time.Now()
	for true {
		c, err := net.Dial("tcp", d.Address)
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

func (d *TestDB) NewSession() (*mgo.Session, error) {
	s, err := mgo.Dial(d.Address)
	if err != nil {
		return nil, err
	}
	return s, nil
}
