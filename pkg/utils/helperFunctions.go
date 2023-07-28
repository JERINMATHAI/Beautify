package utils

import (
	"beautify/pkg/domain"
	"encoding/hex"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// take userId from context
func GetUserIdFromContext(ctx *gin.Context) uint {
	userIdStr, exists := ctx.Get("userId")
	if !exists {
		// userId key does not exist in the context
		return 0 //fmt.Errorf("userId not found in context")
	}
	userId, _ := strconv.ParseUint(userIdStr.(string), 10, 64)
	// if err != nil {
	// 	// Failed to convert userId to uint
	// 	return 0, err
	// }
	return uint(userId)
}

// func GetUserIdFromContext(ctx *gin.Context) uint {
// 	userID := ctx.GetInt("userId")
// 	return uint(userID)
// }

func StringToUint(str string) (uint, error) {
	val, err := strconv.Atoi(str)
	return uint(val), err
}

// generate userName
func GenerateRandomUserName(FirstName string) string {
	suffix := make([]byte, 4)
	numbers := "1234567890"
	rand.Seed(time.Now().UnixMilli())
	for i := range suffix {
		suffix[i] = numbers[rand.Intn(10)]
	}
	userName := (FirstName + string(suffix))
	return strings.ToLower(userName)
}

// generate unique string for sku
func GenerateSKU() string {
	sku := make([]byte, 10)
	rand.Read(sku)
	return hex.EncodeToString(sku)
}

func CompareUserExistingDetails(compare1, compare2 domain.Users) (err error) {
	if compare1.Email == compare2.Email {
		err = appendError(err, "user already exists with this email")
	}
	if compare1.UserName == compare2.UserName {
		err = appendError(err, "user already exists with this username")
	}
	if compare1.Phone == compare2.Phone {
		err = appendError(err, "user already exists with this phone number")
	}
	return err
}

func appendError(err error, message string) error {
	if err == nil {
		return errors.New(message)
	}
	return errors.New(err.Error() + "; " + message)
}

// random coupons
func GenerateCouponCode(couponCodeLenth int) string {
	// letter for coupns
	letters := `ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`
	rand.Seed(time.Now().UnixMilli())
	// creat a byte array of couponCodeLength
	couponCode := make([]byte, couponCodeLenth)
	// loop through the array and randomly pic letter and add to array
	for i := range couponCode {
		couponCode[i] = letters[rand.Intn(len(letters))]
	}
	// convert into string and return the random letter array
	return string(couponCode)
}

func StringToTime(date string) (time.Time, error) {
	layout := "2006-01-02"

	// Parse the string date using the specified layout
	returnDate, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	// Return the parsed time
	return returnDate, nil
}

func GenerateRandomString(length int) string {
	sku := make([]byte, length)
	rand.Read(sku)
	return hex.EncodeToString(sku)
}

func RandomInt(min, max int) int {
	rand.Seed(time.Hour.Nanoseconds())
	return rand.Intn(max-min) + min
}

func GetHashedPassword(password string) (hashedPassword string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return hashedPassword, err
	}
	hashedPassword = string(hash)
	return hashedPassword, nil
}

func ComparePasswordWithHashedPassword(actualpassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(actualpassword))
	return err
}
