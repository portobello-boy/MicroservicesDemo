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
    //     start: new Date(y, m, d),
    //     end: new Date(y, m, d + 1),
    //     allDay: false
    //     url: "http://www...."
    //     guests: ["John Doe", ...]
    // },
    {
        title: "Birthday",
        start: new Date(y, m, d),
        allDay: true,
    },
    {
        title: "Daily Standup",
        start: new Date(y, m, d, 10, 0),
        end: new Date(y, m, d, 10, 15),
    },
    {
        title: "Sprint Planning",
        start: new Date(y, m, d + 1, 14),
        end: new Date(y, m, d + 1, 15),
        url: "blah"
    },
    {
        title: "Vacation",
        start: new Date(y, m, d + 2),
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

