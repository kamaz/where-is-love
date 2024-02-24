package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	assert := assert.New(t)
	// create 2 users
	users := []map[string]any{}
	for range 10 {
		user := createUserEndpoint(assert)
		if len(users) != 2 {
			users = append(users, user)
		}
	}

	// login as user 1
	tokenUser1 := loginEndpoint(assert, users[0])

	matchesForUser1 := discoverEndpoint(assert, tokenUser1, "")
	// check that self does not exist in the result
	selfUser := findUser(matchesForUser1, users[0])
	assert.Nil(selfUser)

	// check that user 2 exists and we can match with them
	userTwo := findUser(matchesForUser1, users[1])
	assert.NotNil(userTwo)

	// match to user 2
	matchUser := users[1]

	// now lets swipe date for user 2
	firstSwipeResult := swipeEndpoint(assert, tokenUser1, map[string]any{
		"userID":     matchUser["id"],
		"preference": "YES",
	})
	assert.Contains(firstSwipeResult, "matched")
	assert.NotContains(firstSwipeResult, "matchID")

	// after swiping user two shouldn't be in matches
	matchesForUser1 = discoverEndpoint(assert, tokenUser1, "")
	userTwo = findUser(matchesForUser1, users[1])
	assert.Nil(userTwo)

	// now lets switch user ot matched and swipe as well so we see they are matched
	tokenUser2 := loginEndpoint(assert, users[1])

	matchesForUser2 := discoverEndpoint(assert, tokenUser2, "")
	// check that user 2 can match to user 1
	matchToUser1 := findUser(matchesForUser2, users[0])
	assert.NotNil(matchToUser1)

	// now lets swipe date with user 1
	secondSwipeResult := swipeEndpoint(assert, tokenUser2, map[string]any{
		"userID":     matchToUser1["id"],
		"preference": "YES",
	})
	// as both have swiped yes, they should be matched
	assert.Contains(secondSwipeResult, "matched")
	assert.Contains(secondSwipeResult, "matchID")

	matchesForUser2 = discoverEndpoint(assert, tokenUser2, "")
	userOne := findUser(matchesForUser2, users[0])
	assert.Nil(userOne)

	// filter by gender
	matchesForUser2FemaleOnly := discoverEndpoint(assert, tokenUser2, "gender=female")
	resultOnlyFemale := usersMatchPropertyValue(matchesForUser2FemaleOnly, "gender", "female")
	assert.True(resultOnlyFemale)

	// filter by age
	// todo: find an age that exists
	searchAge := fmt.Sprintf("%v", matchesForUser2[0].(map[string]any)["age"])
	matchesForUser2AgeOnly := discoverEndpoint(
		assert,
		tokenUser2,
		fmt.Sprintf("age=%s", searchAge),
	)
	resultOnlyAge := usersMatchPropertyValue(matchesForUser2AgeOnly, "age", searchAge)
	assert.True(resultOnlyAge)

	// filter by age and gender
	matchesForUser2FemaleAge := discoverEndpoint(
		assert,
		tokenUser2,
		fmt.Sprintf("gender=female&age=%s", searchAge),
	)
	resultAge22 := usersMatchPropertyValue(matchesForUser2FemaleAge, "age", searchAge)
	resultFemale := usersMatchPropertyValue(matchesForUser2FemaleAge, "gender", "female")
	assert.True(resultAge22 && resultFemale)

	sort := discoverEndpoint(
		assert,
		tokenUser2,
		"sort=%2BdistanceFromMe",
	)
	firstMatch := sort[0].(map[string]any)["distanceFromMe"]
	lastMatch := sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(lastMatch, firstMatch)

	sort = discoverEndpoint(
		assert,
		tokenUser2,
		"sort=-distanceFromMe",
	)
	firstMatch = sort[0].(map[string]any)["distanceFromMe"]
	lastMatch = sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(firstMatch, lastMatch)
}
