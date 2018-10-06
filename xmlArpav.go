package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// http://www.arpa.veneto.it/bollettini/meteo/h24/img12/Graf_231.htm?sens=TEMP
// http://www.arpa.veneto.it/bollettini/meteo/h24/img12/0231.xml

const Tformin = "200601021504"
const Tformout = "02/01/2006 15:04"

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usare:%s <numeroStaz es.231>\n", os.Args[0])
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		xmlStru := new(Conten)
		url := "http://www.arpa.veneto.it/bollettini/meteo/h24/img12/0" + os.Args[i] + ".xml"
		connt, err := http.Get(url)
		if err != nil {
			fmt.Printf("errore %s", err.Error())
			os.Exit(1)

		}
		exit, err := ioutil.ReadAll(connt.Body)
		if err != nil {
			fmt.Printf("errore %s", err.Error())
			os.Exit(1)

		}
		/*
		   d := xml.NewDecoder(reader)
		   d.CharsetReader = CharsetReader
		   err := d.Decode(&dst)
		*/

		Undec := xml.NewDecoder(bytes.NewReader(exit))
		Undec.CharsetReader = CharsetReader
		errUnM := Undec.Decode(xmlStru)

		// errUnM := xml.Unmarshal(exit, xmlStru)

		if errUnM != nil {
			fmt.Printf("Errore %s", errUnM.Error())

		} else {
			for _, st := range xmlStru.Stazioni {
				for _, sens := range st.Sensori {
					for _, SingD := range sens.Valori {
						dataOk, _ := time.Parse(Tformin, SingD.Dataora)
						fmt.Printf("\nStaz:%s;Unità di Misura %s;tipo %s;Data;%s;valore;%f", st.Nome, sens.Un, sens.Typ, dataOk.Format(Tformout), SingD.Valore)
					}

				}
			}
		}

	}
	//  	fmt.Printf("HELLO WORLD %s\n", os.Args[i])

	fmt.Println("")
}
