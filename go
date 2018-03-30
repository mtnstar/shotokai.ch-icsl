#!/usr/bin/env ruby

require 'net/http'
require 'icalendar'

ORGANIZERS = { 'h.itten' => 'Heiner Itten',
               'stefan.mumenthaler' => 'Stefan Mumenthaler',
               'erika.sidler' => 'Erika Sidler',
               'dieter.zehr' => 'Dieter Zehr',
               'marc.schneiter' => 'Marc Schneiter',
               'ps' => 'Pascal Simon'
}.freeze

DOJO_ADDRESS='Äussere Ringstr. 7, Thun, 3600, Schweiz'.freeze

@trainings = []

def list_trainings
  events = @trainings.uniq { |t| t.dtstart }

  events = events.collect do |e|
    date = e.dtstart
    str_day = date.strftime("%A")
    str_date = date.strftime("%d.%m.%Y")
    o_email = e.organizer.to_s.gsub('mailto:', '')
    o_label = o_email.gsub(/@\S*$/, '')
    organizer = ORGANIZERS[o_label]
    "#{str_date} | #{str_day} | #{organizer}"
  end

  puts events.uniq.join("\n")
end

def collect_trainings(ics)
  cals = Icalendar::Calendar.parse(ics)
  events = cals.first.events
  events = events.select {|e| e.categories.first == 'Trainings'}
  events = events.select {|e| e.location == DOJO_ADDRESS}
  @trainings += events
end

abort 'USAGE: ./go 1,2,3' unless ARGV.first

months = ARGV.first.split(',')

months.each do |m|
  year = Date.today.year
  md = "%02d" % m 
  uri = URI("https://www.shotokai.ch/events/#{year}-#{md}/?ical=1&tribe_display=month")
  ics = Net::HTTP.get(uri).force_encoding('utf-8')
  collect_trainings(ics)
end

list_trainings
