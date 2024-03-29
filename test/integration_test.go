package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBothUserMatch(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

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
}

func TestBothUserDoNotMatch(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

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
		"preference": "NO",
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
		"preference": "NO",
	})
	// as both have swiped yes, they should be matched
	assert.Contains(secondSwipeResult, "matched")
	assert.NotContains(secondSwipeResult, "matchID")

	matchesForUser2 = discoverEndpoint(assert, tokenUser2, "")
	userOne := findUser(matchesForUser2, users[0])
	assert.Nil(userOne)
}

func TestFirstUserMatches(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

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
		"preference": "NO",
	})
	// as both have swiped yes, they should be matched
	assert.Contains(secondSwipeResult, "matched")
	assert.NotContains(secondSwipeResult, "matchID")

	matchesForUser2 = discoverEndpoint(assert, tokenUser2, "")
	userOne := findUser(matchesForUser2, users[0])
	assert.Nil(userOne)
}

func TestSecondUserMatches(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

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
		"preference": "NO",
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
	assert.NotContains(secondSwipeResult, "matchID")

	matchesForUser2 = discoverEndpoint(assert, tokenUser2, "")
	userOne := findUser(matchesForUser2, users[0])
	assert.Nil(userOne)
}

func TestFiltering(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

	// login as user 1
	tokenUser1 := loginEndpoint(assert, users[0])
	matchesForUser2 := discoverEndpoint(assert, tokenUser1, "")

	// filter by age
	searchAge := fmt.Sprintf("%v", matchesForUser2[0].(map[string]any)["age"])
	searchGender := fmt.Sprintf("%v", matchesForUser2[0].(map[string]any)["gender"])
	matchesForUser2AgeOnly := discoverEndpoint(
		assert,
		tokenUser1,
		fmt.Sprintf("age=%s", searchAge),
	)
	resultOnlyAge := usersMatchPropertyValue(matchesForUser2AgeOnly, "age", searchAge)
	assert.True(resultOnlyAge)

	// filter by age and gender
	matchesForUser2GenderAge := discoverEndpoint(
		assert,
		tokenUser1,
		fmt.Sprintf("gender=%s&age=%s", searchGender, searchAge),
	)
	resultAge := usersMatchPropertyValue(matchesForUser2GenderAge, "age", searchAge)
	resultGender := usersMatchPropertyValue(matchesForUser2GenderAge, "gender", searchGender)
	assert.True(resultAge && resultGender)

	sort := discoverEndpoint(
		assert,
		tokenUser1,
		"sort=%2BdistanceFromMe",
	)
	firstMatch := sort[0].(map[string]any)["distanceFromMe"]
	lastMatch := sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(lastMatch, firstMatch)

	sort = discoverEndpoint(
		assert,
		tokenUser1,
		"sort=-distanceFromMe",
	)
	firstMatch = sort[0].(map[string]any)["distanceFromMe"]
	lastMatch = sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(firstMatch, lastMatch)
}

func TestOrdering(t *testing.T) {
	assert := assert.New(t)
	// create a few users
	users := createFewUsers(assert)

	// login as user 1
	tokenUser1 := loginEndpoint(assert, users[0])

	sort := discoverEndpoint(
		assert,
		tokenUser1,
		"sort=%2BdistanceFromMe",
	)
	firstMatch := sort[0].(map[string]any)["distanceFromMe"]
	lastMatch := sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(lastMatch, firstMatch)

	sort = discoverEndpoint(
		assert,
		tokenUser1,
		"sort=-distanceFromMe",
	)
	firstMatch = sort[0].(map[string]any)["distanceFromMe"]
	lastMatch = sort[len(sort)-1].(map[string]any)["distanceFromMe"]
	assert.GreaterOrEqual(firstMatch, lastMatch)
}
