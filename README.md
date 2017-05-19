Deltakere:

Erdvik, Magnus

Lie, Eva Kristine

Nguyen, Philip

Tellefsen, Erlend Frøysnes

Van Dijk, Richard

Younas, Osman

		
		
						ICA 04 - Gruppe 11


						Oppgave 1a)
Les text1.txt og text2.txt vha. fileutils inn i “byte slices” og sammenlign disse. Filene har tilsynelatende lik innhold, men hvorfor er den ene filen større enn den andre? Bortsett fra forskjeller i størrelsen, hvorfor må man være oppmerksom på dette “fenomenet”? 


Tekstfil 1 er større enn tekstfil 2 fordi den har en karakter for carriage return “/r” (unicode: U+0013) i tillegg til line feed “/n”, mens tekstfil 2 har kun line feed. Dette gjør at tekstfilen for flere bytes (til sammen 4 bytes).
Dette kommer av at tekstfil 1 er skrevet ved bruk av “CRLF” linjeskift metoden, som brukes på windows operativssystemer, mens tekstfil 2 er skrevet ved bruk av “LF” linjeskift, som brukes på Mac os X systemer.. 

Bruk main.go som tester hvilken av tekstene som er windows eller mac sin metode for linjeskift.

Bilde:
http://imgur.com/a/3XLGg 



							Oppgave 1b)
Skriv en funksjon, som finner ut hvilken kode for linjeskift en tekst-fil bruker. Dere skal selv finne eget navn for deres funksjoner, men programmet skal hete “lineshift.go”. Programmet skal kunne bli utført fra kommandolinje og returnere en eller annen type konklusjon. Programmet skal ta et som in-data. 


Kommentar: Funksjonen “IsLineshiftMacOrWin” tar inn en fil som parameter fra kommandolinje. Videre lager funksjonen en byte slice av teksten i filen. Bytene i filen blir “matchet” ved ‘\r\n\’. Om en byte i filen er den, vil det si at filen bruker Windows linjeskift (CRLF), om ikke er det Mac (LF) 
Bilde: 


Package regexp implements regular expression search. Bruker regexp for å søke etter lineshift \r\n på windows.  ioutil.ReadAll og os for å lese / åpne filene.  https://golang.org/pkg/regexp/

Bilde:
http://imgur.com/a/JI0NO






							Oppgave 2a)
Skriv et Golang program fileinfo.go​, som kan utføres fra kommandolinje. Programmet skal ta et argument, som skal være et filnavn. Stien til filnavnet kan være enten relativ eller absolutt.

Bruk main.go med filnavn som parameter, for å skrive ut filinformasjon.



bilde:
http://imgur.com/a/rPtiM



							Oppgave 2b)
Forklar resultatet, som deres program returnerer for filene “/dev/stdin” og “/dev/ram0”, når det utføres i et Linux-miljø (i sky instansen). 


Bilde:
http://imgur.com/a/mEjD5

ram0 viser mindre rettigheter, selv med administratorrettigheter. Filene er mindre en 1 byte. Begge filene er verken device files, dermed blir det heller ikke Unix character device eller Unix block device.



							Oppgave 2c)

Bilde:
http://imgur.com/a/xQMSx


På windows systemet har filen text1.txt write rettigheter for “alle”, mens den ikke har det på linux systemet.





							Oppgave 3a)
Klassifiser og forklar hvilken type metoder finnes det i Go for å jobbe med filer. 

De grunnleggende metodene:
-Lage en ny tom fil uten noe innhold. Spesifiser navnet på filen og bruker err variabel for å sjekke om det er en lovlig funksjon å utføre. Printer så ut den nye filen.

-Sette en fil til å inneholde en spesifisert grense av bytes i parameter etter filnavnet. Er fila under grensen, så vil alt innholdet være intakt og resten av plassen bli fylt med tomme bytes. Er filen over grensen satt, så vil alt over bli tapt.

-Få fil info. En metode for å returnere statistikk over filen. Her kan vi få printer ut all info om størrelse. filtype, adgang restriksjoner, dato for modifikasjon osv. 

-Sette nytt navn på filen og flytte den. Metoden bruker den originale plassering av filen og spesifiserer den nye plassen filen skal bli flyttet til.

-Slette en fil ved å spesifisere filen i parameteret.

-Åpne og lukke filer. Her kan man spesifisere i parametere hvilken tilgang metoder filen skal ha. Som lese, skrive, eller forskjellige metoder som skjer med filen når den åpnes. Som å lage en fil hvis filnavnet ikke eksiteres fra før eller sette filen til en grense størrelse. 

-Sjekke om filer eksisterer. Metoden tar et parameter om filplassering og returnerer informasjon om filen, hvis ikke returnerer en feil for at filen ikke finnes.

-Sjekke lese og skrive tilganger for en fil. I parameter for funksjonen så definerer vi hvilke tilganger vi skal sjekke filen har for brukeren. Den vil deretter printe ut om man har tilgang eller ikke.

- Forandre tilgang, eierskap og når sist filene er modifisert.
Her kan man også modifisere filene med tidspunkter som er satt til fremtiden” som gjør at filene oppdaterer seg selv på tidspunktet som er satt.

-Harde lenker og symbolske lenker. Her skaper man flere pekere/veikart til en valgt fil. Hvis nå en fil blir slettet eller får tildelt et nytt navn så vil den ikke påvirke den andre lenken til filen, og vil fortsatt ligge i systemet. Symbolske lenker refererer ikke til selve filplasseringen, bare navnet på filen.

Lese og skrive.

-Kopiere filer- Metoden tar filnavnet som parameter. Filen blir “satt til siden” mens filen blir kopiert. Her overføres bytes fra originale filen til kopien i minne, før det blir flushet til harddisken. Deretter blir filene satt i sync.

-Skrive bytes til en fil. Metoden kan brukes til å åpne eksisterende filer eller skrive nye filer og dumpe overflødig bytes eller byte slices til en fil. 

-Å skrive til disk med en buffer. Metoden tar en fil med en opprettet buffer på minne. Metoden kan brukes til å jobbe med filer i bufferen, før man overføres bytes til disken. En grei måte å modifisere store filer, før man lagrer alt på harddisken.
Metoder gir muligheter til bestemme hvor stor bufferen skal være og hvor mye som skal skrive til disk. Man kan også gjenopprette forandringer i bufferen før det blir skrevet permanent.

-Lese eksakt, minst og alle bytes funksjonen gir muligheter å spesifisere akkurat hvor mange bytes av en fil som skal bli lest ved hjelp av byteslice. Den printer deretter ut hvor mange bytes som er blitt lest i hvilket format som er spesifisert.

-Lese filer med en scanner.
Metoden åpner en fil med en scanner funksjon som steg for steg (default er linjeskift) går igjennom filen og returnerer det som er spesifisert i metoden. Dette kan være ord, tekst eller bytes f.eks. Funksjonen kan kjøres kontuniterling til alle elementene er funnet ved å forstørre bufferen sin.

Zipping

-Lage og ekstraktere zip filer. 
Metoden i go bruker zip, men standard biblioteket støttes også arkivering av filer.
Man spesifiserer hvilken fil som skal zippes, for å lette en zip skriver på toppen av fil skriveren. Man kan legge til hardkodet data, eller navn på filene, eller alle filene i en mappestruktur som skal arkiveres. Man kan også velge filer og arkivere spesifiserte setninger i en tekstfil som skal med. Innholdet blir skrevet og arkivert som blir sendt til den underliggende zip skriveren som lager en zip fil.

-Ekstraktere filene fungerer ved å åpne zip filen med en zip leser. Zip leseren itererer over alle filene før man lukker den, og pakker ut innholdet til filplassering som er valgt. En mappe vil bli laget hvis den ikke finnes fra før. Innholdet i de arkiverte filene vil overføre restriksjonene for tilgangen til filene fra det originale innholdet.
Den nye filplasseringen blir opprettet og funksjonen kopierer innholdet fra den kopierte zip filen til nye plasseringen.

- Komprimere filer tar i bruk en metode som skriver filene til gz filer(zip), ved hjelp av en gz skriver. Man definerer hvilken fil som skal komprimeres og sender denne til skriveren som krymper alt innholdet.

-U komprimere filene åpner gzip filen ved hjelp av en gzip reader. Dette blir ofte brukt for å hente filene fra en webserver for å minske lastetiden for filen og bruke mindre plass. Etter gzip leseren er ferdig med å hente innholdet i filen, bruker vi en output skriver ved å kopiere gzip filen for å  lage en ny fil med opprinnelig innhold.
Diverse metoder

-Midlertidig filer og mapper. Metoden oppretter et midlertidig direktiv i en midlertidig mappe. Oppretter en fil, der man kan spesifisere i koden, hva som skal bli gjort med filen før den blir lukket og fjernet. Alt blir lagret i tmp mappen på et linux system som er systemets mappe for midlertidig innhold.

-Laste ned filer over http. Metoden oppretter en ny fil fra en url som man legger til. Metodene skriver bytes fra http området til en fil ved hjelp av en kopieringsfunksjon (io.Copy)

-Hashing og checksum. Metoden henter bytes fra en fil og printer ut resultatene i et mer komprimert og kortere verdier som representer original innholdet. Checksum brukes til å sjekke om innholdet er identisk.






							Oppgave 3b)

Skriv programmer for alle type metoder for å lese og skrive til filer. Programmene skal lese inn en tekst-fil (flere middels store filer er i repository), skrive ut totalt antall linjer og antall for fem “runes” (bokstaver/tegn), som forekommer mest i disse filene. 


Bruk main.go filene i oppgave 3 mappen.
Copy : Kopierer en fil, og forteller hvor mange linjer som er i filen.
Write : Oppretter en ny fil, for å så skrive bytes til den.
Read : Leser hele filen, og gir informasjon om bytes.



							Oppgave 3c) 
Lag benchmarks-tester for alle type metoder for lesing av fil. Forklar fordeler og ulemper for hver av disse metodene i forhold til oppgaven deres. Dere kan ha behov til å se mer på slices i Go: https://blog.golang.org/slices






							Oppgave 4a) 

Høstsemester 2014:
UIAs fakultet
Antall studenter (10539)

Sannsynlighet
Helse og idrettsfag:
1829 studenter =
17.346358118361152 %

Humaniora og pedagogikk:
1525 studenter =
14.463201820940819 %

Kunstfag
420 studenter =
3.983308042488619 %

Teknologi og realfag
2166 studenter =
20.54248861911988 %

Lærerutdanningen
1506 studenter =
14.283004552352049 %

Økonomi og samfunnsvitenskap
3093 studenter =
29.38163884673748 %

							Oppgave4b) 
	
Fakultet
Informasjonsmengde:

Helse og Idrettsfag 1829 = 11100100101
Humaniora og pedagogikk 1525 = 10111110101
Kunstfag 420 =   110100100
Teknologi og realfag 2166 = 100001110110
Lærerutdanning 1506 = 10111100010
Økonomi og samfunnsvitenskap 3093 = 110000010101




Vi kan se at om man får informasjon om en student fra Kunstfag får man minst informasjon. 

							Oppgave4c)

Bilde:
http://imgur.com/a/OOCDG


Følgende fakuliteter har denne huffmankoden

A: Helse og idrettsfag 100
B: Humaniora og pedagogikk 101
C: Kunstfag 000
D: Teknologi og realfag 01
E: Lærerutdanning 001
F: Økonomi og samfunnsvitenskap 11


							Oppgave4d) 


000 = 17  -  001 = 14  -  111 =  4  - 10 = 21   -110 = 14   -  01 = 29
Gjennomsnittslengden for meldingen som inneholder fakultets koder for 100 tilfeldig valgte studenter er lik 250













