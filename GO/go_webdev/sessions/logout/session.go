package main

import (
	"net/http"
)

func getUser(req *http.Request) user {
	var u user

	// I did it for the cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// return user if there is a user that exists
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}