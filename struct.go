package main
import 	"encoding/xml"

/*
type Anagrafe struct {
	XMLNAME xml.Name `xml:"anagrafica"`
	Codice  string   `xml:"codice"`
	Nome    string   `xml:"nome"`
	NomeB   string   `xml:"nomebreve"`
	Quota   int      `xml:"quota"`
	Lat     float32  `xml:"latitudine"`
	Long    float32  `xml:"longitudine"`
	Est     float32  `xml:"est"`
	North   float32  `xml:"north"`
}

type AllAnag struct {
	XMLNAME xml.Name   `xml:"ArrayOfAnagrafica"`
	attr1   string     `xml:"xmlns:xsi,attr"`
	attr2   string     `xml:"xmlns:xsd,attr"`
	attr3   string     `xml:"xmlns,attr"`
	Scheda  []Anagrafe `xml:"anagrafica"`
}

type Temp struct {
	XMLNAME     xml.Name `xml:"temperatura_aria"`
	UnitaMis    string   `xml:"UM,attr"`
	Data        string   `xml:"data"`
	Temperatura float32  `xml:"temperatura"`
}

type Prec struct {
	XMLNAME  xml.Name `xml:"precipitazione"`
	UnitaMis string   `xml:"UM,attr"`
	Data     string   `xml:"data"`
	Pioggia  float32  `xml:"pioggia"`
}
type Vento struct {
	XMLNAME   xml.Name `xml:"vento_al_suolo"`
	UnitaMisV string   `xml:"UM_VV,attr"`
	UnitaMisD string   `xml:"UM_DV,attr"`
	Data      string   `xml:"data"`
	Vel       float32  `xml:"v"`
	Dir       int      `xml:"d"`
}
type RadiazioneG struct {
	XMLNAME  xml.Name `xml:"radiazioneglobale"`
	UnitaMis string   `xml:"UM,attr"`
	Data     string   `xml:"data"`
	Rad      float32  `xml:"rsg"`
}



type Dati struct {
	XMLNAME xml.Name `xml:"datiOggi"`
	attr1   string   `xml:"xmlns:xsi,attr"`
	attr2   string   `xml:"xmlns:xsd,attr"`
	attr3   string   `xml:"xmlns,attr"`
	Data    string   `xml:"data"`
	TempMin struct {
		Tm        float32 `xml:",chardata"`
		FlagNillm string  `xml:"xsi:nil,attr"`
	} `xml:"tmin"`
	TempMax struct {
		TM        float32 `xml:",chardata"`
		FlagNillM string  `xml:"xsi:nil,attr"`
	} `xml:"tmax"`

	P              float32       `xml:"rain"`
	Temperature    []Temp        `xml:"temperature>temperatura_aria"`
	Precipitazioni []Prec        `xml:"precipitazioni>precipitazione"`
	Venti          []Vento       `xml:"venti>vento_al_suolo"`
	Radiazione     []RadiazioneG `xml:"radiazione>radiazioneglobale"`
}

*/

// ==========================================================
// inizio definizione struct
// http://www.arpa.veneto.it/bollettini/meteo/h24/img12/Graf_231.htm?sens=TEMP
// http://www.arpa.veneto.it/bollettini/meteo/h24/img12/0231.xml
// ==========================================================

type DatoS struct {
          XMLNAME xml.Name `xml:"DATI"`
          Dataora   string  `xml:"ISTANTE,attr"`
         Valore       float64 `xml:"VM"`
}

type Sensore struct {
         XMLNAME xml.Name 	`xml:"SENSORE"`
         Id   	string				`xml:"ID"`
         Par 	string         			`xml:"PARAMNM"`
	 Typ  string 				`xml:"TYPE"`
	 Un	string				 `xml:"UNITNM"`
	 Note string 				 `xml:"NOTE"`
	 Freq int					`xml:"FREQ"`
	 Valori []DatoS                           `xml:"DATI"`


}

type Stazione struct {
         XMLNAME xml.Name 	`xml:"STAZIONE"`
         IDSt  int          `xml:"ISTAZ"`
	       Nome  string       `xml:"NOME"`
				 PosX  float64      `xml:"X"`
				 PosY  float64      `xml:"Y"`
        Quota  int 					`xml;"QUOTA"`
				TipoSt string       `xml:"TIPOSTAZ"`
        Prov   string				`xml:"PROVINCIA"`
				Comune string				`xml:"COMUNE"`
				DataAtt string			`xml:"ATTIVAZIONE"`
				Sensori []Sensore   `xml:"SENSORE"`
}

type Conten struct  {
  	XMLNAME xml.Name   `xml:"CONTENITORE"`
    Forn    string     `xml:"FORNITORE"`
	  IstRun	string     `xml:"ISTANTERUN"`
		Inizio  string     `xml:"INIZIO"`
		Fine   	string     `xml:"FINE"`
		Stazioni []Stazione `xml:"STAZIONE"`
}
