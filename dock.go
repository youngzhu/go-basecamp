package basecamp

type dockType string

const (
	TypeMessageBoard  dockType = "message_board"
	TypeTodoSet       dockType = "todoset"
	TypeVault         dockType = "vault"
	TypeChat          dockType = "chat"
	TypeSchedule      dockType = "schedule"
	TypeQuestionnaire dockType = "questionnaire"
	TypeInbox         dockType = "inbox"
	TypeKanbanBoard   dockType = "kanban_board"
)

type dock struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Name     dockType `json:"name"`
	Enabled  bool     `json:"enabled"`
	Position int      `json:"position"`
	Url      string   `json:"url"`
	AppUrl   string   `json:"app_url"`
}
