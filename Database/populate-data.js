'use strict';

// @see https://docs.mongodb.com/manual/tutorial/write-scripts-for-the-mongo-shell/

print("AAAAAAAAAA")

var date = new Date();
var d = date.getDate();
var m = date.getMonth();
var y = date.getFullYear();

var data = [
    // {
    //     title: "test",
    //     startTime: new Date(y, m, d),
    //     endTime: new Date(y, m, d + 1),
    //     allDay: false
    //     url: "http://www...."
    //     attendees: ["John Doe", ...]
    // },
    {
        title: "Birthday",
        startTime: new Date(y, m, d),
        allDay: true,
    },
    {
        title: "Daily Standup",
        startTime: new Date(y, m, d, 10, 0).toISOString(),
        endTime: new Date(y, m, d, 10, 15).toISOString(),
    },
    {
        title: "Sprint Planning",
        startTime: new Date(y, m, d + 1, 14).toISOString(),
        endTime: new Date(y, m, d + 1, 15).toISOString(),
        url: "blah"
    },
    {
        title: "Vacation",
        startTime: new Date(y, m, d + 2).toISOString(),
        allDay: true
    },

]

var MONGODB_URI = "mongodb://127.0.0.1:27017/calendar";

var db = connect(MONGODB_URI);

db.events.insertMany(data)

var cursor = db.events.find();
while ( cursor.hasNext() ) {
   printjson( cursor.next() );
}







print(db.events.find())

