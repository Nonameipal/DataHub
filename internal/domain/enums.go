package domain


type FolderType int

const (
	FolderTypeVideos    FolderType = 0
	FolderTypeDocuments FolderType = 1
)


type TestType int

const (
	TestTypeQuiz TestType = 0
	TestTypeTest TestType = 1
)

type QuestionType int

const (
	QuestionTypeMultipleChoice QuestionType = 1
	QuestionTypeOpenEnded      QuestionType = 2
)

type SessionStatus int

const (
	SessionStatusInProgress SessionStatus = 1
	SessionStatusFinished   SessionStatus = 2
	SessionStatusCancelled  SessionStatus = 3
)

type Roles int

const (
	RolesBasic      Roles = 0
	RolesAdmin      Roles = 1
	RolesSuperAdmin Roles = 2
)
