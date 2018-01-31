#!/usr/bin/env ruby

require 'net/http'
require 'icalendar'

ORGANIZERS = { 'h.itten' => 'Heiner Itten',
               'stefan.mumenthaler' => 'Stefan Mumenthaler',
               'dieter.zehr' => 'Dieter Zehr',
               'ps' => 'Pascal Simon'
}.freeze

DOJO_ADDRESS='Äussere Ringstr. 7, Thun, 3600, Schweiz'.freeze

def list_trainings(ics)
  cals = Icalendar::Calendar.parse(ics)
  events = cals.first.events
  events = events.select {|e| e.categories.first == 'Trainings'}
  events = events.select {|e| e.location == DOJO_ADDRESS}

  events.each do |e|
    date = e.dtstart
    str_day = date.strftime("%A")
    str_date = date.strftime("%d.%m.%Y")
    o_email = e.organizer.to_s.gsub('mailto:', '')
    o_label = o_email.gsub(/@\S*$/, '')
    organizer = ORGANIZERS[o_label]
    puts "#{str_date} | #{str_day} | #{organizer}"
  end
end

raise 'please provide months' unless ARGV.first

months = ARGV.first.split(',')

months.each do |m|
  year = Date.today.year
  md = "%02d" % m 
  uri = URI("https://www.shotokai.ch/events/#{year}-#{md}/?ical=1&tribe_display=month")
  ics = Net::HTTP.get(uri).force_encoding('utf-8')
  list_trainings(ics)
end

