package database

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/pbkdf2"
)

/*
	PASSWORD
*/

type Password_cache struct {
	Hash       string
	Salt       string
	Iterations int
}

const (
	PASSWORD_SALT_LENGTH          = 32
	PASSWORD_HASH_LENGTH          = 32
	PASSWORD_SALT_CHARS           = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	PASSWORD_MIN_ITERATIONS_COUNT = 5000
)

func Password_cache_gen(password string) Password_cache {
	var cache Password_cache

	cache.Salt = Random(uint(PASSWORD_SALT_LENGTH), []rune(PASSWORD_SALT_CHARS))
	cache.Iterations = int(rand.Uint32())%1000 + PASSWORD_MIN_ITERATIONS_COUNT
	cache.Hash = hex.EncodeToString(pbkdf2.Key(
		[]byte(password),
		[]byte(cache.Salt),
		int(cache.Iterations),
		PASSWORD_HASH_LENGTH,
		sha256.New,
	))

	return cache
}

func Password_cache_check(password string, cache Password_cache) bool {
	rehash := hex.EncodeToString(pbkdf2.Key(
		[]byte(password),
		[]byte(cache.Salt),
		int(cache.Iterations),
		PASSWORD_HASH_LENGTH,
		sha256.New,
	))

	if rehash == cache.Hash {
		return true
	}

	return false
}

func Password_verify(db *sql.DB, something string, password string) (bool, int) {
	var err error
	var userid int
	var cache Password_cache
	check, _ := regexp.MatchString("^(?=.{1,64}@)[A-Za-z0-9_-]+(\\.[A-Za-z0-9_-]+)*@[^-][A-Za-z0-9-]+(\\.[A-Za-z0-9-]+)*(\\.[A-Za-z]{2,})$", something)
	if check {
		err = db.QueryRow(fmt.Sprintf("SELECT password_hash, password_salt,"+
			"password_iterations, userid FROM users WHERE email = '%s';",
			something)).Scan(&cache.Hash, &cache.Salt, &cache.Iterations, &userid)
	} else {
		err = db.QueryRow(fmt.Sprintf("SELECT password_hash, password_salt,"+
			"password_iterations, userid FROM users WHERE username = '%s';",
			something)).Scan(&cache.Hash, &cache.Salt, &cache.Iterations, &userid)
	}
	if err != nil {
		fmt.Println(err.Error())
		return false, userid
	}

	if Password_cache_check(password, cache) {
		return true, userid
	} else {
		return false, userid
	}
}

/*
	TOKEN
*/

const (
	TOKEN_SALT_LENGTH = 32
	TOKEN_SALT_CHARS  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

type tokens_array struct {
	Objects []string `json:"objects"`
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func Token_gen(db *sql.DB, userid int) string {
	token_str := Random(TOKEN_SALT_LENGTH, []rune(TOKEN_SALT_CHARS))

	// append tokens
	var tokens tokens_array
	db.QueryRow(fmt.Sprintf("SELECT tokens FROM users WHERE userid = %d;", userid)).Scan(&tokens)
	tokens.Objects = append(tokens.Objects, string(token_str))
	tokens_str, err := json.Marshal(tokens)
	if err != nil {
		return ""
	}
	db.Exec(fmt.Sprintf("UPDATE users SET tokens = '%s';", string(tokens_str)))

	return token_str
}

func Token_delete(db *sql.DB, token string, userid int) {
	var tokens tokens_array
	db.QueryRow(fmt.Sprintf("SELECT tokens FROM users WHERE userid = %d;", userid)).Scan(&tokens)
	for i := 0; i < len(tokens.Objects); i++ {
		if tokens.Objects[i] == token {
			tokens.Objects = remove(tokens.Objects, i)
			break
		}
	}
	tokens_str, err := json.Marshal(tokens)
	if err != nil {
		return
	}
	db.Exec(fmt.Sprintf("UPDATE users SET tokens = '%s';", string(tokens_str)))
}

func Token_find(db *sql.DB, token string, userid int) bool {
	var tokens tokens_array
	var tokens_str []byte
	db.QueryRow(fmt.Sprintf("SELECT tokens FROM users WHERE userid = %d;", userid)).Scan(&tokens_str)
	err := json.Unmarshal(tokens_str, tokens)
	if err != nil {
		return false
	}
	for i := 0; i < len(tokens.Objects); i++ {
		if tokens.Objects[i] == token {
			return true
		}
	}
	tokens_str, err = json.Marshal(tokens)
	if err != nil {
		return false
	}
	db.Exec(fmt.Sprintf("UPDATE users SET tokens = '%s';", string(tokens_str)))

	return false
}
