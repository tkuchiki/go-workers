FROM ruby:2.7.1-slim-buster

RUN mkdir /app
ENV APP_ROOT /app
WORKDIR $APP_ROOT

ADD ./Gemfile $APP_ROOT/
ADD ./Gemfile.lock $APP_ROOT/
ADD ./worker.rb $APP_ROOT/

RUN bundle update --bundler && bundle install

