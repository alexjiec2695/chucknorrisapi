package externalservices

import (
	"chucknorris/entities"
	"fmt"
	"github.com/go-resty/resty/v2"
	"sync"
)

type chuckNorris struct {
	client *resty.Client
	url    string
}

type ChuckNorris interface {
	GetItem(chan<- entities.ChuckNorrisItem, *sync.WaitGroup, chan int)
}

func NewChuckNorrisServices(urlService string) ChuckNorris {
	client := resty.New()
	return &chuckNorris{
		client: client,
		url:    urlService,
	}
}

func (c *chuckNorris) GetItem(chanel chan<- entities.ChuckNorrisItem, wg *sync.WaitGroup, worker chan int) {
	defer func() {
		wg.Done()
		<-worker
	}()

	item := entities.ChuckNorrisItem{}

	_, err := c.client.R().
		SetResult(&item).
		ForceContentType("application/json").
		Get(c.url)

	if err != nil {
		fmt.Println(err)
	}

	chanel <- item
}
