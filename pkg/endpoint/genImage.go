package endpoint

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/g91TeJl/AiBot/pkg/model"
)

func GenImage(message string, i int, apiKey string) { //имя сменить
	// Build REST endpoint URL w/ specified engine
	engineId := "stable-diffusion-v1-5"
	// apiHost, hasApiHost := os.LookupEnv("API_HOST")
	// if !hasApiHost {
	// 	apiHost = "https://api.stability.ai"
	// }
	apiHost := "https://api.stability.ai"
	reqUrl := apiHost + "/v1/generation/" + engineId + "/text-to-image"

	// Acquire an API key from the environment
	var data = []byte(fmt.Sprintf(`{
		"text_prompts": [
			  {
				"text": "%s"
			  }
			],
			"cfg_scale": 7,
			"clip_guidance_preset": "FAST_BLUE",
			"height": 512,
			"width": 512,
			"samples": 1,
			"steps": 150}`, message))
	req, _ := http.NewRequest("POST", reqUrl, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)

	// Execute the request & read all the bytes of the body
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		var body map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
			panic(err)
		}
		panic(fmt.Sprintf("Non-200 response: %s", body))
	}

	// Decode the JSON body
	var body model.TextToImageResponse
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		panic(err)
	}

	// Write the images to disk
	for _, image := range body.Images {
		//outFile := fmt.Sprintf("./out/v1_txt2img_%d.png", i)
		outFile := fmt.Sprintf("./out/%s_%s.png", message, strconv.Itoa(i))
		file, err := os.Create(outFile)
		if err != nil {
			panic(err)
		}

		imageBytes, err := base64.StdEncoding.DecodeString(image.Base64)
		if err != nil {
			panic(err)
		}

		if _, err := file.Write(imageBytes); err != nil {
			panic(err)
		}

		if err := file.Close(); err != nil {
			panic(err)
		}
	}
}
