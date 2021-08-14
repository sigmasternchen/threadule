package logic

import (
	"threadule/backend/internal/data/models"
)
import "github.com/dghubble/oauth1"
import "github.com/dghubble/go-twitter/twitter"

func (l *Logic) sendTweet(client *twitter.Client, tweet *models.Tweet, prevId int64) (int64, error) {
	status, _, err := client.Statuses.Update(
		tweet.Text,
		&twitter.StatusUpdateParams{
			InReplyToStatusID: prevId,
		},
	)

	if err != nil {
		tweet.Status = models.TweetFailed
		errorString := new(string)
		*errorString = err.Error()
		tweet.Error = errorString
		_ = l.ctx.Data.UpdateTweet(tweet)
		// TODO log data error

		return 0, err
	}

	tweet.Status = models.TweetDone
	_ = l.ctx.Data.UpdateTweet(tweet)
	// TODO log data error

	return status.ID, nil
}

func (l *Logic) getTwitterClient(account *models.Account) *twitter.Client {
	config := oauth1.NewConfig(
		l.ctx.Config.Twitter.ConsumerKey,
		l.ctx.Config.Twitter.ConsumerSecret,
	)
	token := oauth1.NewToken(
		*account.AccessToken,
		*account.AccessTokenSecret,
	)
	httpClient := config.Client(oauth1.NoContext, token)

	return twitter.NewClient(httpClient)
}

func (l *Logic) sendThread(thread *models.Thread) {
	client := l.getTwitterClient(thread.Account)

	thread.Status = models.ThreadProcessing
	_ = l.ctx.Data.UpdateThread(thread)
	// TODO log data error

	tweets, err := l.ctx.Data.GetTweetsForThread(thread)
	if err != nil {
		// TODO log error
		errorString := new(string)
		*errorString = err.Error()
		thread.Status = models.ThreadFailed
		thread.Error = errorString
		_ = l.ctx.Data.UpdateThread(thread)
		// TODO log error

		return
	}
	failed := false
	var tweetIds []int64
	lastId := int64(0)
	for _, tweet := range tweets {
		lastId, err = l.sendTweet(client, &tweet, lastId)
		if err != nil {
			failed = true
			break
		} else {
			tweetIds = append(tweetIds, lastId)
		}
	}

	if failed {
		for _, id := range tweetIds {
			_, _, err = client.Statuses.Destroy(id, nil)
			if err != nil {
				// TODO log error
			}
		}

		thread.Status = models.ThreadFailed
		if err != nil {
			// should always be != nil but the compiler can't see that
			errorString := new(string)
			*errorString = err.Error()
			thread.Error = errorString
		}

		_ = l.ctx.Data.UpdateThread(thread)
		// TODO log data error
	}

	thread.Status = models.ThreadDone
	_ = l.ctx.Data.UpdateThread(thread)
	// TODO log data error
}

func (l *Logic) scheduleTrigger() {
	threads, err := l.ctx.Data.GetScheduledThreads()
	if err != nil {
		// TODO log error
		return
	}

	for _, thread := range threads {
		l.sendThread(&thread)
	}
}
