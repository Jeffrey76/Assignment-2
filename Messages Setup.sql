CREATE database edufi;
USE edufi;

/*ID
SenderID
ReceiverID
TimeSent
TimeRead (possible implementation)
Header
Content*/

create table Conversation(
	ID               int PRIMARY KEY,
	InitiatorID      int,
	RecipientID     int,
	StartTime        datetime,
	NoofMessages     int,
	ConversationName varchar(30)
    );
insert into Conversation(ID,InitiatorID,RecipientID,StartTime,NoofMessages,ConversationName) 
values("1","1","2",now(),0,"TestConversation");
insert into Conversation(ID,InitiatorID,RecipientID,StartTime,NoofMessages,ConversationName) 
values("2","2","1",now(),0,"SecondConversation");
use edufi;
create table Replies(
	ID               int PRIMARY KEY,
    ConversationID   int,
	SenderID         int,
	ReceiverID       int,
	TimeSent         datetime,
	TimeRead         datetime,
	Header varchar(30),
    Content varchar(500)
    );
insert into Replies(ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content) 
values("1","1","1","2",now(),"First Header","This Message is the First Message");
insert into Replies(ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content) 
values("2","1","1","2",now(),"Second Header","This Message is the Second Message");
insert into Replies(ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content) 
values("3","1","2","1",now(),"Third Header","This is Person 2's Reply");
insert into Replies(ID,ConversationID,SenderID,ReceiverID,TimeSent,Header,Content) 
values("4","1","2","1",now(),"Fourth Header","This is Person 2's Second Reply");


/*select ID,SenderID,ReceiverID,TimeSent,Header,Content from Replies ;*/
select * from Conversation ;
select * from Replies;

/*select * from replies;
SELECT * FROM Conversation WHERE InitiatorID='101' OR RecipientID='101' ORDER BY ID DESC;
*/

