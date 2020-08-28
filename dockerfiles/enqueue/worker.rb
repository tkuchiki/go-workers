require 'sidekiq'

Sidekiq.configure_server do |config|
  config.redis = { url: 'redis://redis:6379/0' }
end

Sidekiq.configure_client do |config|
  config.redis = { url: 'redis://redis:6379/0' }
end

class TestWorker
  include Sidekiq::Worker

  def perform(*opts)
    puts self.jid
    puts opts[0]
  end
end
