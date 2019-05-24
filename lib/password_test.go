package lib

import (
	"encoding/base64"
	"testing"
)

func TestRandomSalt64(t *testing.T) {
	randomSalt := RandomSalt64()

	deserializedRandomSalt, err := base64.StdEncoding.DecodeString(randomSalt)

	if err != nil {
		t.Fatalf("random salt %s could not be decoded", randomSalt)
	}
	if len(deserializedRandomSalt) > saltRandomBytes {
		t.Fatalf("random salt %s was not %d bytes but was %d", randomSalt, saltRandomBytes, len(deserializedRandomSalt))
	}
}

func TestHashPassword64(t *testing.T) {
	firstUsername := "jack"
	firstSalt := "89123n4kjlv"
	firstPassword := "password"

	secondUsername := "jill"
	secondSalt := firstSalt
	secondPassword := firstPassword

	thirdUsername := firstUsername
	thirdSalt := "89123n4kjlva"
	thirdPassword := firstPassword

	fourthUsername := firstUsername
	fourthSalt := firstSalt
	fourthPassword := "pashword"

	firstHash := HashPassword64(firstUsername, firstSalt, firstPassword)
	secondHash := HashPassword64(secondUsername, secondSalt, secondPassword)
	thirdHash := HashPassword64(thirdUsername, thirdSalt, thirdPassword)
	fourthHash := HashPassword64(fourthUsername, fourthSalt, fourthPassword)
	fifthHash := HashPassword64(firstUsername, firstSalt, firstPassword)

	if firstHash == secondHash {
		t.Fatalf("change in username from %s to %s did not cause change in hash", firstUsername, secondUsername)
	}
	if firstHash == thirdHash {
		t.Fatalf("change in salt from %s to %s did not cause change in hash", firstHash, thirdHash)
	}
	if firstHash == fourthHash {
		t.Fatalf("change in password from %s to %s did not cause change in hash", firstHash, fourthHash)
	}
	if firstHash != fifthHash {
		t.Fatalf("same input combination produced a different output %s vs %s", firstHash, fifthHash)
	}
}

func TestSecureStringsEqual_True(t *testing.T) {
	first := "aoboajf819n1jt4nlb;;dfpa__f=as;dfm,."
	second := "aoboajf819n1jt4nlb;;dfpa__f=as;dfm,."

	equal := SecureStringsEqual(first, second)

	if !equal {
		t.FailNow()
	}
}

func TestSecureStringsEqual_False(t *testing.T) {
	first := "aoboajf819n1jt4nlb;;dfpa__f=as;dfm,."
	second := "aoboajf819n1jt4nlb;;dfpa__f=as;dfm,.+"

	equal := SecureStringsEqual(first, second)

	if equal {
		t.FailNow()
	}
}
