package models

import "time"

type Produto struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Nome           string    `gorm:"not null" json:"nome"`
	Descricao      string    `json:"descricao"`
	Preco          float64   `gorm:"not null" json:"preco"`
	Categoria      string    `gorm:"not null" json:"categoria"`
	DataCriacao    time.Time `gorm:"autoCreateTime" json:"data_criacao"`
	DataAtualizacao time.Time `gorm:"autoUpdateTime" json:"data_atualizacao"`
}
