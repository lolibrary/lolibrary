package portforward

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"syscall"

	"github.com/logrusorgru/aurora"
	"github.com/phayes/freeport"
)

type background struct {
	exit     chan bool
	finished chan error
	cmd      *exec.Cmd
}

func (b *background) Start() error {
	if err := b.cmd.Start(); err != nil {
		return err
	}

	go b.Monitor()
	go b.Wait()

	return nil
}

func (b *background) Monitor() {
	// block until we need to exit or the program ends prematurely
	<-b.exit

	if err := b.cmd.Process.Signal(syscall.SIGTERM); err != nil {
		// we hit a bad state here, just exit and log.
		log.Fatalf("Failed to kill port-forward process: %v\n", err)
	}
}
func (b *background) Wait() {
	err := b.cmd.Wait()

	b.finished <- err
}

func (b *background) WaitC() <-chan error {
	return b.finished
}

func (b *background) Close() error {
	fmt.Println("📉  Shutting down port-forward process.")

	b.exit <- true

	err := <-b.WaitC()

	close(b.finished)
	close(b.exit)

	return err
}

// Enable creates a backend process to kubectl port-forward the edge proxy to localhost.
// It returns a port, which the proxy is running on, and a Closer, which should be closed to shut down the proxy.
func Enable() (uint16, io.Closer) {
	remotePort := 80
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatalf("😩 Could not get an open port to forward to: %v\n", err)
	}

	ports := fmt.Sprintf("%d:%d", port, remotePort)
	c := exec.Command("kubectl", "port-forward", "service/edge-proxy-internal", ports)
	bg := &background{
		cmd:      c,
		exit:     make(chan bool),
		finished: make(chan error),
	}

	fmt.Printf("↔️  Starting a port forward for %v on port %v\n", aurora.Green("service/edge-proxy-internal"), aurora.Green(port))
	if err := bg.Start(); err != nil {
		log.Fatalf("❌ Failed to start port forward: %v\n", err)
	}

	return uint16(port), bg
}
