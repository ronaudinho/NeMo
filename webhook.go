package main

import (
	"log"
	req "github.com/imroc/req"
)

func SentToWebhook(url string, Sessions Session) (int, error) {
	r, err := req.Post(url, req.BodyJSON(Sessions))
	if err != nil {
		return 500, err
	}

	resp := r.Response()

	return resp.StatusCode, nil
}

func SentToDiscord(url string, Sessions Session) (bool, error) {
	// req.Debug = true
	compiled_message := "\n[" + Sessions.Created + "]\n" + Sessions.CurrentProcess
	for index := range(Sessions.Datas) {
		compiled_message = compiled_message + "\n[" + Sessions.Datas[index].Created + "]"
		compiled_message = compiled_message + "\nnemo: " + Sessions.Datas[index].Slug + " - " + Sessions.Datas[index].Question
		compiled_message = compiled_message + "\n" + Sessions.PhoneNumber + ": " + Sessions.Datas[index].Answer
	}

	var Discord = discord {
		Content: compiled_message,
	}
	_, err := req.Post(url, req.BodyJSON(Discord))

	if err != nil {
		return false, err
	}

	return true, nil
}

func LogToWebhook(url string, logGreeting LogGreeting) (int, error) {
	r, err := req.Post(url, req.BodyJSON(logGreeting))
	if err != nil {
		return 500, err
	}

	resp := r.Response()

	return resp.StatusCode, nil
}

func LogToDiscord(url string, logGreeting LogGreeting) (bool, error) {
	// req.Debug = true
	compiled_message := logGreeting.PhoneNumber + " replied with " + logGreeting.Message

	var Discord = discord {
		Content: compiled_message,
	}
	_, err := req.Post(url, req.BodyJSON(Discord))

	if err != nil {
		log.Println("ERROR: "+ err.Error())
		return false, err
	}

	return true, nil
}
