TCP Chat Application

A simple TCP-based chat application implemented in Go, featuring multiple chat rooms, private messaging, and nickname support.

Features

	•	Multi-Room Chat: Join or create rooms with /join ROOM_NAME and chat with others in the same room.
	•	Nicknames: Set a custom nickname using /nick NAME.
	•	Messaging: Send messages within the current room using /msg MESSAGE.
	•	Room Listings: View available rooms with /rooms.
	•	Exit: Leave the chat with /quit.

Usage

	1.	Run the server:

go run main.go


	2.	Connect via Telnet or Netcat:

telnet localhost 8888
# or
nc localhost 8888



Commands

	•	/nick NAME - Set your nickname.
	•	/join ROOM_NAME - Join or create a chat room.
	•	/rooms - List all available rooms.
	•	/msg MESSAGE - Send a message in the current room.
	•	/quit - Leave the chat.