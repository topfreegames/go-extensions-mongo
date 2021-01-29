package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
	"github.com/topfreegames/go-extensions-mongo/interfaces"
)

// Client is the struct that connects to PostgreSQL
type Client struct {
	Config  *viper.Viper
	MongoDB interfaces.MongoDB
}

//NewClient connects to Mongo server and return its client
func NewClient(prefix string, config *viper.Viper, mongoOrNil ...interfaces.MongoDB) (*Client, error) {
	client := &Client{
		Config: config,
	}
	var mongoDB interfaces.MongoDB
	if len(mongoOrNil) > 0 {
		mongoDB = mongoOrNil[0]
	}
	err := client.Connect(prefix, mongoDB)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func makeKey(prefix, sufix string) string {
	return fmt.Sprintf("%s.%s", prefix, sufix)
}

//Connect connects to mongo database and saves on Client
func (c *Client) Connect(prefix string, db interfaces.MongoDB) error {
	url := c.Config.GetString(makeKey(prefix, "url"))
	user := c.Config.GetString(makeKey(prefix, "user"))
	pass := c.Config.GetString(makeKey(prefix, "pass"))
	database := c.Config.GetString(makeKey(prefix, "database"))
	timeout := c.Config.GetDuration(makeKey(prefix, "connectionTimeout"))

	if db != nil {
		c.MongoDB = db
	} else {
		var session *mgo.Session
		var err error

		if timeout > 0 {
			session, err = mgo.DialWithTimeout(url, timeout)
		} else {
			session, err = mgo.Dial(url)
		}

		if err != nil {
			return err
		}
		mongoDB := session.DB(database)
		if user != "" && pass != "" {
			err = mongoDB.Login(user, pass)
			if err != nil {
				return err
			}
		}
		c.MongoDB = NewMongo(session, mongoDB)
	}

	return nil
}

//Close closes the session and the connection with database
func (c *Client) Close() {
	c.MongoDB.Close()
}