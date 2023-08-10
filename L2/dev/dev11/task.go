package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var yearsMap = make(map[int]*Year)

type Command interface {
	do(year, month, day int, e *Event) error
}
type appendEvent struct{}

func (a *appendEvent) do(year, month, day int, e *Event) error {
	if year < time.Now().Year() {
		return errors.New("this year has already passed")
	}
	if _, ok := yearsMap[year]; !ok {
		yearsMap[year] = NewYear()
	}
	err := yearsMap[year].append(month, day, e)
	return err
}

type deleteEvent struct{}

func (d *deleteEvent) do(year, month, day int, e *Event) error {
	if _, ok := yearsMap[year]; !ok {
		return errors.New("year not found")
	}
	if year < time.Now().Year() {
		return errors.New("this year has already passed")
	}
	err := yearsMap[year].delete(month, day, e)
	return err
}

type updateEvent struct{}

func (u *updateEvent) do(year, month, day int, e *Event) error {
	if year < time.Now().Year() {
		return errors.New("this year has already passed")
	}
	if _, ok := yearsMap[year]; !ok {
		return errors.New("year not found")
	}
	err := yearsMap[year].update(month, day, e)
	return err
}

type Event struct {
	userId      int
	event       string
	description string
	time        string
}

func NewEvent() *Event {
	return &Event{}
}

type Period interface {
	getEvents(int) []*Event
}
type Day struct {
	events []*Event
	week   *Week
}

func NewDay() *Day {
	return &Day{make([]*Event, 0), nil}
}
func (d *Day) getEvents(userId int) []*Event {
	res := make([]*Event, 0)
	for _, e := range d.events {
		if e.userId == userId {
			res = append(res, e)
		}
	}
	return res
}
func (d *Day) append(e *Event) {
	d.events = append(d.events, e)
}
func (d *Day) update(e *Event) error {
	isDone := false
	for _, event := range d.events {
		if event.userId == e.userId && event.event == e.event {
			event.description = e.description
			isDone = true
		}
	}
	if isDone {
		return nil
	}
	return errors.New("event not found")
}
func (d *Day) delete(e *Event) error {
	isDone := false
	for i, event := range d.events {
		if event.userId == e.userId && event.event == e.event {
			event.description = e.description
			d.events = append(d.events[:i], d.events[i+1:]...) ///???
			isDone = true
		}
	}
	if isDone {
		return nil
	}
	return errors.New("event not found")
}

type Week struct {
	days [7]*Day
}

func (w *Week) getEvents() []*Event {
	result := make([]*Event, 0)
	for _, d := range w.days {
		if d != nil {
			result = append(result, d.events...)
		}
	}
	return result
}

type Month struct {
	days    [31]*Day
	maxDays int
}

func NewMonth(maxD int) *Month {
	return &Month{maxDays: maxD}
}
func (m *Month) getEvents(userId int) []*Event {
	result := make([]*Event, 0)
	for _, d := range m.days {
		for _, e := range d.events {
			if e.userId == userId {
				result = append(result, e)
			}
		}

	}
	return result
}
func (m *Month) checkDays(day int) error {
	if day > m.maxDays {
		return errors.New("incorrect day")
	}
	return nil
}
func (m *Month) append(day int, e *Event) error {
	err := m.checkDays(day)
	if err != nil {
		return err
	}
	if m.days[day-1] == nil {
		m.days[day-1] = NewDay()
	}
	m.days[day-1].append(e)
	return nil
}
func (m *Month) update(day int, e *Event) error {
	err := m.checkDays(day)
	if err != nil {
		return err
	}
	if m.days[day-1] == nil {
		return errors.New("day not found")
	}
	err = m.days[day-1].update(e)
	return err
}
func (m *Month) delete(day int, e *Event) error {
	err := m.checkDays(day)
	if err != nil {
		return err
	}
	if m.days[day-1] == nil {
		return errors.New("day not found")
	}
	err = m.days[day-1].delete(e)
	return err
}

type Year struct {
	month    [12]*Month
	Weeks    [53]Week
	maxMonth int
}

func NewYear() *Year {

	return &Year{maxMonth: 12}
}
func (y *Year) checkMonth(month int) error {
	if month > y.maxMonth {
		return errors.New("incorrect month")
	}
	return nil
}
func (y *Year) getEvents() []*Event {
	return make([]*Event, 0)
}
func (y *Year) append(month, day int, e *Event) error {
	err := y.checkMonth(month)
	if err != nil {
		return err
	}
	if y.month[month] == nil {
		y.month[month] = NewMonth(maxDaysInMonth(month))
	}
	err = y.month[month].append(day, e)
	if err != nil {
		return err
	}
	return err
}
func (y *Year) update(month, day int, e *Event) error {
	err := y.checkMonth(month)
	if err != nil {
		return err
	}
	if y.month[month] == nil {
		return errors.New("month not found")
	}
	err = y.month[month].update(day, e)
	return err
}

func (y *Year) delete(month, day int, e *Event) error {
	err := y.checkMonth(month)
	if err != nil {
		return err
	}
	if y.month[month] == nil {
		return errors.New("month not found")
	}
	err = y.month[month].delete(day, e)
	return err
}
func maxDaysInMonth(monthNumber int) int {
	if monthNumber == 2 {
		return 29
	}
	if (monthNumber%2 == 0) == (monthNumber < 8) {
		return 31
	}
	return 30
}
func doEvent(com Command, e *Event, time time.Time) error {

	year, week := time.ISOWeek()
	month, day := int(time.Month()), time.Day()
	err := com.do(year, month, day, e)
	return err
}
func getEvent(per Period, id int) {
	per.getEvents(id)
}
func getPeriod(date time.Time) Period {

}
func CreateEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "there are must be POST method", 503)
		return
	}

	event, err := Parse(req)
	if err != nil {
		http.Error(rw, "parse error", 400)
		return
	}
	eventTime, err := time.Parse("2000-01-01", event.time)
	if err != nil {
		http.Error(rw, "incorrect date", 400)
		return
	}

	err = doEvent(&appendEvent{}, event, eventTime)
	if err != nil {
		http.Error(rw, "err", 503)
		return
	}
}
func UpdateEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "there are must be POST method", 503)
		return
	}

	event, err := Parse(req)
	if err != nil {
		http.Error(rw, "parse error", 400)
		return
	}
	eventTime, err := time.Parse("2000-01-01", event.time)
	if err != nil {
		http.Error(rw, "incorrect date", 400)
		return
	}

	err = doEvent(&updateEvent{}, event, eventTime)
	if err != nil {
		http.Error(rw, "err", 503)
		return
	}
}
func DeleteEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(rw, "there are must be POST method", 503)
		return
	}

	event, err := Parse(req)
	if err != nil {
		http.Error(rw, "parse error", 400)
		return
	}
	eventTime, err := time.Parse("2000-01-01", event.time)
	if err != nil {
		http.Error(rw, "incorrect date", 400)
		return
	}

	err = doEvent(&deleteEvent{}, event, eventTime)
	if err != nil {
		http.Error(rw, "err", 503)
		return
	}
}
func EventsDay(rw http.ResponseWriter, req *http.Request) {
	userId, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(rw, "user_id error", 400)
		return
	}
	date, err := time.Parse("2000-01-01", req.URL.Query().Get("date"))

	if err != nil {
		http.Error(rw, "incorrect date", 400)
		return
	}
	err = getEvent()
	if err != nil {
		http.Error(rw, "err", 503)
		return
	}
}
func EventsWeek(rw http.ResponseWriter, r *http.Request) {

}
func EventsMonth(rw http.ResponseWriter, r *http.Request) {

}
func main() {
	port := "8080"
	http.HandleFunc("/create_event", CreateEvent)
	http.HandleFunc("/update_event", UpdateEvent)
	http.HandleFunc("/delete_event", DeleteEvent)
	http.HandleFunc("/events_for_day", EventsDay)
	http.HandleFunc("/events_for_week", EventsWeek)
	http.HandleFunc("/events_for_month", EventsMonth)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Print(err)
		return
	}
}
func Parse(req *http.Request) (*Event, error) {
	var event *Event
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, *event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
