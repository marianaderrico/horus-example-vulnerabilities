package product

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/wiliansilvazup/horus/example/vulnerabilities/internal/entities"
	"time"
)

type Product struct {
	entities.Base
	Name string `json:"name"`
}

func (p *Product) TableName() string {
	return "products"
}

func (p *Product) Bytes() []byte {
	bytes, _ := json.Marshal(p)
	return bytes
}

func (p *Product) GenerateID() {
	p.ID = uuid.New()
}

func (p *Product) SetCreatedAt() {
	p.CreatedAt = time.Now()
}

func (p *Product) SetUpdatedAt() {
	p.UpdatedAt = time.Now()
}
