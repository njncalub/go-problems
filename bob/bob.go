/*
Package bob simulates a conversation with Bob, a lackadaisical teenager.

His responses are very limited and quite predictable for his age:

	Bob answers 'Sure.' if you ask him a question.
	He answers 'Whoa, chill out!' if you yell at him.
	He answers 'Calm down, I know what I'm doing!' if you yell a question
		at him.
	He says 'Fine. Be that way!' if you address him without actually saying
		anything.
	He answers 'Whatever.' to anything else.
*/
package bob

import (
	"regexp"
	"strings"
)

// Due to his limited vocabular and perhaps teenage angst, these are Bob's
// replies to most, if not all, kinds of conversations.
var (
	ResponseToQuestions         = "Sure."
	ResponseToShouting          = "Whoa, chill out!"
	ResponseToShoutingQuestions = "Calm down, I know what I'm doing!"
	ResponseToSilence           = "Fine. Be that way!"
	ResponseToAnythingElse      = "Whatever."
)
var alphaNumMatcher *regexp.Regexp = regexp.MustCompile(`[A-Za-z]`)

// Hey is used to start communicating with Bob. Parsing a string remark, Bob
// returns the appropriate response from his set of pre-programmed replies.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	containsLetters := alphaNumMatcher.MatchString(remark)
	wasSilence := remark == ""
	wasShouting := containsLetters && remark == strings.ToUpper(remark)
	wasAsking := strings.HasSuffix(remark, "?")

	switch {
	case wasSilence:
		return ResponseToSilence
	case wasShouting && wasAsking:
		return ResponseToShoutingQuestions
	case wasAsking:
		return ResponseToQuestions
	case wasShouting:
		return ResponseToShouting
	default:
		return ResponseToAnythingElse
	}
}
