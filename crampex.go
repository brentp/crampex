// Package crampex allows us to read and write cram files in go using a system call to samtools view
// that is then piped to biogo bam reader.
package crampex

import (
	"io"
	"os"
	"os/exec"

	"github.com/biogo/hts/bam"
)

type reader struct {
	io.ReadCloser
	cmd *exec.Cmd
}

func (r *reader) Close() error {
	if err := r.cmd.Wait(); err != nil {
		return err
	}
	return r.ReadCloser.Close()
}

// NewReader returns a bam.Reader from any path that samtools can read.
func NewReader(path string, rd int, fasta string, region string) (*bam.Reader, error) {
	var cmd *exec.Cmd
	if region == "" {
		cmd = exec.Command("samtools", "view", "-T", fasta, "-b", "-u", "-h", path)
	} else {
		cmd = exec.Command("samtools", "view", "-T", fasta, "-b", "-u", "-h", path, region)
	}

	cmd.Stderr = os.Stderr
	pipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	if err = cmd.Start(); err != nil {
		pipe.Close()
		return nil, err
	}
	cr := &reader{ReadCloser: pipe, cmd: cmd}
	return bam.NewReader(cr, rd)
}
