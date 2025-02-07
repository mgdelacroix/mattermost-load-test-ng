// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package simplecontroller

import (
	"errors"
	"math"
	"time"

	"github.com/mattermost/mattermost-load-test-ng/loadtest/user"
)

type SimpleController struct {
	user user.User
	stop chan bool
	rate float64
}

func (c *SimpleController) Init(user user.User) {
	c.user = user
	c.stop = make(chan bool)
	c.rate = 1.0
}

func (c *SimpleController) Run(status chan<- user.UserStatus) {
	if c.user == nil {
		c.sendFailStatus(status, "controller was not initialized")
		return
	}

	actions := []UserAction{
		{
			run:       c.signUp,
			waitAfter: 1000,
		},
		{
			run:       c.login,
			waitAfter: 1000,
		},
		{
			run:       c.createPost,
			waitAfter: 1000,
		},
		{
			run:       c.createGroupChannel,
			waitAfter: 1000,
		},
		{
			run:       c.viewChannel,
			waitAfter: 1000,
		},
		{
			run:       c.logout,
			waitAfter: 1000,
		},
	}

	status <- user.UserStatus{User: c.user, Info: "user started", Code: user.STATUS_STARTED}

	defer c.sendStopStatus(status)

	for {
		for i := 0; i < len(actions); i++ {
			status <- actions[i].run()

			idleTime := time.Duration(math.Round(float64(actions[i].waitAfter) * c.rate))

			select {
			case <-c.stop:
				return
			case <-time.After(time.Millisecond * idleTime):
			}
		}

		// status <- user.UserStatus{User: c.user, Info: "user loop done", Code: user.STATUS_DONE}
	}
}

func (c *SimpleController) SetRate(rate float64) error {
	if rate < 0 {
		return errors.New("rate should be a positive value")
	}
	c.rate = rate
	return nil
}

func (c *SimpleController) Stop() {
	close(c.stop)
}

func (c *SimpleController) sendFailStatus(status chan<- user.UserStatus, reason string) {
	status <- user.UserStatus{User: c.user, Code: user.STATUS_FAILED, Err: errors.New(reason)}
}

func (c *SimpleController) sendStopStatus(status chan<- user.UserStatus) {
	status <- user.UserStatus{User: c.user, Info: "user stopped", Code: user.STATUS_STOPPED}
}
