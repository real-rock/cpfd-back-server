package repo

import (
	"cpfd-back/internal/core/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateActivityLogParams struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Action bool   `json:"action"`
}

type GetIndoorPropertyLogsParams struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

type CreateIndoorPropertyLogParams struct {
	Type      string  `gorm:"column:type json:"type"`
	Value     float32 `gorm:"column:value json:"value"`
	MachineID string  `gorm:"column:machine json:"machine_id"`
}

type GetParticleLogsWithDatesParams struct {
	Start time.Time `json:"start" form:"start"`
	End   time.Time `json:"end" form:"end"`
}

type CreateParticleLogParams struct {
	MachineID string  `gorm:"column:machine" json:"machine_id"`
	Pm1       float32 `gorm:"column:pm1" json:"pm1"`
	Pm25      float32 `gorm:"column:pm2_5" json:"pm2_5"`
	Pm10      float32 `gorm:"column:pm10" json:"pm10"`
}

type CreateMachineParams struct {
	Description string `json:"description"`
	MachineNum  int    `json:"machine_num"`
	Location    string `json:"location"`
	PostalCode  string `json:"postal_code"`
}

type RepoManager interface {
	CreateMachine(p CreateMachineParams) error
	GetMachine(machineUid uuid.UUID) (*model.Machine, error)
	GetMachineWithNum(num int) (*model.Machine, error)
	GetMachineNum(machineUid uuid.UUID) (int, error)

	GetActivityLogs() ([]model.Activity, error)
	GetCurrentActivity() (map[string]bool, error)
	CreateActivityLog(p CreateActivityLogParams) error

	CreateIndoorPropertyLog(p CreateIndoorPropertyLogParams) error
	GetAllIndoorPropertyLogs() ([]model.IndoorProperty, error)
	GetIndoorPropertyLogsWithDates(p GetIndoorPropertyLogsParams) ([]model.IndoorProperty, error)

	GetAllParticleLogs() ([]model.Particle, error)
	GetParticleLogsWithDates(p GetParticleLogsWithDatesParams) ([]model.Particle, error)
	CreateParticleLog(p CreateParticleLogParams) error
}

func New(mysqlDb *gorm.DB, redisDb *redis.Client) RepoManager {
	r := &Repo{
		Mysql: mysqlDb,
		Redis: redisDb,
	}
	r.Init()
	return r
}

type Repo struct {
	Mysql *gorm.DB
	Redis *redis.Client
}
