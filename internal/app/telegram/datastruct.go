package telegram

type Message struct {
	MessageID int64
	Chat      Chat
	From      User
	Date      int64
	Text      string
}

type Chat struct {
	ID        int64
	FirstName string
	Username  string
	Type      string
}

type User struct {
	ID           int64
	IsBot        bool
	FirstName    string
	Username     string
	LanguageCode string
}

// type Chat struct {
// id 	Integer 	Unique identifier for this chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
// type 	string 	Type of chat, can be either “private”, “group”, “supergroup” or “channel”
// title 	string 	Optional. Title, for supergroups, channels and group chats
// username 	string 	Optional. Username, for private chats, supergroups and channels if available
// first_name 	string 	Optional. First name of the other party in a private chat
// last_name 	string 	Optional. Last name of the other party in a private chat
// photo 	ChatPhoto 	Optional. Chat photo. Returned only in getChat.
// bio 	string 	Optional. Bio of the other party in a private chat. Returned only in getChat.
// has_private_forwards 	True 	Optional. True, if privacy settings of the other party in the private chat allows to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
// description 	string 	Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
// invite_link 	string 	Optional. Primary invite link, for groups, supergroups and channel chats. Returned only in getChat.
// pinned_message 	Message 	Optional. The most recent pinned message (by sending date). Returned only in getChat.
// permissions 	ChatPermissions 	Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
// slow_mode_delay 	Integer 	Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user; in seconds. Returned only in getChat.
// message_auto_delete_time 	Integer 	Optional. The time after which all messages sent to the chat will be automatically deleted; in seconds. Returned only in getChat.
// has_protected_content 	True 	Optional. True, if messages from the chat can't be forwarded to other chats. Returned only in getChat.
// sticker_set_name 	string 	Optional. For supergroups, name of group sticker set. Returned only in getChat.
// can_set_sticker_set 	True 	Optional. True, if the bot can change the group sticker set. Returned only in getChat.
// linked_chat_id 	Integer 	Optional. Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats. This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier. Returned only in getChat.
// location 	ChatLocation 	Optional. For supergroups, the location to which the supergroup is connected. Returned only in getChat.
// }

// type Message struct {

// }

// type User struct {
// 	id                          int64
// 	is_bot                      bool
// 	first_name                  string
// 	last_name                   string
// 	username                    string
// 	language_code               string
// 	can_join_groups             bool
// 	can_read_all_group_messages bool
// 	supports_inline_queries     bool
// }
