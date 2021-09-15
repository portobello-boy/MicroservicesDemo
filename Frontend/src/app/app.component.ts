import { Component, OnInit } from '@angular/core';
import { CalendarView, CalendarEvent, DAYS_OF_WEEK } from 'angular-calendar';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import dayjs from 'dayjs';
import en from 'dayjs/locale/en';


dayjs.locale({
  ...en,
  weekStart: DAYS_OF_WEEK.MONDAY,
});


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  date: Date = new Date();

  events: CalendarEvent[] = [
    // {
    //   start: new Date(this.date.getFullYear(), this.date.getMonth(), this.date.getDate(), 10, 0),
    //   end: new Date(this.date.getFullYear(), this.date.getMonth(), this.date.getDate(), 10, 45),
    //   title: 'interesting title'
    // }
  ]

  viewDate: Date = new Date();
  view: CalendarView = CalendarView.Month;
  CalendarView = CalendarView;

  constructor(private http: HttpClient) {
    this.getData()
  }

  ngOnInit() {
  }

  getData() {
    const url = "http://localhost:3001/enrich/"
    const getAllDataRequest = {
      "type": "getAllEvents",
    }

    // this.http.get("http://localhost:3001/health", {responseType: 'text'}).subscribe(data => {
    //   console.log(data)
    // })
    
    this.http.post(url, getAllDataRequest)
      .subscribe(data => {
        Object.values(data).forEach(event => {
          console.log(event)
          const e = {
            title: event.title,
            start: new Date(event.startTime),
            // end: event.endTime,
            // allDay: event.allDay,
          }
          console.log(e)
          this.events = [
            ...this.events, e
          ]
        })
      })
  }

  setView(view: CalendarView) {
    this.view = view;
  }
}
