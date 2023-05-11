package usecase

import (
	"github.com/moonidelight/go_course/lab3/internal/domain/entity"
	"github.com/moonidelight/go_course/lab3/internal/repository"
)

type PsqlImp struct {
	repo *repository.BookRepository
}

func NewPsqlImp(repo *repository.BookRepository) *PsqlImp {
	return &PsqlImp{repo: repo}
}
func (pi *PsqlImp) Create(title, desc string, cost float64) {
	pi.repo.Create(title, desc, cost)
}
func (pi *PsqlImp) GetById(id int) (string, string, float64) {
	return pi.repo.GetById(id)
}
func (pi *PsqlImp) GetAll() []entity.Book {
	return pi.repo.GetAll()
}
func (pi *PsqlImp) UpdateTitleAndDesc(id int, title, desc string) (string, string) {
	return pi.repo.UpdateTitleAndDesc(id, title, desc)
}

func (pi *PsqlImp) DeleteById(id int) {
	pi.repo.DeleteById(id)
}
func (pi *PsqlImp) SearchByTitle(title string) (string, string, float64) {
	return pi.repo.SearchByTitle(title)
}
func (pi *PsqlImp) GetSortedBooksAsc() interface{} {
	return pi.repo.GetSortedBooksAsc()
}
func (pi *PsqlImp) GetSortedBooksDesc() interface{} {
	return pi.repo.GetSortedBooksDesc()
}
