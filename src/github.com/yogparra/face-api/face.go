package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {

	//contraseña proporciona por nuestra api
	const key = "e8d8ed3978994c55bd69096475ccd634"
	
	//url proporciona por nuestra api
	const uriBase = "https://centralus.api.cognitive.microsoft.com/face/v1.0/detect"

	//url de nuestra a imagen a analizar
	//const imageUrl1 = "https://upload.wikimedia.org/wikipedia/commons/3/37/Dagestani_man_and_woman.jpg"
	const imageUrl2 = "https://avatars2.githubusercontent.com/u/3090002?s=460&v=4"


	//parametros del detalle de la imagen, solo solicitamos para el metodo faceAttributes los parametro de age y gender
	const params = "?returnFaceAttributes=age,gender"
	//,headPose,smile,facialHair,glasses,emotion,hair,makeup,occlusion,accessories,blur,exposure,noise"

	const uri = uriBase + params
	const imageUrlEnc = "{\"url\":\"" + imageUrl2 + "\"}"

	reader := strings.NewReader(imageUrlEnc)

	//creación del cliente http
	client := &http.Client	{
		Timeout: time.Second * 2,
	}

	//crea la solicitud pasando la url de la imagen al cuerpo de esta
	req, err := http.NewRequest("POST", uri, reader)
	if err != nil {
		panic(err)
	}

	//incorpora los encabezados
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Ocp-Apim-Subscription-Key", key)

	//envia la solicitud y espera la respuesta
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//Lee el cuerpo de la respuesta
	//Deja los datos en una matriz
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
	//Analiza los datos de json
	var f interface{}
	json.Unmarshal(data, &f)

	//Da formato y muestra los resultados del json
	jsonFormatted, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(jsonFormatted))
}
