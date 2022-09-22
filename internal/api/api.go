package api

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/valyala/fasthttp"
	"golang.org/x/sync/errgroup"
	"pr0.bot/internal/models"
	"pr0.bot/pkg/configs"
	"pr0.bot/pkg/database"
	"pr0.bot/pkg/logger"
	"pr0.bot/pkg/utils"
)

type PostType struct {
	URL              string
	CurrentTime      int64
	Response         *models.Pr0Response
	SuccessfullPosts []*models.PostItem
	Caption          string
	Random           int
	FailedPosts      []*models.PostItem
}

func sendGetAsync(cookie string, url string, rc chan []byte) error {
	client := fasthttp.Client{
		DisablePathNormalizing: true,
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.Add("cookie", cookie)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err := client.Do(req, resp)

	if err == nil {
		rc <- resp.Body()
	}
	return err
}

// gets post from given category in yaml
func (p *PostType) GetPosts() error {
	respChan := make(chan []byte, 1)
	errGrp, _ := errgroup.WithContext(context.Background())

	cookie := configs.Config.Items.Cookie
	errGrp.Go(func() error { return sendGetAsync(cookie, p.URL, respChan) })

	err := errGrp.Wait()
	if err != nil {
		return err
	}

	var pr0Response models.Pr0Response
	response := <-respChan
	cleaned := strings.ReplaceAll(string(response), "\\", "")

	err = json.Unmarshal([]byte(cleaned), &pr0Response)
	if err != nil {
		return err
	}
	p.Response = &pr0Response
	return nil
}

// check if posts are already in db and if older than 48 hours
func (p *PostType) CleanPosts() error {
	for _, x := range p.Response.Items {
		if (p.CurrentTime - 172800) < int64(x.Created) {
			post := database.MongoDBInstance.PostExists(int32(x.ID))
			if !post {
				url := utils.BuildMediaURL(x.Image, x.Audio)
				p.SuccessfullPosts = append(p.SuccessfullPosts, &models.PostItem{
					MediaURL: url,
					Caption:  p.Caption,
					IsVideo:  x.Audio,
				})

				err := database.MongoDBInstance.Insert(int64(x.ID))
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func Watcher(flags int, tags string) error {
	postCategory := &PostType{}

	// to generate captions, we have to do this - sorry
	postCategory.Caption = strings.ReplaceAll(tags, "! s:-100 ", "")
	postCategory.Caption = strings.ReplaceAll(postCategory.Caption, " ", "")

	postCategory.URL = utils.BuildURL(flags, tags)
	postCategory.CurrentTime = utils.GenCurrentTime()
	postCategory.Random = utils.GenRandomTime()

	err := postCategory.GetPosts()
	if err != nil {
		return err
	}

	err = postCategory.CleanPosts()
	if err != nil {
		return err
	}

	postCategory.ProcessPosts()
	err = postCategory.ProcessFailedPosts()
	if err != nil {
		return err
	}

	postCategory = nil

	return nil
}

func (p *PostType) ProcessPosts() {
	for _, x := range p.SuccessfullPosts {
		_, err := utils.SendPost(x.MediaURL, x.Caption, x.IsVideo)
		if err != nil {
			p.FailedPosts = append(p.FailedPosts, x)
			logger.SugarLogger.Error(fmt.Sprintf("%s Media: %s\nAdded to failed posts and will try it again", err.Error(), x.MediaURL))
		}

		// use random generated waittime here, to prevent tg flood
		time.Sleep(time.Second * time.Duration(p.Random))
	}
}

func (p *PostType) ProcessFailedPosts() error {
	for _, x := range p.FailedPosts {
		resp, err := DownloadFile(x.MediaURL)
		if err != nil {
			// keep this logging statement
			logger.SugarLogger.Error(fmt.Sprintf("%s Media: %s", err.Error(), x.MediaURL))
			return err
		}

		_, err = utils.SendPostByte(x.MediaURL, x.Caption, x.IsVideo, resp)
		if err != nil {
			logger.SugarLogger.Error(fmt.Sprintf("%s Media: %s", err.Error(), x.MediaURL))
			return err
		}

		resp = nil
		// use random generated waittime here, to prevent tg flood
		time.Sleep(time.Second * time.Duration(p.Random))
	}
	return nil
}

func DownloadFile(url string) ([]byte, error) {
	var dst []byte
	statusCode, body, err := fasthttp.Get(dst, url)
	if statusCode != 200 {
		return nil, err
	}

	return body, nil
}
