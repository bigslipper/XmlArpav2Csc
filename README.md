##XmlArpav2Csv

This a simple program written in Go language that read a XML file created by Veneto Regional Agency for Protection on Environment 
(ARPAV) and print to stdout all the field divided by a simple semicolon

Program have only one parameter, the number of station without zero ( three cypher ) and read from Arpav website  the xml file. 

to build 

'''
#>go build
'''

### The structure
The xml file structure is decoded in _struct.go_ 


the principle node

'''
type Conten struct  {
    XMLNAME     xml.Name   `xml:"CONTENITORE"`
    Forn        string     `xml:"FORNITORE"`
    IstRun	string     `xml:"ISTANTERUN"`
    Inizio      string     `xml:"INIZIO"`
    Fine        string     `xml:"FINE"`
    Stazioni    []Stazione `xml:"STAZIONE"`
}
'''

this is stations structure

'''
type Stazione struct {
         XMLNAME xml.Name   `xml:"STAZIONE"`
         IDSt  int          `xml:"ISTAZ"`
         Nome  string       `xml:"NOME"`
	 PosX  float64      `xml:"X"`
	 PosY  float64      `xml:"Y"`
         Quota  int         `xml;"QUOTA"`
  	 TipoSt string      `xml:"TIPOSTAZ"`
         Prov   string	    `xml:"PROVINCIA"`
	 Comune string	    `xml:"COMUNE"`
	 DataAtt string	    `xml:"ATTIVAZIONE"`
	 Sensori []Sensore  `xml:"SENSORE"`
}
'''

