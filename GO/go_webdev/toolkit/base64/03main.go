package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := `You've got the lot to burn
A shelve of pig smotherd cries
Is there a spirit that spits
Upon the exit of signs
Is anybody there
These steps keep on growing long
Bayonet trials rust propellers await
No
Nobody is heard
Rowing sheep smiles for the dead
Nobody is heard
An antiquated home
Afloat with engines on mute
Sui generis ship spined around the yard
Is anybody there
These craft only multiply
At the nape of ruins rust propellers await
No
Nobody is heard compass wilting in the wind
Nobody is heard
Rowing sheep smile for the dead
Tansoceanic depth in this earth
In this cenotaph
Lash of one thousand eye brows clicking
Counting the toll
Counting the toll
You've got the lot to burn
A shelve of pig smothered cries
Is there a spirit that spits upon the exit of signs
Is anybody there
These steps keep on growing long
Bayonet trials rust propellers await
No
Nobody is heard compass wilting in the wind
Nobody is heard rowing sheep smile for the dead
Transoceanic depth in this earth in this cenotaph
Carpel jets
Hit the ground
Carpel jets
Hit the ground
Carpel jets
Hit the ground
Carpel jets
Hit the ground
Lash of one thousand eyebrows clicking
Counting the toll
Counting the toll
Lash of one thousand eyebrows clicking
Counting the toll
Counting the toll`

s64 := base64.StdEncoding.EncodeToString([]byte(s))

fmt.Println(s64)

bs, err := base64.StdEncoding.DecodeString(s64)
if err != nil {
	log.Fatalln("We've crashed the drunkship")
}
fmt.Println(string(bs))
}
