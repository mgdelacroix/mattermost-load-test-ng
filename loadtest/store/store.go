// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package store

import (
	"github.com/mattermost/mattermost-server/model"
)

type UserStore interface {
	Id() string
}

type MutableUserStore interface {
	UserStore
	
	// users

	SetUser(user *model.User) error
	User() (*model.User, error)

	// posts
	SetPost(post *model.Post) error
	Post(postId string) (*model.Post, error)

	// preferences
	SetPreferences(preferences *model.Preferences) error
	Preferences() (*model.Preferences, error)

	// channels
	SetChannel(channel *model.Channel) error
	Channel(channelId string) (*model.Channel, error)
	SetChannelMembers(channelId string, channelMembers *model.ChannelMembers) error
	ChannelMembers(channelId string) (*model.ChannelMembers, error)
	SetChannelMember(channelId string, channelMember *model.ChannelMember) error
	ChannelMember(channelId, userId string) (*model.ChannelMember, error)
	RemoveChannelMember(channelId string, userId string) error

	// teams
	SetTeamMember(teamId string, teamMember *model.TeamMember) error
	SetTeamMembers(teamId string, teamMember []*model.TeamMember) error
	TeamMember(teamdId, userId string) (*model.TeamMember, error)
}
