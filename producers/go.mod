module producers

go 1.20

replace models => ../models

require (
	models v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.40
)

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/sony/sonyflake v1.1.0 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	gorm.io/driver/mysql v1.5.1 // indirect
	gorm.io/gorm v1.25.1 // indirect
)
