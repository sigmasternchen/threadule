package logic

import (
	"net/url"
	"threadule/backend/internal/data/models"
)
import "github.com/dghubble/oauth1"
import "github.com/dghubble/go-twitter/twitter"
import twitterOAuth "github.com/dghubble/oauth1/twitter"

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

func (l *Logic) getAcountClient(account *models.Account) *twitter.Client {
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
	client := l.getAcountClient(thread.Account)

	thread.Status = models.ThreadProcessing
	err := l.ctx.Data.UpdateThreadWithoutTweets(thread)
	l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)

	tweets, err := l.ctx.Data.GetTweetsForThread(thread)
	if err != nil {
		l.ctx.Log.Errorf("couldn't get tweets from DB: %v", err)

		errorString := new(string)
		*errorString = err.Error()
		thread.Status = models.ThreadFailed
		thread.Error = errorString
		err = l.ctx.Data.UpdateThreadWithoutTweets(thread)
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

		err = l.ctx.Data.UpdateThreadWithoutTweets(thread)
		l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)

		return
	}

	thread.Status = models.ThreadDone
	err = l.ctx.Data.UpdateThreadWithoutTweets(thread)
	l.ctx.Log.Errorf("couldn't update thread in DB: %v", err)
}

func (l *Logic) scheduleTriggerTwitter() {
	threads, err := l.ctx.Data.GetScheduledThreads()
	if err != nil {
		l.ctx.Log.Errorf("couldn't get scheduled threads from DB: %v", err)
		return
	}

	for _, thread := range threads {
		l.sendThread(&thread)
	}
}

func (l *Logic) getTwitterOAuthConfig() *oauth1.Config {
	return &oauth1.Config{
		ConsumerKey:    l.ctx.Config.Twitter.ConsumerKey,
		ConsumerSecret: l.ctx.Config.Twitter.ConsumerSecret,
		CallbackURL:    "oob",
		Endpoint:       twitterOAuth.AuthorizeEndpoint,
	}
}

func (l *Logic) twitterLoginInit(account *models.Account) (string, *url.URL, error) {
	oauth1Config := l.getTwitterOAuthConfig()

	requestToken, requestSecret, err := oauth1Config.RequestToken()
	if err != nil {
		l.ctx.Log.Errorf("couldn't get requestToken: %v", err)
		return "", nil, ErrInternalError
	}

	account.RequestToken = &requestToken
	account.RequestSecret = &requestSecret
	err = l.ctx.Data.UpdateAccount(account)
	if err != nil {
		l.ctx.Log.Errorf("couldn't update account in database: %v", err)
		return "", nil, ErrInternalError
	}

	authUrl, err := oauth1Config.AuthorizationURL(requestToken)
	if err != nil {
		l.ctx.Log.Errorf("couldn't get authorization url: %v", err)
		return "", nil, ErrInternalError
	}

	return account.ID.String(), authUrl, nil
}

func (l *Logic) twitterLoginResolve(account *models.Account, pin string) error {
	oauth1Config := l.getTwitterOAuthConfig()

	accessToken, accessSecret, err := oauth1Config.AccessToken(*account.RequestToken, *account.RequestSecret, pin)
	if err != nil {
		l.ctx.Log.Errorf("couldn't get access token: %v", err)
		return ErrInternalError
	}

	account.AccessToken = &accessToken
	account.AccessTokenSecret = &accessSecret
	account.RequestToken = nil
	account.RequestSecret = nil

	twitterClient := l.getAcountClient(account)

	accountVerifyParams := &twitter.AccountVerifyParams{
		IncludeEntities: twitter.Bool(false),
		SkipStatus:      twitter.Bool(true),
		IncludeEmail:    twitter.Bool(false),
	}
	user, _, err := twitterClient.Accounts.VerifyCredentials(accountVerifyParams)
	if err != nil {
		l.ctx.Log.Errorf("couldn't verify credentials: %v", err)
		return ErrInternalError
	}

	account.TwitterID = &user.ID
	account.Name = user.Name
	account.ScreenName = user.ScreenName
	account.AvatarURL = user.ProfileImageURL

	err = l.ctx.Data.UpdateAccount(account)
	if err != nil {
		l.ctx.Log.Errorf("couldn't update account in database: %v", err)
		return ErrInternalError
	}

	return nil
}
