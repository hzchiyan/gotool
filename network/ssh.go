package network

import (
	"bytes"
	"io/ioutil"
	"net"
	"os/exec"
	"time"

	"golang.org/x/crypto/ssh"
)

func LocalCommand(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func SSHSecretLogin(addr, user, secret string) (*ssh.Session, error) {
	var session *ssh.Session
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(secret),
		},
		Timeout: time.Second * 5,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		return session, err
	} else {
		return client.NewSession()
	}
}

func SSHKeyLogin(addr, user, keyFilePath string) (session *ssh.Session, err error) {
	key, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return session, err
	}
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return session, err
	}
	client, err := ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})
	if err != nil {
		return session, err
	}
	return client.NewSession()
}

func RemoteCommand(session *ssh.Session, command string) (string, error) {
	var b bytes.Buffer
	session.Stdout = &b
	defer session.Close()
	if err := session.Run(command); err == nil {
		return b.String(), err
	} else {
		return "", err
	}
}
