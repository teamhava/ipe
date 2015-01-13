// Copyright 2014 Claudemiro Alves Feitosa Neto. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"testing"
)

func newApp() App {
	return App{Name: "Test", AppID: "123", Key: "123", Secret: "123", OnlySSL: false, ApplicationDisabled: false, UserEvents: true}
}

func Test_add_channels(t *testing.T) {

	app := newApp()

	// Public

	if len(app.PublicChannels) != 0 {
		t.Error("Length of public channels must be 0 before test")
	}

	app.AddChannel(NewChannel("ID", ""))

	if len(app.PublicChannels) != 1 {
		t.Error("Length os public channels after insert must be 1")
	}

	// Presence

	if len(app.PresenceChannels) != 0 {
		t.Error("Length of presence channels must be 0 before test")
	}

	app.AddChannel(NewChannel("presence-test", ""))

	if len(app.PresenceChannels) != 1 {
		t.Error("Length os presence channels after insert must be 1")
	}

	// Private

	if len(app.PrivateChannels) != 0 {
		t.Error("Length of private channels must be 0 before test")
	}

	app.AddChannel(NewChannel("private-test", ""))

	if len(app.PrivateChannels) != 1 {
		t.Error("Length os private channels after insert must be 1")
	}

}

func Test_AllChannels(t *testing.T) {
	app := newApp()
	app.AddChannel(NewChannel("private-test", ""))
	app.AddChannel(NewChannel("presence-test", ""))
	app.AddChannel(NewChannel("test", ""))

	if len(app.AllChannels()) != 3 {
		t.Error("Must have 3 channels")
	}
}

func Test_New_Subscriber(t *testing.T) {
	app := newApp()

	if len(app.Subscribers) != 0 {
		t.Error("Length of subscribers before test must be 0")
	}

	conn := NewSubscriber("1", "", nil)
	app.AddSubscriber(conn)

	if len(app.Subscribers) != 1 {
		t.Error("Length os subscribers after test must be 1")
	}
}

func Test_find_subscriber(t *testing.T) {
	app := newApp()
	conn := NewSubscriber("1", "", nil)
	app.AddSubscriber(conn)

	conn, err := app.FindSubscriber("1")

	if err != nil {
		t.Error(err)
	}

	if conn.SocketID != "1" {
		t.Error("Wrong subscriber.")
	}

	// Find a wrong subscriber

	conn, err = app.FindSubscriber("DoesNotExists")

	if err == nil {
		t.Error("Opps, Must be nil")
	}

	if conn != nil {
		t.Error("Opps, Must be nil")
	}
}

func Test_find_or_create_channels(t *testing.T) {
	app := newApp()

	// Public
	if len(app.PublicChannels) != 0 {
		t.Error("Length of public channels must be 0 before test")
	}

	c := app.FindOrCreateChannelByChannelID("id", "")

	if len(app.PublicChannels) != 1 {
		t.Error("Length os public channels after insert must be 1")
	}

	if c.ChannelID != "id" {
		t.Error("Opps wrong channel")
	}

	// Presence
	if len(app.PresenceChannels) != 0 {
		t.Error("Length of presence channels must be 0 before test")
	}

	c = app.FindOrCreateChannelByChannelID("presence-test", "")

	if len(app.PresenceChannels) != 1 {
		t.Error("Length os presence channels after insert must be 1")
	}

	if c.ChannelID != "presence-test" {
		t.Error("Opps wrong channel")
	}

	// Private
	if len(app.PrivateChannels) != 0 {
		t.Error("Length of private channels must be 0 before test")
	}

	c = app.FindOrCreateChannelByChannelID("private-test", "")

	if len(app.PrivateChannels) != 1 {
		t.Error("Length os private channels after insert must be 1")
	}

	if c.ChannelID != "private-test" {
		t.Error("Opps wrong channel")
	}

}