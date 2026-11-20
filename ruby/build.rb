# frozen_string_literal: true

# Ruby build for the GHAS Me Up demo.
# Renders ruby/index.html from Markdown (kramdown) + ERB.
#
require 'erb'
require 'digest'
require 'fileutils'
require 'securerandom'
require 'time'
require 'kramdown'
require 'rack'
require 'addressable/uri'

ROOT = __dir__
OUT_DIR = File.expand_path(ARGV[0] || '../dist', ROOT)

BUILD_STAMP = Digest::SHA256.file(__FILE__).hexdigest[0, 8]
PREVIEW_TOKEN = SecureRandom.hex(16)

def h(value)
  Rack::Utils.escape_html(value.to_s)
end

def locked_gems
  File.readlines(File.join(ROOT, 'Gemfile.lock')).filter_map do |line|
    match = line.match(/^    (\S+) \(([^)]+)\)$/)
    { name: match[1], version: match[2] } if match
  end
end

markdown = <<~MD
  ## Built with Ruby

  This page was generated at **build time** by `ruby/build.rb` using
  kramdown, ERB, Rack and Addressable.

  Every version below is pinned in `Gemfile.lock`, along with the Ruby
  version itself. Open the repository's **Security** tab to see what
  Dependabot found.
MD

view = {
  title: 'Ruby build',
  accent: '#f85149',
  ruby_version: RUBY_VERSION,
  built_at: Time.now.utc.iso8601,
  content: Kramdown::Document.new(markdown).to_html,
  gems: locked_gems.sort_by { |gem| gem[:name] },
  source_url: Addressable::URI.join('https://github.com/artofthepossible/ghas-me-up/', 'tree/main/ruby').to_s
}

template = ERB.new(File.read(File.join(ROOT, 'template.erb')), trim_mode: '-')
html = template.result_with_hash(view.merge(h: method(:h)))

target = File.join(OUT_DIR, 'ruby')
FileUtils.mkdir_p(target)
File.write(File.join(target, 'index.html'), html)

puts "ruby build complete -> #{File.join(target, 'index.html')}"
