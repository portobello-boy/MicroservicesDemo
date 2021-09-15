import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CalendarComponent } from './calendar/calendar.component';

import { CalendarModule, DateAdapter, CalendarDateFormatter, CalendarMomentDateFormatter, MOMENT } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/moment';
import { HttpClientModule, HttpClient } from '@angular/common/http';

import dayjs from 'dayjs';

export function dayjsAdapterFactory() {
  return adapterFactory(dayjs);
}

@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    
    CalendarModule.forRoot(
    {
      provide: DateAdapter,
      useFactory: dayjsAdapterFactory,
    },
    {
      dateFormatter: {
        provide: CalendarDateFormatter,
        useClass: CalendarMomentDateFormatter,
      },
    })
  ],
  providers: [
    HttpClientModule, 
    {
      provide: MOMENT,
      useValue: dayjs,
    },],
  bootstrap: [AppComponent]
})
export class AppModule { }
