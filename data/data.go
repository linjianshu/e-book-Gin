package data

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"github.com/jinzhu/configor"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func init() {
	var err error

	var Config = struct {
		APPName string `default:"app name"`

		DB struct {
			Name     string
			User     string `default:"root"`
			Password string `required:"true" env:"DBPassword"`
			Port     uint   `default:"3306"`
			Host     string `default:"localhost"`
		}

		Contacts []struct {
			Name  string
			Email string `required:"true"`
		}
	}{}
	err = configor.Load(&Config, "config.yaml")

	//Db, err = gorm.Open(mysql.Open("root:Lalala123#@tcp(localhost:3306)/datashare?charset=utf8&parseTime=true"), &gorm.Config{})
<<<<<<< HEAD
	//Db, err = gorm.Open(mysql.Open("root:Lalala123#@tcp(localhost:3306)/datashare?charset=utf8&parseTime=true"), &gorm.Config{})
	Db, err = gorm.Open(mysql.Open("root:123456@tcp(124.222.114.171:3306)/datashare?charset=utf8&parseTime=true"), &gorm.Config{})
=======
	Db, err = gorm.Open(mysql.Open("root:Lalala123#@tcp(localhost:3306)/datashare?charset=utf8&parseTime=true"), &gorm.Config{})
	//Db, err = gorm.Open(mysql.Open("root:123456@tcp(124.222.114.171:3306)/datashare?charset=utf8&parseTime=true"), &gorm.Config{})
>>>>>>> 558c452 (20220809)
	//Db, err = gorm.Open(mysql.Open(fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", Config.DB.User, Config.DB.Password, Config.DB.Host, Config.DB.Port, Config.DB.Name)), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	return
}

// CreateUUID create a random UUID with from RFC 4122
// adapted from http://github.com/nu7hatch/gouuid
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// Encrypt hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
