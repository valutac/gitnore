class Gitnore < Formula
  desc "👨‍💻 gitignore super power"
  homepage "https://valutac.com"
  url "https://github.com/valutac/gitnore/archive/0.1.0.tar.gz"
  sha256 "0c8b7f9d5e79b87aa7fe508ae4613c0feb93685a57e930298c322d999ed27af7"
  head "https://github.com/valutac/gitnore.git"

  depends_on "go" => :build
  depends_on :hg => :build

  def install
    ENV["GOPATH"] = buildpath

    system "go", "build", "-o", "gitnore"
    bin.install "gitnore"
  end

  test do
   system "#{bin}/gitnore"
  end
end
