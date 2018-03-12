package jianshu

type SpecialSubject struct {
}

func NewSpecialSubject() *SpecialSubject {
	return &SpecialSubject{}
}

func (s *SpecialSubject) GetSpecialSubjectTitle() {}

func (s *SpecialSubject) GetSpecialSubjectPassageNumber() {}

func (s *SpecialSubject) GetSpecialSubjectFollowerNumber() {}

func (s *SpecialSubject) GetSpecialSubjectNotice() {}

func (s *SpecialSubject) GetSpecialSubjectTitleAdministrator() {}

func (s *SpecialSubject) GetSpecialSubjectNewComment() {}

func (s *SpecialSubject) GetSpecialSubjectNewAdd() {}

func (s *SpecialSubject) GetSpecialSubjectHot() {}
