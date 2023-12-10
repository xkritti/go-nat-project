package iaarrepository

import (
	"go-nat-project/models"

	"gorm.io/gorm"
)

type IaarRepositoryDB struct {
	DB *gorm.DB
}

func NewIaarRepositoryDB(db *gorm.DB) *IaarRepositoryDB {
	return &IaarRepositoryDB{
		DB: db,
	}
}

func (repo *IaarRepositoryDB) GetEngIaar(hashCid string) (*models.EngIaar, error) {
	var engIaar models.EngIaar

	sql := `
		select
			c."name",
			c.level_range ,
			c.school ,
			c.province ,
			c.exam_type ,
			es.score_pt_expression ,
			es.score_pt_reading ,
			es.score_pt_structure ,
			es.score_pt_vocabulary ,
			es.total_score ,
			es.province_rank ,
			es.region_rank
		from
			eng_scores es
		inner join competitors c 
		on
			c.cid = es.hash_cid
		where
			c.cid = ?;
	`

	err := repo.DB.Raw(sql, hashCid).Scan(&engIaar).Error
	if err != nil {
		return nil, err
	}
	return &engIaar, nil
}
