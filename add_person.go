package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"io"
	"strings"

	// "io/ioutil"
	// "log"
	// "os"
	// "strings"
	pb "elfiyang16/go-protobuf/protoc"
	// "google.golang.org/protobuf/proto"
)

func promptForAddress(r io.Reader) (*pb.Person, error) {
	p := &pb.Person{}
	rd := bufio.NewReader(r)

	fmt.Print("Enter person ID number:")
	if _, err := fmt.Fscanf(rd, "%d\n", &p.Id); err != nil {
		return p, err
	}
	fmt.Print("Enter name:")
	name, err := rd.ReadString('\n')
	if err != nil {
		return p, err
	}
	p.Name = strings.TrimSpace(name)

	for {
		fmt.Print("Enter a phone number (or leave blank to finish): ")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}
		pn := &pb.Person_PhoneNumber{
			Number: phone,
		}
		fmt.Print("Is this a mobile, home, or work phone? ")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return p, err
		}
		ptype = strings.TrimSpace(ptype)
		// A proto enum results in a Go constant for each enum value.
		switch ptype {
		case "mobile":
			pn.Type = pb.Person_MOBILE
		case "home":
			pn.Type = pb.Person_HOME
		case "work":
			pn.Type = pb.Person_WORK
		default:
			fmt.Printf("Unknown phone type %q.  Using default.\n", ptype)
		}

		p.Phones = append(p.Phones, pn)
	}

}

func main() {

}
