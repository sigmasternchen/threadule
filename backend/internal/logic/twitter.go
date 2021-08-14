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
		l.ctx.Log.Warningf("couldn't send tweet: %v", err)
		tweet.Status = models.TweetFailed
		errorString := new(string)
		*errorString = err.Error()
		tweet.Error = errorString
		err2 := l.ctx.Data.UpdateTweet(tweet)
		if err2 != nil {
			l.ctx.Log.Errorf("couldn't update tweet in DB: %v", err2)
		}

		return 0, err
	}

	tweet.Status = models.TweetDone
	err = l.ctx.Data.UpdateTweet(tweet)
	l.ctx.Log.Errorf("couldn't update tweet in DB: %v", err)

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
	err := l.ctx.Data.UpdateThread(thread)
	l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)

	tweets, err := l.ctx.Data.GetTweetsForThread(thread)
	if err != nil {
		l.ctx.Log.Errorf("couldn't get tweets from DB: %v", err)

		errorString := new(string)
		*errorString = err.Error()
		thread.Status = models.ThreadFailed
		thread.Error = errorString
		err = l.ctx.Data.UpdateThread(thread)
		l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)

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
				l.ctx.Log.Errorf("couldn't destroy tweets: %v", err)
			}
		}

		thread.Status = models.ThreadFailed
		if err != nil {
			// should always be != nil but the compiler can't see that
			errorString := new(string)
			*errorString = err.Error()
			thread.Error = errorString
		}

		err = l.ctx.Data.UpdateThread(thread)
		l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)

		return
	}

	thread.Status = models.ThreadDone
	err = l.ctx.Data.UpdateThread(thread)
	l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)
}

func (l *Logic) scheduleTrigger() {
	threads, err := l.ctx.Data.GetScheduledThreads()
	if err != nil {
		l.ctx.Log.Errorf("couldn't get scheduled threads from DB: %v", err)
		return
	}

	for _, thread := range threads {
		l.sendThread(&thread)
	}
}
