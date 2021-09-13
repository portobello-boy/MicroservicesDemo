'use strict';

// @see https://docs.mongodb.com/manual/tutorial/write-scripts-for-the-mongo-shell/

print("AAAAAAAAAA")

var date = new Date();
var d = date.getDate();
var m = date.getMonth();
var y = date.getFullYear();

var MONGODB_URI = "mongodb://127.0.0.1:27017/calendar";

var db = connect(MONGODB_URI);

db.events.insertOne({
    title: "test",
    start: new Date(y, m, d),
    end: new Date(y, m, d + 1),
})

var cursor = db.events.find();
while ( cursor.hasNext() ) {
   printjson( cursor.next() );
}







print(db.events.find())

