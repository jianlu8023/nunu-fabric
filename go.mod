module nunu-fabric

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/gin-gonic/gin v1.9.1
	github.com/go-co-op/gocron v1.28.2
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redis/redismock/v9 v9.0.3
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/pkg/errors v0.9.1
	github.com/sony/sonyflake v1.1.0
	github.com/spf13/viper v1.16.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.24.0
	golang.org/x/crypto v0.9.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gorm.io/driver/mysql v1.5.1
	gorm.io/gorm v1.25.1
)

require (
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/hxx258456/fabric-sdk-go-gm v0.0.7
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/pelletier/go-toml/v2 v2.0.8 // indirect
	github.com/spf13/cast v1.5.1 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace (
	github.com/gin-gonic/gin v1.9.1 => github.com/gin-gonic/gin v1.8.0
	github.com/go-kit/kit => github.com/go-kit/kit v0.8.0
	github.com/mitchellh/mapstructure => github.com/mitchellh/mapstructure v1.3.3
	github.com/spf13/viper => github.com/spf13/viper v0.0.0-20150908122457-1967d93db724
	github.com/zmap/zcrypto => github.com/zmap/zcrypto v0.0.0-20190729165852-9051775e6a2e
	github.com/zmap/zlint => github.com/zmap/zlint v0.0.0-20190806154020-fd021b4cfbeb
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20181228115726-23731bf9ba55
)
