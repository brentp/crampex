package crampex

import (
	"io"
	"log"
	"testing"
)

func TestCramReader(t *testing.T) {

	cr, err := NewReader("/home/brentp/src/bcftools/test/mpileup/mpileup.1.cram",
		1, "/home/brentp/src/bcftools/test/mpileup/mpileup.ref.fa", "")
	if err != nil {
		t.Fatal(err)
	}
	var a int
	for {
		b, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		a = b.Start()
	}
	log.Println(a)
}

func TestCramReaderRegion(t *testing.T) {

	cr, err := NewReader("test/mpileup.1.cram",
		1, "test/mpileup.ref.fa", "17:3992-3998")
	if err != nil {
		t.Fatal(err)
	}
	var k int
	for {
		b, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if b.End() < 2992 {
			t.Fatalf("found end below requested region: %d", b.End())
		}
		_ = b.Start()
		k++
	}
	if k != 16 {
		t.Fatalf("expected 16 rows, found %d", k)
	}
}
